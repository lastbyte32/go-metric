package storage

import (
	"database/sql"
	"encoding/json"
	"errors"
	"os"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/config"
	"github.com/lastbyte32/go-metric/internal/metric"
)

const (
	sqlUpdateCounter = `UPDATE metrics SET counter = ? WHERE id = ?`
	sqlUpdateGauge   = `UPDATE metrics SET gauge = ? WHERE id = ?`

	sqlInsertCounter = `INSERT INTO metrics (id, mtype, counter) VALUES(?,?,?)`
	sqlInsertGauge   = `INSERT INTO metrics (id, mtype, gauge) VALUES(?,?,?)`
)

type sqlStorage struct {
	db         *sqlx.DB
	logger     *zap.SugaredLogger
	pathToFile string
	fileMutex  sync.RWMutex
}

type rowMetric struct {
	ID      string
	Mtype   metric.MType
	Gauge   sql.NullFloat64
	Counter sql.NullInt64
}

func (store *sqlStorage) Get(name string) (metric.IMetric, bool) {
	store.logger.Infof("store: get %s", name)
	var row rowMetric

	err := store.db.Get(&row, "SELECT * FROM metrics WHERE id = $1", name)
	if err != nil {
		store.logger.Infof("get err: %s", err)
		return nil, false
	}
	switch row.Mtype {
	case metric.GAUGE:

		return metric.NewGauge(name, row.Gauge.Float64), true
	case metric.COUNTER:
		return metric.NewCounter(name, row.Counter.Int64), true
	default:
		return nil, false
	}

}

func (store *sqlStorage) All() map[string]metric.IMetric {
	var rows []rowMetric
	if err := store.db.Select(&rows, "SELECT * FROM metrics"); err != nil {
		return nil
	}
	metrics := make(map[string]metric.IMetric)
	for _, row := range rows {
		store.logger.Info(row)
		switch row.Mtype {
		case metric.GAUGE:
			metrics[row.ID] = metric.NewGauge(row.ID, row.Gauge.Float64)
		case metric.COUNTER:
			metrics[row.ID] = metric.NewCounter(row.ID, row.Counter.Int64)
		}
	}
	return metrics
}

func (store *sqlStorage) Update(name string, value string, metricType metric.MType) error {
	store.logger.Infof("Update [%s] %s -> %s", metricType, name, value)
	oneMetric, ok := store.Get(name)
	if !ok {
		store.insert(name, value, metricType)
		return nil
	}
	store.logger.Infof("SetValue [%s] %s -> %s", metricType, name, value)

	err := oneMetric.SetValue(value)
	if err != nil {
		return err
	}
	store.logger.Infof("SetValue END [%s] %s -> %s", metricType, name, value)

	sqlQuery := ""
	switch metricType {
	case metric.GAUGE:
		sqlQuery = `UPDATE metrics SET gauge = cast($1 as double precision) WHERE id = $2`
	case metric.COUNTER:
		sqlQuery = `UPDATE metrics SET counter = $1 WHERE id = $2`
	default:
		return errors.New("unknown type")
	}
	store.logger.Infof("SQL: %s, value %s -> %v", sqlQuery, name, value)

	_, err = store.db.Exec(sqlQuery, oneMetric.ToString(), name)
	if err != nil {
		store.logger.Error(err)
		return err
	}
	return nil
}

func (store *sqlStorage) insert(name string, value string, metricType metric.MType) error {
	store.logger.Infof("INSERT %s -> %s", name, value)

	sqlQuery := ""
	switch metricType {
	case metric.GAUGE:
		sqlQuery = `INSERT INTO metrics (id, mtype, gauge) VALUES($1,$2,$3)`
	case metric.COUNTER:
		sqlQuery = `INSERT INTO metrics (id, mtype, counter) VALUES($1,$2,$3)`
	default:
		return errors.New("unknown type")
	}
	_, err := store.db.Exec(sqlQuery, name, metricType, value)
	if err != nil {
		store.logger.Error(err)
		return err
	}
	return nil
}

func (store *sqlStorage) Close() error {
	err := store.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewSQLStorage(l *zap.SugaredLogger, config config.Configurator) IStorage {
	db, err := sqlx.Connect("pgx", config.GetDSN())
	if err != nil {
		//return nil, err
		l.Error(err)
	}
	store := &sqlStorage{
		pathToFile: config.GetStoreFile(),
		db:         db,
		logger:     l,
	}

	if err := store.migration(); err != nil {
		l.Error(err)
	}
	if config.IsRestore() {
		store.restore()
	}
	return store
}

func (store *sqlStorage) migration() error {
	store.logger.Info("Migration")

	createTable := `CREATE TABLE IF NOT EXISTS metrics (
	id varchar(255) NOT NULL,
	mtype varchar(255) NOT NULL,
	gauge double precision,
	counter bigint,
	CONSTRAINT metrics_pkey PRIMARY KEY (id))`
	//dropTableIfExisis := `DROP TABLE IF EXISTS metrics`
	//createCounterTable := `CREATE TABLE IF NOT EXISTS metrics (name VARCHAR (128) UNIQUE NOT NULL, value BIGINT NOT NULL)`
	//createGaugeTable := `CREATE TABLE IF NOT EXISTS gauge    (name VARCHAR (128) UNIQUE NOT NULL, value DOUBLE PRECISION NOT NULL)`

	migrations := [...]string{
		//dropTableIfExisis,
		createTable,
		//createCounterTable,
		//createGaugeTable,
	}

	for result, migration := range migrations {
		store.logger.Infof("SQL: %s", migration)
		_, err := store.db.Exec(migration)
		if err != nil {
			return err
		}
		store.logger.Infof("Row: %d", result)
	}
	return nil
}

func (store *sqlStorage) openFile(mode int) (*os.File, error) {
	file, err := os.OpenFile(store.pathToFile, mode, filePermissionsDefault)
	if err != nil {
		store.logger.Infof("err open pathToFile: %s, [%s]", store.pathToFile, err.Error())
		return nil, err
	}
	return file, err
}

func (store *sqlStorage) restore() {
	store.logger.Infof("restore from file: %s", store.pathToFile)
	store.fileMutex.RLock()
	defer store.fileMutex.RUnlock()
	file, err := store.openFile(readOnlyMode)
	if err != nil {
		return
	}
	defer file.Close()

	metricsFromFile := make(map[string]metric.Metrics)
	err = json.NewDecoder(file).Decode(&metricsFromFile)
	if err != nil {
		store.logger.Infof("file decode err: %s", err)
		return
	}

	for _, item := range metricsFromFile {
		if err := store.Update(item.ID, item.GetValueAsString(), metric.MType(item.MType)); err != nil {
			store.logger.Infof("update %s", err.Error())
			return
		}
	}
}
