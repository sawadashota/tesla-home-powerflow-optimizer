package chargecontroller

import (
	"context"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/event"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/service"
)

type (
	Subscriber interface {
		event.MetricInsertedEventSubscriber
	}
	subscriber struct {
		r dependencies
	}
	dependencies interface {
		service.ChargeServiceProvider
	}
	Provider interface {
		Subscriber() Subscriber
	}
)

func New(r dependencies) Subscriber {
	return &subscriber{r: r}
}

func (s *subscriber) On(_ *event.MetricInsertedEvent) error {
	ctx := context.Background()
	return s.r.ChargeService().Adjust(ctx)
}
