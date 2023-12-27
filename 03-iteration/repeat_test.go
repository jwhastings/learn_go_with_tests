package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", repeated, expected)
	}
}

func BenchmarkRepeatCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("abc", 5)
	}
}

func BenchmarkRepeatStdLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("abc", 5)
	}
}

func ExampleRepeat() {
	repeat := Repeat("abc", 3)
	fmt.Println(repeat)
	// Output: abcabcabc
}
