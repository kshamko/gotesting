package fib

import "errors"

// ErrInputSmall is a error definition for bad input for fib func
var ErrInputSmall = errors.New("bad input. too small")

// ErrInputBig is a error definition for bad input for fib func
var ErrInputBig = errors.New("bad input. too big")

//Fib function takes number of fid sequence element and returns value of the element or error
//Recursive approach is used here. It's elegant and looks beautiful but uses a lot of redundunt computations
//and could lead to stack overflow error if case of relatively big input (please google about it)
func Fib(a int64) (int64, error) {
	if a < 0 {

		//line below was used in the first iteration
		//return 0, errors.New("bad input. too small")

		return 0, ErrInputSmall
	}

	if a > 100 {

		//line below was used in the first iteration
		//return 0, errors.New("bad input. too big")

		return 0, ErrInputBig
	}
	return fib(a), nil
}

func fib(a int64) int64 {
	if a == 0 || a == 1 {
		return a
	}

	return fib(a-2) + fib(a-1)
}

//FibImproved function takes number of fid sequence element and returns value of the element or error
//Tail-recursive approach is used here. Tail recursion uses accumulated values from the previous step to
//make a calculation, could be replaced with a loop (please read about that)
func FibImproved(a int64) (int64, error) {
	if a < 0 {
		return 0, ErrInputSmall
	}

	if a > 100 {
		return 0, ErrInputBig
	}

	return fibTail(a, 0, 1), nil
}

func fibTail(a int64, first, second int64) int64 {
	if a == 0 {
		return first
	}
	return fibTail(a-1, second, first+second)
}
