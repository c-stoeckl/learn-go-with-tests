package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	Repeat("a", b.N)
}

func ExampleRepeat() {
	repeatedString := Repeat("a", 6)
	fmt.Println(repeatedString)
	// Output: aaaaaa
}
