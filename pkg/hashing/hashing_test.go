package hashing

import (
	"reflect"
	"testing"
)

func TestMakePass(t *testing.T) {

	t.Run("with valid password", func(t *testing.T) {
		got, _ := MakePass("password")
		want := 60

		if len(got) != want {
			t.Errorf("got %d want %d given, %v", len(got), want, "password")
		}
	})

	t.Run("with empty password", func(t *testing.T) {
		got, err := MakePass("")
		want := "Empty name"

		if err.Error() != want {
			t.Errorf("got %s want %s given, %v", err, want, got)
		}
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Run("with valid password", func(t *testing.T) {
		hesh := "$2a$14$ZOubd0goKj9Dhfkgd3GsPOZfAHkvGG/0ih8zkSx0.bI1JmbJljSNe"
		pass := "1"
		got := CheckPasswordHash(pass, hesh)
		var want bool

		if reflect.TypeOf(got) != reflect.TypeOf(want) {
			t.Errorf("got %t want %t given, %v", got, want, "password")
		}
	})
}
