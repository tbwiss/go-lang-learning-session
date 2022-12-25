package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Tom", "")
		want := "Hello, Tom"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Rant", "Spanish")
		want := "Hola, Rant"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
