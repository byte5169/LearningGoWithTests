package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	var result string
	result = Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}