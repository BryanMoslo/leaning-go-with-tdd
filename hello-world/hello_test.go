package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q and want: %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Santiago", "")
		want := "Hello, Santiago"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world as default", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Moslo", "Spanish")
		want := "Hola, Moslo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in french", func(t *testing.T) {
		got := Hello("Moslo", "French")
		want := "Bonjour, Moslo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in portuguese", func(t *testing.T) {
		got := Hello("Bryan", "Portuguese")
		want := "Olá, Bryan"
		assertCorrectMessage(t, got, want)
	})
}
