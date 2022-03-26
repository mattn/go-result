package result

import (
	"errors"
	"reflect"
	"testing"
)

func doSomething1() Result1[bool] {
	return Ok1(true)
}

func doSomething2() Result1[int] {
	return Err1[int](errors.New("something wrong"))
}

func doSomething3() Result2[int, bool] {
	return Ok2(123, true)
}

func TestUnwrapOk(t *testing.T) {
	r := doSomething1()
	v := r.Unwrap()
	if reflect.TypeOf(v).Kind() != reflect.Bool {
		t.Fatalf("Unwrap must return bool but got %T", v)
	}

	if r.IsErr() == true {
		t.Fatal("IsErr must return false")
	}

	if r.Err() != nil {
		t.Fatal("Err must return nil")
	}
}

func TestUnwrapErr(t *testing.T) {
	r := doSomething2()
	var err error
	func() {
		defer func() {
			err = recover().(error)
		}()
		r.Unwrap()
	}()
	if err == nil {
		t.Fatal("Unwrap must fail")
	}

	if r.IsErr() == false {
		t.Fatal("IsErr must return true")
	}

	if r.Err() == nil {
		t.Fatal("Err must return non-nil")
	}
}

func TestResult2(t *testing.T) {
	r := doSomething3()
	v1, v2 := r.Unwrap()
	if reflect.TypeOf(v1).Kind() != reflect.Int || reflect.TypeOf(v2).Kind() != reflect.Bool {
		t.Fatalf("Unwrap must return int, bool but got %T, %T", v1, v2)
	}

	if r.IsErr() == true {
		t.Fatal("IsErr must return false")
	}

	if r.Err() != nil {
		t.Fatal("Err must return nil")
	}

	v1 = r.V1()
	if reflect.TypeOf(r.V1()).Kind() != reflect.Int {
		t.Fatalf("Unwrap must return int, bool but got %T, %T", v1, v2)
	}
	v2 = r.V2()
	if reflect.TypeOf(v2).Kind() != reflect.Bool {
		t.Fatalf("Unwrap must return int, bool but got %T, %T", v1, v2)
	}
}
