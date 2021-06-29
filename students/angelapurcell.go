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
		var zero string
		s := strings.Split(file, " ")
		first := s[0] + s[1]
		num := strings.Replace(s[2],"(", "", 1)
		count := strings.Count(s[4], "0")
		for i:= 0; i < count; i++ {
			zero = zero + "0"

		}
		newPath := first + "_" + zero + num + ".txt"
		fmt.Println(newPath)
	}
}
