package main

//list of imports
import (
	"fmt"
	"os"
	"path/filepath"
)


//list of constants

//list of variables
var files []string
var root = "/sample"


//main types for files
//list of functions
// func Walk(root string, fn WalkFunc) error {
// 	fmt.Println("test")
// }

// IF FOLDER, IGNORE
// if info.IsDir() {
//     return nil
// } 

func getFileNames (files *[]string) filepath.walkFunc{
	return
}
func WalkDir(root string, fn fs.WalkDirFunc) error{
	fmt.Println("test")
}
//list of methods

//func main
func main() {
	fmt.Println("test")
}
