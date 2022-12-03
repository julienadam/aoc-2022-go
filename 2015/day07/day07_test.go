package day07

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_signal_for_an_assignment_gives_the_signal_value(t *testing.T) {
	layout := make(map[string]string)
	layout["x"] = "123"
	signal := findSignalForWire("x", layout, make(map[string]uint16))
	assert.Equal(t, uint16(123), signal)
}

func Test_signal_with_assignment_into_not(t *testing.T) {
	layout := make(map[string]string)
	layout["y"] = "123"
	layout["x"] = "NOT y"
	signal := findSignalForWire("x", layout, make(map[string]uint16))
	assert.Equal(t, uint16(65412), signal)
}

func Test_signal_with_assignment_into_lshift_2(t *testing.T) {
	layout := make(map[string]string)
	layout["y"] = "123"
	layout["x"] = "y LSHIFT 2"
	signal := findSignalForWire("x", layout, make(map[string]uint16))
	assert.Equal(t, uint16(492), signal)
}

func Test_signal_with_assignment_into_rshift_2(t *testing.T) {
	layout := make(map[string]string)
	layout["y"] = "123"
	layout["x"] = "y RSHIFT 2"
	signal := findSignalForWire("x", layout, make(map[string]uint16))
	assert.Equal(t, uint16(30), signal)
}

func Test_signal_with_assignments_for_x_y_into_AND(t *testing.T) {
	layout := make(map[string]string)
	layout["x"] = "123"
	layout["y"] = "456"
	layout["d"] = "x AND y"
	signal := findSignalForWire("d", layout, make(map[string]uint16))
	assert.Equal(t, uint16(72), signal)
}

func Test_signal_with_assignments_for_x_y_into_OR(t *testing.T) {
	layout := make(map[string]string)
	layout["x"] = "123"
	layout["y"] = "456"
	layout["d"] = "x OR y"
	signal := findSignalForWire("d", layout, make(map[string]uint16))
	assert.Equal(t, uint16(507), signal)
}


func Test_signal_with_chained_assignments_follows_chain(t *testing.T) {
	layout := make(map[string]string)
	layout["x"] = "y"
	layout["y"] = "456"
	signal := findSignalForWire("x", layout, make(map[string]uint16))
	assert.Equal(t, uint16(456), signal)
}

func Test_sample_gives_expected_answers(t *testing.T) {
	layout := map[string]string{
		"x": "123",
		"y": "456",
		"d": "x AND y",
		"e": "x OR y",
		"f": "x LSHIFT 2",
		"g": "y RSHIFT 2",
		"h": "NOT x",
		"i": "NOT y",
	}

	assert.Equal(t, uint16(72), findSignalForWire("d", layout, make(map[string]uint16)))
	assert.Equal(t, uint16(507), findSignalForWire("e", layout, make(map[string]uint16)))
	assert.Equal(t, uint16(492), findSignalForWire("f", layout, make(map[string]uint16)))
	assert.Equal(t, uint16(114), findSignalForWire("g", layout, make(map[string]uint16)))
	assert.Equal(t, uint16(65412), findSignalForWire("h", layout, make(map[string]uint16)))
	assert.Equal(t, uint16(65079), findSignalForWire("i", layout, make(map[string]uint16)))
	assert.Equal(t, uint16(123), findSignalForWire("x", layout, make(map[string]uint16)))
	assert.Equal(t, uint16(456), findSignalForWire("y", layout, make(map[string]uint16)))
}

