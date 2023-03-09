package storage

import (
	"context"
	"database/sql"
	"errors"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" //
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/config"
	"github.com/lastbyte32/go-metric/internal/metric"
)

const dbTimeoutDefault = time.Second * 10

const (
	sqlCreateTable = `CREATE TABLE IF NOT EXISTS metrics (
	id varchar(255) NOT NULL,
	mtype varchar(255) NOT NULL,
	gauge double precision,
	counter bigint,
	CONSTRAINT metrics_pkey PRIMARY KEY (id))`
	sqlGetMetric     = `SELECT * FROM metrics WHERE id = $1`
	sqlAllMetrics    = `SELECT * FROM metrics`
	sqlUpdateCounter = `UPDATE metrics SET counter = $1 WHERE id = $2`
	sqlUpdateGauge   = `UPDATE metrics SET gauge = cast($1 as double precision) WHERE id = $2`
	sqlInsertCounter = `INSERT INTO metrics (id, mtype, counter) VALUES($1,$2,$3)`
	sqlInsertGauge   = `INSERT INTO metrics (id, mtype, gauge) VALUES($1,$2,$3)`
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
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeoutDefault)
	defer cancel()
	var row rowMetric
	err := store.db.GetContext(ctx, &row, sqlGetMetric, name)
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
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeoutDefault)
	defer cancel()
	if err := store.db.SelectContext(ctx, &rows, sqlAllMetrics); err != nil {
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
		sqlQuery = sqlUpdateGauge
	case metric.COUNTER:
		sqlQuery = sqlUpdateCounter
	default:
		return errors.New("unknown type")
	}
	store.logger.Infof("SQL: %s, value %s -> %v", sqlQuery, name, value)

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeoutDefault)
	defer cancel()
	_, err = store.db.ExecContext(ctx, sqlQuery, oneMetric.ToString(), name)
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
		sqlQuery = sqlInsertGauge
	case metric.COUNTER:
		sqlQuery = sqlInsertCounter
	default:
		return errors.New("unknown type")
	}
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeoutDefault)
	defer cancel()
	_, err := store.db.ExecContext(ctx, sqlQuery, name, metricType, value)
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
	db, err := connect(config.GetDSN())
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
	return store
}

func (store *sqlStorage) migration() error {
	store.logger.Info("Migration")

	migrations := [...]string{
		sqlCreateTable,
	}
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeoutDefault)
	defer cancel()
	for result, migration := range migrations {
		store.logger.Infof("SQL: %s", migration)
		_, err := store.db.ExecContext(ctx, migration)
		if err != nil {
			return err
		}
		store.logger.Infof("Row: %d", result)
	}
	return nil
}

func connect(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Ping(dsn string) error {
	db, err := connect(dsn)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}
