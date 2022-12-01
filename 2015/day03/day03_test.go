package day03

import (
	"testing"

	"gotest.tools/assert"
)

func Test_east_only_visits_2(t *testing.T) {
	visited := deliver(">")
	assert.Equal(t, 2, visited)
}

func Test_square_visits_4(t *testing.T) {
	visited := deliver("^>v<")
	assert.Equal(t, 4, visited)
}

func Test_back_and_forth_visits_2(t *testing.T) {
	visited := deliver("^v^v^v^v^v")
	assert.Equal(t, 2, visited)
}

func Test_up_down_with_robo_santa_visits_3(t *testing.T) {
	visited := deliverWithRoboSanta("^v")
	assert.Equal(t, 3, visited)
}

func Test_square_with_robo_santa_visits_3(t *testing.T) {
	visited := deliverWithRoboSanta("^>v<")
	assert.Equal(t, 3, visited)
}

func Test_back_and_for_with_robo_santa_visits_11(t *testing.T) {
	visited := deliverWithRoboSanta("^v^v^v^v^v")
	assert.Equal(t, 11, visited)
}
