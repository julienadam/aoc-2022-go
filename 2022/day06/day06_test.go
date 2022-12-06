package day06

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Test_can_find_marker_in_mjqjpqmgbljsphdztnvjfqwrcgsmlb(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	assert.Equal(t, 7, findStartOfPacketMarker(input))
}

func Test_can_find_marker_in_bvwbjplbgvbhsrlpgdmjqwftvncz(t *testing.T) {
	input := "bvwbjplbgvbhsrlpgdmjqwftvncz"

	assert.Equal(t, 5, findStartOfPacketMarker(input))
}

func Test_can_find_marker_in_nppdvjthqldpwncqszvftbrmjlhg(t *testing.T) {
	input := "nppdvjthqldpwncqszvftbrmjlhg"

	assert.Equal(t, 6, findStartOfPacketMarker(input))
}

func Test_can_find_marker_in_nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg(t *testing.T) {
	input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"

	assert.Equal(t, 10, findStartOfPacketMarker(input))
}

func Test_can_find_marker_in_zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw(t *testing.T) {
	input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

	assert.Equal(t, 11, findStartOfPacketMarker(input))
}

func Test_can_find_start_of_message_marker_in_mjqjpqmgbljsphdztnvjfqwrcgsmlb(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	assert.Equal(t, 19, findStartOfMessageMarker(input))
}

func Test_can_find_start_of_message_marker_in_bvwbjplbgvbhsrlpgdmjqwftvncz(t *testing.T) {
	input := "bvwbjplbgvbhsrlpgdmjqwftvncz"

	assert.Equal(t, 23, findStartOfMessageMarker(input))
}

func Test_can_find_start_of_message_marker_in_nppdvjthqldpwncqszvftbrmjlhg(t *testing.T) {
	input := "nppdvjthqldpwncqszvftbrmjlhg"

	assert.Equal(t, 23, findStartOfMessageMarker(input))
}

func Test_can_find_start_of_message_marker_in_nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg(t *testing.T) {
	input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"

	assert.Equal(t, 29, findStartOfMessageMarker(input))
}

func Test_can_find_start_of_message_marker_in_zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw(t *testing.T) {
	input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

	assert.Equal(t, 26, findStartOfMessageMarker(input))
}
