package main

import (
	"fmt"
	"os"
	"testing"
)

func TestActionCopy(t *testing.T) {
	var err error
	var newFile *os.File

	// Create test directory and files
	if err = os.Mkdir("tmp_test", 0700); err != nil {
		t.Fatal(err)
	}

	if newFile, err = os.Create("tmp_test/1.txt"); err != nil {
		t.Fatal(err)
	}

	newFile.Close()

	aItemCopy := actionItem{
		Action:        "copy",
		Src:           "tmp_test",
		Dst:           "bin_test",
		DstClearFirst: true,
	}
	actionCopy(aItemCopy)

	_, existsErr := os.Stat("bin_test/1.txt")

	if err = os.RemoveAll("tmp_test"); err != nil {
		fmt.Println("error removing test files. Please delete manually")
	}

	if err = os.RemoveAll("bin_test"); err != nil {
		fmt.Println("error removing test files. Please delete manually")
	}

	if existsErr != nil {
		t.Fatal(existsErr, "ie: the file wasn't copied to the destination.")
	}

}
