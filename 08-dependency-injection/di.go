package di

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, x string) {
	// convoluted approach:
	// 	- build string
	// 	- write to the buffer after converting to a buffer string
	// 		and then getting the bytes
	// greeting := fmt.Sprintf("Hello, %s", x)
	// w.Write(bytes.NewBufferString(greeting).Bytes())

	// instead, we can use Fprintf which:
	// 	- accepts anything that implements io.Writer
	// 	- allows for formatting text
	// 	- uses io.Writer.Write below the hood to write to the writer
	// 		after formatting
	fmt.Fprintf(w, "Hello, %s", x)
}
