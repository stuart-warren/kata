package fizzbuzz_test

import (
	"strconv"
	"testing"

	"github.com/stuart-warren/kata/fizzbuzz"
)

func TestNumberShouldNotFizzOrBuzz(t *testing.T) {
	for input := range []int{1, 2, 4, 17} {
		output, err := fizzbuzz.Do(input)
		if err != nil {
			t.Errorf("Got unexpected error: %v", err)
		}
		if strconv.Itoa(input) != output {
			t.Errorf("Got unexpected output: %s from %d", output, input)
		}
	}
}
