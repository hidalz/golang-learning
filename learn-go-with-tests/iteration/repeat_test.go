package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("want %q but got %q", expected, repeated)
	}
}

func TestConcatenateFoo(t *testing.T) {
	concatenated := ConcatenateFoo("b")
	expected := "bfoo"

	if concatenated != expected {
		t.Errorf("want %q but got %q", expected, concatenated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}
