package goldcrest_test

import (
	"testing"

	"github.com/Nigel2392/goldcrest"
)

func TestHooks(t *testing.T) {
	type MyIntFunc func([]int) []int

	goldcrest.Register(
		"construct_int_list", -1,
		MyIntFunc(func(list []int) []int {
			return append(list, 1)
		}),
	)
	var myIntList = make([]int, 0)
	myIntList = append(myIntList, 1, 2, 3)

	var hookList = goldcrest.Get[MyIntFunc]("construct_int_list")
	for _, hook := range hookList {
		myIntList = hook(myIntList)
	}

	if len(myIntList) != 4 {
		t.Errorf("Expected length of 4, got %d", len(myIntList))
	}

	if myIntList[3] != 1 {
		t.Errorf("Expected 1, got %d", myIntList[3])
	}
}

func TestHooksUnregister(t *testing.T) {
	type TestFunc func()

	goldcrest.Register(
		"test_unregister", -1,
		TestFunc(func() {}),
	)

	var hookList = goldcrest.Get[TestFunc]("test_unregister")
	if len(hookList) != 1 {
		t.Errorf("Expected length of 1, got %d", len(hookList))
	}

	goldcrest.Unregister("test_unregister")

	hookList = goldcrest.Get[TestFunc]("test_unregister")
	if len(hookList) != 0 {
		t.Errorf("Expected length of 0, got %d", len(hookList))
	}
}

func TestHooksTypeCast(t *testing.T) {
	type TestFunc func()

	goldcrest.Register(
		"test_type_cast", -1,
		func() {},
	)

	var hookList = goldcrest.Get[TestFunc]("test_type_cast")
	if len(hookList) != 1 {
		t.Errorf("Expected length of 1, got %d", len(hookList))
	}

	var hooklist = goldcrest.Get[func()]("test_type_cast")
	if len(hooklist) != 1 {
		t.Errorf("Expected length of 1, got %d", len(hooklist))
	}
}

func TestHooksTypeCastPanic(t *testing.T) {

	type TestFunc func()

	goldcrest.Register(
		"test_type_cast", -1,
		TestFunc(func() {}),
	)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected no panic, got %v", r)
		} else {
			t.Logf("Expected and caught panic: %v", r)
		}
	}()

	var _ = goldcrest.Get[func() int]("test_type_cast")

}
