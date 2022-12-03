package day08

import (
	"testing"

	"gotest.tools/v3/assert"
)

const TestEmpty = "\"\""
const TestAbc = "\"abc\""
const TestAaaAaa = "\"aaa\\\"aaa\""
const TestX27 = "\"\\x27\""

func Test_empty_counts_as_0(t *testing.T) {
	charsInMemory := CountCharsInMemory(TestEmpty)
	assert.Equal(t, 0, charsInMemory)
}

func Test_abc_counts_as_3(t *testing.T) {
	charsInMemory := CountCharsInMemory(TestAbc)
	assert.Equal(t, 3, charsInMemory)
}

func Test_escaped_quote_counts_as_1_char(t *testing.T) {
	charsInMemory := CountCharsInMemory(TestAaaAaa)
	assert.Equal(t, 7, charsInMemory)
}

func Test_escaped_unicode_chat_counts_as_1_char(t *testing.T) {
	charsInMemory := CountCharsInMemory(TestX27)
	assert.Equal(t, 1, charsInMemory)
}

func Test_chars_minus_memory_for_sample_is_12(t *testing.T) {
	input := []string{
		TestEmpty,
		TestAbc,
		TestAaaAaa,
		TestX27,
	}

	got := TotalCharsMinusTotalMemory(input)
	assert.Equal(t, 12, got)
}

func Test_encoding_empty_string_gives_6_chars_escaped(t *testing.T) {
	result := Escape(TestEmpty)
	assert.Equal(t, "\"\\\"\\\"\"", result)
	assert.Equal(t, 6, len(result))
}

func Test_encoding_abc_gives_9_chars_escaped(t *testing.T) {
	result := Escape(TestAbc)
	assert.Equal(t, "\"\\\"abc\\\"\"", result)
	assert.Equal(t, 9, len(result))
}

func Test_encoding_aaa_aaa_gives_16_chars_escaped(t *testing.T) {
	result := Escape(TestAaaAaa)
	assert.Equal(t, "\"\\\"aaa\\\\\\\"aaa\\\"\"", result)
	assert.Equal(t, 16, len(result))
}

func Test_encoding_x27_gives_11_chars_escaped(t *testing.T) {
	result := Escape(TestX27)
	assert.Equal(t, "\"\\\"\\\\x27\\\"\"", result)
	assert.Equal(t, 11, len(result))
}

func Test_encoded_chars_minus_input_chars_for_sample_is_19(t *testing.T) {
	input := []string{
		TestEmpty,
		TestAbc,
		TestAaaAaa,
		TestX27,
	}

	got := TotalEncodedCharsMinusInitialChars(input)
	assert.Equal(t, 19, got)
}
