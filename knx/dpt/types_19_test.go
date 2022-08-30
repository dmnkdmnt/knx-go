package dpt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ONLY DATE AND TIME MAPPED YET - NO FLAGS

// Test DPT 19.001 (Date Time)
func TestDPT_19001(t *testing.T) {
	knxValue := []byte{0, 122, 8, 7, 240, 55, 42, 0, 0}
	dptValue := DPT_19001{
		DPT_10001: DPT_10001{
			Weekday: 7,
			Hour:    16,
			Minutes: 55,
			Seconds: 42,
		},
		DPT_11001: DPT_11001{
			Year:  2022,
			Month: 8,
			Day:   7,
		},
	}

	assert.True(t, dptValue.IsValid())

	var tmpDPT DPT_19001
	assert.NoError(t, tmpDPT.Unpack(knxValue))
	assert.Equal(t, dptValue, tmpDPT)

	assert.Equal(t, knxValue, dptValue.Pack())

	assert.Equal(t, "Sunday, 07.08.2022, 16:55:42", dptValue.String())
}
