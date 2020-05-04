package files

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// WriteToNewFile demonstrates use of io package to create a
// new file and write strings into it
func WriteToNewFile() {
	fmt.Println("Creating a new file and writing content into it...")

	f, err := os.Create("testfile.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	// Making sure the file will be closed
	defer f.Close()
	defer removeFile(*f)

	// Creating this reader as io.Copy needs a reader i.e.
	// type which implements Reader interface
	fmt.Println("Writing contents to file: ", f.Name())
	reader := strings.NewReader("Hello World")
	io.Copy(f, reader)
}

func removeFile(file os.File) {
	err := os.Remove(file.Name())
	if err != nil {
		fmt.Println("Could not delete file..")
		return
	}
	fmt.Printf("File %s has been deleted..\n", file.Name())
}
