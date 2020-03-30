package inject

/**
	go test ./test -v -run Countdown
 */
import (
	"bytes"
	"go-example/learn"
	"testing"
)


func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	main.Greet(&buffer,"Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

//func Greet(writer *bytes.Buffer, name string) {
//	fmt.Fprintf(writer, "Hello, %s", name)
//}
