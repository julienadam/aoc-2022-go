package day02

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_against_rock_paper_wins_and_gives_8_points(t *testing.T) {
	assert.Equal(t, 8, scoreRound(charToRps('A'), charToRps('Y')))
}

func Test_against_rocks_scissors_loses_and_gives_3_point(t *testing.T) {
	assert.Equal(t, 3, scoreRound(charToRps('X'), charToRps('C')))
}

func Test_against_paper_paper_draws_and_gives_5_points(t *testing.T) {
	assert.Equal(t, 5, scoreRound(charToRps('Y'), charToRps('B')))
}

func Test_can_parse_round(t *testing.T) {
	o, m := parseRound("A X")
	assert.Equal(t, rock, o)
	assert.Equal(t, rock, m)
}

func Test_sample_game_gives_15_points(t *testing.T) {
	input := []string { "A Y", "B X", "C Z"}
	assert.Equal(t, 15, scoreAllRounds(input))
}
