package dictionary

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	/*	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})*/

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknow")

		assertStrings(t, got, ErrNotFound)
	})
}

func assertStrings(t *testing.T, got, want error) {
	t.Helper()

	if !errors.Is(want, got) {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
