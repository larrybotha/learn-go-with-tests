package iteration

//import "github.com/davecgh/go-spew/spew"

func Repeat(value string, repeats int) string {
	var result string

	for n := 0; n < repeats; n++ {
		result = result + value
	}

	//spew.Dump(result)

	return result
}
