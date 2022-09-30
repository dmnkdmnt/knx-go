package dpt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test DPT 5.001 (Scaling) with values within range
func TestDPT_5001(t *testing.T) {
	knxValue := []byte{0, 107}
	dptValue := DPT_5001(42)

	var tmpDPT DPT_5001
	assert.NoError(t, tmpDPT.Unpack(knxValue))
	assert.Equal(t, dptValue, tmpDPT)

	assert.Equal(t, knxValue, dptValue.Pack())

	assert.Equal(t, "42%", dptValue.String())
}

// Test DPT 5.003 (Angle) with values within range
func TestDPT_5003(t *testing.T) {
	knxValue := []byte{0, 30}
	dptValue := DPT_5003(42)

	var tmpDPT DPT_5003
	assert.NoError(t, tmpDPT.Unpack(knxValue))
	assert.Equal(t, dptValue, tmpDPT)

	assert.Equal(t, knxValue, dptValue.Pack())

	assert.Equal(t, "42°", dptValue.String())
}

func TestDPT_5005(t *testing.T) {
	knxValue := []byte{0, 42}
	dptValue := DPT_5005(42)

	var tmpDPT DPT_5005
	assert.NoError(t, tmpDPT.Unpack(knxValue))
	assert.Equal(t, dptValue, tmpDPT)

	assert.Equal(t, knxValue, dptValue.Pack())

	assert.Equal(t, "42", dptValue.String())
}
