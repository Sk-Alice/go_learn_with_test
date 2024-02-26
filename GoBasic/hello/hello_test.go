package main

// https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/hello-world

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// t.Helper() 方法通常用于标记一个测试辅助函数，以指示这个函数不是真正的测试函数，而是被测试函数所调用的辅助函数。
		// 通常在辅助函数中的第一行调用，它告诉测试框架将当前函数标记为辅助函数，而不是真正的测试函数。
		// 从而能够正确地报告测试失败时的文件名和行号
		t.Helper()

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("K", "French")
		want := "Bonjour, K"

		assertCorrectMessage(t, got, want)
	})

}
