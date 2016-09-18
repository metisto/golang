package rpn

import "testing"

func Test_rpn_should_return_single_number_when_input_is_a_single_number(t *testing.T) {
	assert(t, 7, Rpn("7"))
	assert(t, 2.8, Rpn("2.8"))
}

func Test_rpn_should_return_result_of_simple_operation(t *testing.T) {
	assert(t, 1, Rpn("0.2 0.8 +"))
	assert(t, 2, Rpn("3 1 -"))
	assert(t, 3, Rpn("9 3 /"))
	assert(t, 4, Rpn("2 2 *"))
}

func assert(t *testing.T, expected float64, actual float64) {
	if actual != expected {
		t.Fatalf("Expected %f but actual is %f", expected, actual)
	}
}
