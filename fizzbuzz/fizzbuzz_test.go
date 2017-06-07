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
	for _, input := range []int{3, 6, 9, 99} {
		output := fizzbuzz.Do(input)
		if output != "fizz" {
			t.Errorf("Got unexpected output: %q from %d", output, input)
		}
	}
}

func TestNumberShouldBuzz(t *testing.T) {
	for _, input := range []int{5, 10, 2000} {
		output := fizzbuzz.Do(input)
		if output != "buzz" {
			t.Errorf("Got unexpected output: %q from %d", output, input)
		}
	}
}

func TestNumberShouldFizzBuzz(t *testing.T) {
	for _, input := range []int{15, 150} {
		output := fizzbuzz.Do(input)
		if output != "fizzbuzz" {
			t.Errorf("Got unexpected output: %q from %d", output, input)
		}
	}
}
