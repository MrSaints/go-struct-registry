package registry

import (
	"reflect"
	"testing"
)

type TestSimple struct{}

func TestRegisterStruct(t *testing.T) {
	t.Run("copy", func(t *testing.T) {
		if err := RegisterStruct("simple_copy", TestSimple{}); err != nil {
			t.Errorf("unexpected error registering struct: %+v", err)
		}
	})

	t.Run("pointer", func(t *testing.T) {
		if err := RegisterStruct("simple_pointer", &TestSimple{}); err != nil {
			t.Errorf("unexpected error registering struct: %+v", err)
		}
	})

	t.Run("zero pointer", func(t *testing.T) {
		if err := RegisterStruct("simple_zero_pointer", (*TestSimple)(nil)); err != nil {
			t.Errorf("unexpected error registering struct: %+v", err)
		}
	})

	t.Run("conflict", func(t *testing.T) {
		if err := RegisterStruct("simple_copy", (*TestSimple)(nil)); err == nil {
			t.Error("expected error registering a conflicting struct")
		}
	})
}

func TestNewStruct(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		st := TestSimple{}
		if err := RegisterStruct("basic", st); err != nil {
			t.Errorf("unexpected error registering struct: %+v", err)
		}

		v, err := NewStruct("basic")
		if err != nil {
			t.Errorf("unexpected error initialising struct: %+v", err)
		}

		got := reflect.TypeOf(v)
		want := reflect.TypeOf(st)

		if got != want {
			t.Errorf("got: %+v, want: %+v", got, want)
		}
	})

	t.Run("non-existent", func(t *testing.T) {
		if _, err := NewStruct("missing"); err == nil {
			t.Error("expected error initialising a non-existent struct")
		}
	})
}
