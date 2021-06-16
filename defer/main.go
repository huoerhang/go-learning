package main

//import "fmt"
import (
	"io"
	"os"
)

func main() {
	f, err := os.Create("notes.txt")

	if err != nil {
		return
	}

	defer f.Close()

	if _, err = io.WriteString(f, "Learning GO!"); err != nil {
		return
	}

	f.Sync()
	// for i := 1; i <= 4; i++ {
	// 	defer fmt.Println("deferred", -i)
	// 	fmt.Println("regular", i)
	// }
}
