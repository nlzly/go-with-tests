package iteration

import (
	"fmt"
	"testing"
)

func assertCorrectMessage(t testing.TB, expected, repeated string) {
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
func TestRepeat(t *testing.T) {
	t.Run("print the letter a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("print letter n number of times", func(t *testing.T) {
		repeated := Repeat("c", 6)
		expected := "cccccc"
		assertCorrectMessage(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
