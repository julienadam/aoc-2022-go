package puzzleLoader

import (
	"fmt"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

const TEST_YEAR = 1900
const TEST_DAY = 1

func Test_file_name_is_correct_for_real_data(t *testing.T) {
	expected := fmt.Sprintf("1900%cday01%cdata%cday01.txt", os.PathSeparator, os.PathSeparator, os.PathSeparator)
	assert.Equal(t, expected, getFileName(TEST_YEAR, TEST_DAY, Real))
}

func Test_file_name_is_correct_for_sample_data1(t *testing.T) {
	expected := fmt.Sprintf("1900%cday01%cdata%cday01_sample1.txt", os.PathSeparator, os.PathSeparator, os.PathSeparator)
	assert.Equal(t, expected, getFileName(TEST_YEAR, TEST_DAY, Sample1))
}

func Test_file_name_is_correct_for_sample_data2(t *testing.T) {
	expected := fmt.Sprintf("1900%cday01%cdata%cday01_sample2.txt", os.PathSeparator, os.PathSeparator, os.PathSeparator)
	assert.Equal(t, expected, getFileName(TEST_YEAR, TEST_DAY, Sample2))
}

func Test_can_load_all_lines_from_file(t *testing.T) {
	got, err := ReadAllLines("sample.txt")
	assert.NilError(t, err)
	assert.DeepEqual(t, got, []string{"blah", "foo"})
}

func Test_can_load_all_text_from_file(t *testing.T) {
	got, err := ReadAllText("sample.txt")
	assert.NilError(t, err)
	assert.DeepEqual(t, got, "blah\r\nfoo")
}
