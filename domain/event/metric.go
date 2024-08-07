package event

type (
	Publisher[T any] interface {
		Publish(e *T) error
	}
	Subscriber[T any] interface {
		On(e *T) error
	}
)

type (
	MetricInsertedEvent                  struct{}
	MetricInsertedEventPublisher         Publisher[MetricInsertedEvent]
	MetricInsertedEventPublisherProvider interface {
		MetricInsertedEventPublisher() MetricInsertedEventPublisher
	}

	MetricInsertedEventSubscriber         Subscriber[MetricInsertedEvent]
	MetricInsertedEventSubscriberProvider interface {
		MetricInsertedEventSubscriber() MetricInsertedEventSubscriber
	}
)
