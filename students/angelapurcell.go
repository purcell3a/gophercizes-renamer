package main

//list of imports
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//list of constants

//list of variables
var files []string
var root = "/Users/angelapurcell/Workspace/personal/gophercizes/renamer/sample"



//main types for files


//list of functions
func getFileNames (files *[]string) filepath.WalkFunc{
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Fatal(err)
        }
		// IF FOLDER, IGNORE
		if info.IsDir(){
			return nil
		}
		if !strings.Contains(info.Name(), " "){
			return nil
		}
        *files = append(*files, info.Name())
        return nil
    }
}
// func Rename(oldpath, newpath string) error{

// }
//list of methods

//func main
func main() { 

	err := filepath.Walk(root, getFileNames(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files{
		fmt.Println(file)
	}
}
