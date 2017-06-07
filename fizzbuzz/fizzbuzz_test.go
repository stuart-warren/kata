package fizzbuzz_test

import (
	"strconv"
	"testing"

	"github.com/stuart-warren/kata/fizzbuzz"
)

func TestNumberShouldNotFizzOrBuzz(t *testing.T) {
	for _, input := range []int{1, 2, 4, 17} {
		output := fizzbuzz.Do(input)
		if strconv.Itoa(input) != output {
			t.Errorf("Got unexpected output: %q from %d", output, input)
		}
	}
}

func TestNumberShouldFizz(t *testing.T) {
	for _, input := range []int{3} {
		output := fizzbuzz.Do(input)
		if output != "fizz" {
			t.Errorf("Got unexpected output: %q from %d", output, input)
		}
	}
}
