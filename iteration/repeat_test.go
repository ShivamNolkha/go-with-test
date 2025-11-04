package repeat

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){
	got := Repeat("*", 5)
	want := "*****"
	
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop(){
		Repeat("a", 10)
	}
}


func ExampleRepeat(){
	repeatedString := Repeat("a", 5)
	fmt.Println(repeatedString)
	// Output: aaaaa
}