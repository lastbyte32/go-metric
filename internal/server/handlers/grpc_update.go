package handlers

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	mproto "github.com/lastbyte32/go-metric/api/proto"
	"github.com/lastbyte32/go-metric/internal/metric"
	"github.com/lastbyte32/go-metric/internal/storage"
)

var _ mproto.MetricsServer = (*grpcUpdateHandler)(nil)

type grpcUpdateHandler struct {
	store  storage.IStorage
	logger *zap.SugaredLogger
	mproto.UnimplementedMetricsServer
}

func (s *grpcUpdateHandler) Update(stream mproto.Metrics_UpdateServer) error {
	for {
		m, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&emptypb.Empty{})
			}
			return err
		}
		var newMetric metric.IMetric
		switch m.GetType() {
		case mproto.Types_COUNTER:
			newMetric = metric.NewCounter(m.GetCounter().GetId(), m.GetCounter().GetDelta())
		case mproto.Types_GAUGE:
			newMetric = metric.NewGauge(m.GetGauge().GetId(), m.GetGauge().GetValue())
		default:
			return fmt.Errorf("unknown metric type: %s", m.GetType())
		}

		if err := s.store.Update(newMetric.GetName(), newMetric.ToString(), newMetric.GetType()); err != nil {
			s.logger.Info(fmt.Sprintf("err: %s", err.Error()), http.StatusBadRequest)
		}
	}
}

func NewGRPCUpdateHandler(s storage.IStorage, l *zap.SugaredLogger) *grpcUpdateHandler {
	return &grpcUpdateHandler{
		store:  s,
		logger: l,
	}
}
