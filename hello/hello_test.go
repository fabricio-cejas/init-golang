package hello

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello, world."

    if result := Hello(); result != expected {
        t.Errorf("Hello() = %q, expected = %q", result, expected)
    }
}

func TestProverb(t *testing.T) {
    want := "Concurrency is not parallelism."
    if got := Proverb(); got != want {
        t.Errorf("Proverb() = %q, want %q", got, want)
    }
}
