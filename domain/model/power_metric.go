package model

import (
	"time"
)

type PowerMetric struct {
	SurplusWatt int       `validate:"required"`
	Timestamp   time.Time `validate:"required"`
}

func (m *PowerMetric) Validate() error {
	return Validate(m)
}

type PowerMetricList []*PowerMetric

func (l PowerMetricList) Validate() error {
	for _, metric := range l {
		if err := metric.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (l PowerMetricList) LatestTimestamp() time.Time {
	var latest time.Time
	for _, metric := range l {
		if metric.Timestamp.After(latest) {
			latest = metric.Timestamp
		}
	}
	return latest
}

func (l PowerMetricList) MaximumInterval() time.Duration {
	if len(l) < 2 {
		return 0
	}
	var maximum time.Duration
	for i := 1; i < len(l); i++ {
		interval := l[i].Timestamp.Sub(l[i-1].Timestamp)
		if interval > maximum {
			maximum = interval
		}
	}
	return maximum
}
