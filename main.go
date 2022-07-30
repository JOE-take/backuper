package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	BACKUPPATH string = "./"
)
func DirSize(path string) (int64, error){
	var size int64
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil{
			return err
		}
		if !info.IsDir(){
			size += info.Size()
		}
		return nil
	})
	if err != nil{
		return -1, err
	}
	return size, nil
}

func main(){
	filenames, err := os.ReadDir(BACKUPPATH);
	if err != nil {
		log.Fatal(err)
	}

	for _, itr := range filenames {
		fileinfo , err:= itr.Info()
		if err != nil{
			log.Fatal(err)
		}
		
		if itr.IsDir(){
			size, err :=  DirSize(itr.Name())
			if err != nil{
				log.Fatal()
			}
			fmt.Println(itr.Name(), size)
			
		} else {
			fmt.Println(itr.Name(), ", " ,fileinfo.Size())	
		}

	}
}
