package iterate

import (
	"fmt"
	"testing"
)

func TestIierate(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("got: %q and want: %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 10)
	fmt.Println(result)
	// Output: aaaaaaaaaa
}
