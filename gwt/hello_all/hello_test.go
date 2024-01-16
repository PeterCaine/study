package main

import "testing"

func TestHello(t *testing.T){
  t.Run("saying hello", func(t *testing.T){
    got := Hello("Chris", "")
    want := "Hello, Chris"
    assertCorrectMessage(t, got, want)
  })
  t.Run("saying hello default", func(t *testing.T){
    got := Hello("", "")
    want := "Hello, World"
    assertCorrectMessage(t, got, want)
  })
  t.Run("saying hello in Spanish", func(t *testing.T){
    got := Hello("", "Spanish")
    want := "Hola, World"
    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
