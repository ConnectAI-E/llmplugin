package calculator

import (
	"context"
	"testing"
)

func TestCalculator_Do(t *testing.T) {
	ts := []struct {
		testname string
		query    string
		expect   string
	}{
		{
			"Add",
			"1+1",
			"2",
		},
		{
			"Sub",
			"1-1",
			"0",
		},
		{
			"Mul",
			"1*1",
			"1",
		},
	}
	for _, tt := range ts {
		t.Run(tt.testname, func(t *testing.T) {
			c := NewCalculator()
			got, err := c.Do(context.Background(), tt.query)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.expect {
				t.Errorf("Calculator.Do() = %v, want %v", got, tt.expect)
			}
		})
	}
}
