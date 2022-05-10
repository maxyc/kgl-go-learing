package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assert := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}

	t.Run("with name", func(t *testing.T) {
		got := hello(name, langEnglish)
		want := EnglishHelloText + " " + name

		assert(t, got, want)
	})

	t.Run("without name", func(t *testing.T) {
		got := hello("", langEnglish)
		want := EnglishHelloText + " world"

		assert(t, got, want)
	})

	t.Run("test spanish with name", func(t *testing.T) {
		got := hello("maxyc", langSpanish)
		want := SpanishHelloText + " " + name

		assert(t, got, want)
	})

	t.Run("test franch with name", func(t *testing.T) {
		got := hello("maxyc", langFranch)
		want := FranchHelloText + " " + name

		assert(t, got, want)
	})
}
