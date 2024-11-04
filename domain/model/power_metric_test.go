package model_test

import (
	"testing"
	"time"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
)

func TestPowerMetricList_MaximumInterval(t *testing.T) {
	tests := []struct {
		name string
		list model.PowerMetricList
		want time.Duration
	}{
		{
			name: "5 minutes",
			list: model.PowerMetricList{
				{
					Name:      "test",
					Watt:      100,
					Timestamp: time.Date(2024, 1, 1, 0, 5, 0, 0, time.UTC),
				},
				{
					Name:      "test",
					Watt:      100,
					Timestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			want: 5 * time.Minute,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.MaximumInterval(); got != tt.want {
				t.Errorf("MaximumInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
