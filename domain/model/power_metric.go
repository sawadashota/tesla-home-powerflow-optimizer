package model

import (
	"cmp"
	"slices"
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

func (l PowerMetricList) IsSorted() bool {
	return slices.IsSortedFunc(l, func(i, j *PowerMetric) int {
		return cmp.Compare(i.Timestamp.UnixNano(), j.Timestamp.UnixNano())
	})
}

func (l PowerMetricList) Sort() PowerMetricList {
	if l.IsSorted() {
		return l
	}
	sorted := make(PowerMetricList, len(l))
	for i, metric := range l {
		sorted[i] = metric
	}
	slices.SortFunc(sorted, func(i, j *PowerMetric) int {
		return cmp.Compare(i.Timestamp.UnixNano(), j.Timestamp.UnixNano())
	})
	return sorted
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
	sorted := l.Sort()
	var maximum time.Duration
	for i := 1; i < len(sorted); i++ {
		interval := sorted[i-1].Timestamp.Sub(sorted[i].Timestamp)
		if interval > maximum {
			maximum = interval
		}
	}
	return maximum
}
