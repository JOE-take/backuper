package main

import (
	"fmt"
	"io/ioutils"
)

const (
	BACKUPPATH string = "./"
)

func main(){
	filenames := ioutils.ReadDir(BACKUPPATH);
	for itr, _ := range filenames {
		fmt.Println(itr)	
	}
}
