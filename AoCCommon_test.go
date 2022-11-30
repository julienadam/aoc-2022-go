package AoCCommon

import (
	"fmt"
	"testing"
)

func Test_can_load_real_data(t *testing.T) {

	got, err := LoadData(0, Real)
	if err != nil {
		t.Errorf("failed to open file for day %d", 0)
	}

	for i, e := range got {
		fmt.Println(e)
		if (i == 0) && (e != "blah") {
			t.Errorf("first line should have been `blah`, was %s", e)
		} else if (i == 1) && (e != "foo") {
			t.Errorf("second line should have been `foo`, was %s", e)
		} else if i > 1 {
			t.Errorf("Invalid value %s at line %d, should not have read a value", e, i)
		}

		i++
	}
}

func Test_can_load_sample1_data(t *testing.T) {

	got, err := LoadData(0, Sample1)
	if err != nil {
		t.Errorf("failed to open file for day %d", 0)
	}

	for i, e := range got {
		fmt.Println(e)
		if (i == 0) && (e != "blah1") {
			t.Errorf("first line should have been `blah`, was %s", e)
		} else if (i == 1) && (e != "foo1") {
			t.Errorf("second line should have been `foo`, was %s", e)
		} else if i > 1 {
			t.Errorf("Invalid value %s at line %d, should not have read a value", e, i)
		}

		i++
	}
}

func Test_can_load_sample2_data(t *testing.T) {

	got, err := LoadData(0, Sample2)
	if err != nil {
		t.Errorf("failed to open file for day %d", 0)
	}

	for i, e := range got {
		fmt.Println(e)
		if (i == 0) && (e != "blah2") {
			t.Errorf("first line should have been `blah`, was %s", e)
		} else if (i == 1) && (e != "foo2") {
			t.Errorf("second line should have been `foo`, was %s", e)
		} else if i > 1 {
			t.Errorf("Invalid value %s at line %d, should not have read a value", e, i)
		}

		i++
	}
}
