package metric

type MType string

const (
	COUNTER MType = "counter"
	GAUGE   MType = "gauge"
)

type Metric interface {
	GetName() string
	GetType() MType
	ToString() string
	Increase(int642 int64)
}
