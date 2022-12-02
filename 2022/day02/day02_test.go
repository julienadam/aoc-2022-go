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
	input := []string{"A Y", "B X", "C Z"}
	assert.Equal(t, 15, scoreAllRounds(input))
}

func Test_can_parse_advanced_instructions_rock_loss(t *testing.T) {
	o, s := parseRoundAdvanced("A X")
	assert.Equal(t, rock, o)
	assert.Equal(t, loss, s)
}

func Test_can_parse_advanced_instructions_paper_draw(t *testing.T) {
	o, s := parseRoundAdvanced("B Y")
	assert.Equal(t, paper, o)
	assert.Equal(t, draw, s)
}

func Test_can_parse_advanced_instructions_scissors_win(t *testing.T) {
	o, s := parseRoundAdvanced("C Z")
	assert.Equal(t, scissors, o)
	assert.Equal(t, win, s)
}

func Test_rock_and_draw_gives_4_points(t *testing.T) {
	score := scoreForOpponentAndAdvancedStrategy(rock, draw)
	assert.Equal(t, 4, score)
}

func Test_paper_and_loss_gives_1_points(t *testing.T) {
	score := scoreForOpponentAndAdvancedStrategy(paper, loss)
	assert.Equal(t, 1, score)
}

func Test_scissors_and_win_gives_7_points(t *testing.T) {
	score := scoreForOpponentAndAdvancedStrategy(scissors, win)
	assert.Equal(t, 7, score)
}


func Test_sample_game_with_advanced_strategy_gives_12_points(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}
	assert.Equal(t, 12, scoreAllRoundsAdvanced(input))
}