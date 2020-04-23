package fib

import "testing"

//Fib Sequence: 1, 1, 2, 3, 5, 8, 13, 21, 34

/*
//TestFib is a first iteration with an attemp to test error
//It's possible to assert errors by calling err.Error() to get error's message,
//but that's not a good idea because err could be nil and call of err.Error()
//will lead to runtime error, so an additional check for nill case is required.
//The solution for this problem is described below
func TestFib(t *testing.T) {
	res, err := Fib(7)
	if res != 13 || err != nil {
		t.Errorf("Fib(7) returns %v, %v; expected %v, %v", res, err, 13, nil)
	}

	res, err = Fib(101)
	if res != 0 || err != errors.New("bad input") {
		t.Errorf("Fib(101) returns %v, %v; expected %v, %v", res, err, 0, errors.New("bad input"))
	}
}
*/

/*
//TestErrFixedFib is a second iteration when constant like globals were introduced.
//That's much easier to check/verify errors if they are defined in global vars. We could rely
//on their values not just in tests but in our application, please refer to the example below when
//there is a special branch in the app when there are no rows found
// row := db.QueryRow(sqlStatement, 3)
// switch err := row.Scan(&id, &email); err {
//  case sql.ErrNoRows:
//    fmt.Println("No rows were returned!")
//  case nil:
//    fmt.Println(id, email)
//  default:
//    panic(err)
//  }
func TestErrFixedFib(t *testing.T) {
	res, err := Fib(7)
	if res != 13 || err != nil {
		t.Errorf("Fib(7) returns %v, %v; expected %v, %v", res, err, 13, nil)
	}

	res, err = Fib(101)
	if res != 0 || err != ErrInputBig {
		t.Errorf("Fib(101) returns %v, %v; expected %v, %v", res, err, 0, ErrInputBig)
	}
}*/

//TestTDFib is a 3rd iteration when table driven tests approach is demonstrated.
//This approach allows to add test cases easilly and makes tests code clean/without copy-paste
func TestTDFib(t *testing.T) {
	testCases := []struct {
		in  int64
		out int64
		err error
	}{
		{-1, 0, ErrInputSmall},
		{101, 0, ErrInputBig},
		{0, 0, nil},
		{1, 1, nil},
		{7, 13, nil},
	}

	for _, c := range testCases {
		res, err := Fib(c.in)
		if res != c.out || err != c.err {
			t.Errorf("Fib(%d) returns %v, %v; expected %v, %v", c.in, res, err, c.out, c.err)
		}
	}
}

//TestTDFibImproved is the same as test above but it tests FibImproved function
//In real life FibImproved should be Fib with new implementation and TestTDFib should be used
//to unit test it. This demonstrates how it's usefull to have unit tests as
//they could be reused to check new implementation.
func TestTDFibImproved(t *testing.T) {
	testCases := []struct {
		in  int64
		out int64
		err error
	}{
		{-1, 0, ErrInputSmall},
		{101, 0, ErrInputBig},
		{0, 0, nil},
		{1, 1, nil},
		{7, 13, nil},
	}

	for _, c := range testCases {
		res, err := FibImproved(c.in)
		if res != c.out || err != c.err {
			t.Errorf("Fib(%d) returns %v, %v; expected %v, %v", c.in, res, err, c.out, c.err)
		}
	}
}
