package worker

import (
	"github.com/mattn/go-pubsub"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/event"
)

type (
	PubSub struct {
		ps *pubsub.PubSub
		r  pubSubDependencies

		metricInsertedEventPublisher event.MetricInsertedEventPublisher
	}
	pubSubDependencies interface {
		event.MetricInsertedEventSubscriberProvider
	}
	PubSubProvider interface {
		event.MetricInsertedEventPublisherProvider
	}
)

func NewPubSub(r pubSubDependencies) PubSubProvider {
	ps := pubsub.New()
	return &PubSub{
		ps: ps,
		r:  r,
		metricInsertedEventPublisher: publisher[event.MetricInsertedEvent](func(e *event.MetricInsertedEvent) error {
			if err := ps.Sub(r.MetricInsertedEventSubscriber().On); err != nil {
				return err
			}
			ps.Pub(e)
			return nil
		}),
	}
}

func (p *PubSub) MetricInsertedEventPublisher() event.MetricInsertedEventPublisher {
	return p.metricInsertedEventPublisher
}

type publisher[T any] func(e *T) error

func (p publisher[T]) Publish(e *T) error {
	return p(e)
}
