package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Basic Addition", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectResult(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectResult(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
