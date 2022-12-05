package day05

import (
	stack "github.com/golang-ds/stack/slicestack"
	"gotest.tools/v3/assert"
	"testing"
)

func readSampleStackState() []stack.SliceStack[rune] {
	return readStackState([]string{
		`    [D]    `,
		`[N] [C]    `,
		`[Z] [M] [P]`,
		`1   2   3 `,
	})
}

func Test_can_read_stack_state_from_input(t *testing.T) {
	result := readSampleStackState()

	assert.Equal(t, 3, len(result))
	assert.Equal(t, 2, result[0].Size())
	p, _ := result[0].Pop()
	assert.Equal(t, 'N', p)
	p, _ = result[0].Pop()
	assert.Equal(t, 'Z', p)
	assert.Equal(t, 3, result[1].Size())
	p, _ = result[1].Pop()
	assert.Equal(t, 'D', p)
	p, _ = result[1].Pop()
	assert.Equal(t, 'C', p)
	p, _ = result[1].Pop()
	assert.Equal(t, 'M', p)
	assert.Equal(t, 1, result[2].Size())
	p, _ = result[2].Pop()
	assert.Equal(t, 'P', p)
}

func Test_can_read_moves_from_input(t *testing.T) {
	result := readSampleMoves()

	assert.Equal(t, 4, len(result))
	assert.DeepEqual(t, Move{1, 2, 1}, result[0])
	assert.DeepEqual(t, Move{3, 1, 3}, result[1])
	assert.DeepEqual(t, Move{2, 2, 1}, result[2])
	assert.DeepEqual(t, Move{1, 1, 2}, result[3])
}

func readSampleMoves() []Move {
	input := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	return readMoves(input)
}

func Test_applying_first_move_to_initial_state_works(t *testing.T) {
	state := readSampleStackState()

	applyMove(state, Move{1, 2, 1})

	// TODO: simplify this testing code
	assert.Equal(t, 3, len(state))
	assert.Equal(t, 3, state[0].Size())
	p, _ := state[0].Pop()
	assert.Equal(t, 'D', p)
	p, _ = state[0].Pop()
	assert.Equal(t, 'N', p)
	p, _ = state[0].Pop()
	assert.Equal(t, 'Z', p)
	assert.Equal(t, 2, state[1].Size())
	p, _ = state[1].Pop()
	assert.Equal(t, 'C', p)
	p, _ = state[1].Pop()
	assert.Equal(t, 'M', p)
	assert.Equal(t, 1, state[2].Size())
	p, _ = state[2].Pop()
	assert.Equal(t, 'P', p)
}

func Test_moves_applied_to_sample_gives_expected_result(t *testing.T) {
	state := readSampleStackState()
	moves := readSampleMoves()
	result := applyMovesAndGetResultWithCrateMover9000(state, moves)

	assert.Equal(t, "CMZ", result)
}

func Test_moves_applied_to_sample_gives_expected_result_with_crate_mover_9001(t *testing.T) {
	state := readSampleStackState()
	moves := readSampleMoves()
	result := applyMovesAndGetResultWithCrateMover9001(state, moves)

	assert.Equal(t, "MCD", result)
}
