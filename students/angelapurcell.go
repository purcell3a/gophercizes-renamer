package main

//list of imports
import (
	"fmt"
	"os"
	"path/filepath"
	"log"
)


//list of constants

//list of variables
var files []string
var root = "/Users/angelapurcell/Workspace/personal/gophercizes/renamer/sample"



//main types for files
// IF FOLDER, IGNORE
// if info.IsDir() {
//     return nil
// } 

//list of functions
func getFileNames (files *[]string) filepath.WalkFunc{
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Fatal(err)
        }
        *files = append(*files, file)
        return nil
    }
}
//list of methods

//func main
func main() { 
	
	fmt.Println(root)
	err := filepath.Walk(root, getFileNames(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files{
		files = append(files, info.Name)
		fmt.Println(file)
	}
	// files = append(files, info.Name)
	// err := filepath.Walk(root, visit(&files))
}
