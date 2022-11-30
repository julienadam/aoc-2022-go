package puzzleLoader

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_can_load_real_data(t *testing.T) {

	got, err := LoadLines(2022, 0, Real)
	assert.NilError(t, err)
	assert.DeepEqual(t, got, []string{"blah", "foo"})
}

func Test_can_load_sample1_data(t *testing.T) {
	got, err := LoadLines(2022, 0, Sample1)
	assert.NilError(t, err)
	assert.DeepEqual(t, got, []string{"blah1", "foo1"})
}

func Test_can_load_sample2_data(t *testing.T) {
	got, err := LoadLines(2022, 0, Sample2)
	assert.NilError(t, err)
	assert.DeepEqual(t, got, []string{"blah2", "foo2"})
}

func Test_can_load_real_data_as_text(t *testing.T) {

	got, err := LoadString(2022, 0, Real)
	assert.NilError(t, err)
	assert.DeepEqual(t, got, "blah\r\nfoo")
}
