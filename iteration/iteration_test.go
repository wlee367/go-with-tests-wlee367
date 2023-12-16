package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Basic Case", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Caller specifies the number of repetition", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"

		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}

func ExampleRepeat() {
	result := Repeat("Jason", 2)
	fmt.Println(result)
	// Output: JasonJason
}
