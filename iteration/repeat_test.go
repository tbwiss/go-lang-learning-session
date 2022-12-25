package iteration

import "testing"
import "fmt"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	repeat := Repeat("f", 5)
	fmt.Println(repeat)
	// Output: fffff
}

func TestRepeat(t *testing.T) {
	t.Run("repeating a string", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("pass amount", func(t *testing.T) {
		got := Repeat("a", 4)
		want := "aaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
