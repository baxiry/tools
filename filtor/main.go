package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	extension = []string{}
	path      string
)

func checkArgs() {
	if len(os.Args[:]) < 2 {
		fmt.Println("Please Type name of directory  you wanna filter!")
		os.Exit(0)
	}
	path = os.Args[1]
}

func readCurrentDir(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileList, _ := file.Readdirnames(0)

	for _, name := range fileList {
		//fmt.Println(i, "  ", name)
		isdir, err := os.Stat(path + "/" + name)
		if err != nil {
			fmt.Println(err)
		}
		if !isdir.IsDir() {
			extension = strings.Split(name, ".")
			if len(extension) == 1 {
				continue
			}
			l := len(extension) - 1

			newDir := extension[l]
			fmt.Println(newDir)

			err = os.MkdirAll(path+"/"+extension[l], 0777)
			if err != nil {
				panic(err)
			}
			err = os.Rename(path+"/"+name, path+"/"+newDir+"/"+name)
			if err != nil {
				//fmt.Println("permession ????")
				panic(err)
			}
		}
	}
}

func main() {
	checkArgs()
	readCurrentDir(path)
}
