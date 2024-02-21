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

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")

	// 表示已有所添加key
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if want != got {
		t.Errorf("got '%s' want '%s'", got, want)
	}

}

func assertStrings(t *testing.T, got, want error) {
	t.Helper()

	if !errors.Is(want, got) {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
