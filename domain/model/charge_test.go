package model_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
)

func TestTimeRange_InRange(t *testing.T) {
	tests := []struct {
		tr     model.TimeRange
		given  model.TimeOnly
		assert assert.BoolAssertionFunc
	}{
		{
			tr: model.TimeRange{
				Start: model.NewTimeOnly(21, 0, 0),
				End:   model.NewTimeOnly(6, 0, 0),
			},
			given:  model.NewTimeOnly(21, 0, 0),
			assert: assert.True,
		},
		{
			tr: model.TimeRange{
				Start: model.NewTimeOnly(21, 0, 0),
				End:   model.NewTimeOnly(6, 0, 0),
			},
			given:  model.NewTimeOnly(6, 0, 0),
			assert: assert.True,
		},
		{
			tr: model.TimeRange{
				Start: model.NewTimeOnly(21, 0, 0),
				End:   model.NewTimeOnly(6, 0, 0),
			},
			given:  model.NewTimeOnly(20, 59, 59),
			assert: assert.False,
		},
		{
			tr: model.TimeRange{
				Start: model.NewTimeOnly(21, 0, 0),
				End:   model.NewTimeOnly(6, 0, 0),
			},
			given:  model.NewTimeOnly(6, 0, 1),
			assert: assert.False,
		},
		{
			tr: model.TimeRange{
				Start: model.NewTimeOnly(6, 0, 0),
				End:   model.NewTimeOnly(21, 0, 0),
			},
			given:  model.NewTimeOnly(6, 0, 0),
			assert: assert.True,
		},
		{
			tr: model.TimeRange{
				Start: model.NewTimeOnly(6, 0, 0),
				End:   model.NewTimeOnly(21, 0, 0),
			},
			given:  model.NewTimeOnly(21, 0, 0),
			assert: assert.True,
		},
		{
			tr: model.TimeRange{
				Start: model.NewTimeOnly(6, 0, 0),
				End:   model.NewTimeOnly(21, 0, 0),
			},
			given:  model.NewTimeOnly(5, 59, 59),
			assert: assert.False,
		},
		{
			tr: model.TimeRange{
				Start: model.NewTimeOnly(6, 0, 0),
				End:   model.NewTimeOnly(21, 0, 0),
			},
			given:  model.NewTimeOnly(21, 0, 1),
			assert: assert.False,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("start:%s, end:%s, given:%s", tt.tr.Start, tt.tr.End, tt.given), func(t *testing.T) {
			tt.assert(t, tt.tr.InRange(tt.given))
		})
	}
}
