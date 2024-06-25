package ch1 


import (
	"log"
	"os"
)

func testMkdir() {
	err := os.Mkdir("testdir", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	err = os.WriteFile("testdir/testFile.txt", []byte("hello, maderator!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}