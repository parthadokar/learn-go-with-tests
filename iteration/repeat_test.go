package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T)  {
	repeated := Repeat("a",5)
	output := "aaaaa"

	if repeated != output {
		t.Errorf("expected %q but got %q",repeated,output)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a",5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a",5)
	fmt.Println(repeated)
	// Output: aaaaa
}