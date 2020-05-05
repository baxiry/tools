package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	extension = []string{}
)

func readCurrentDir(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileList, _ := file.Readdirnames(0)

	for i, name := range fileList {
		fmt.Println(i, "  ", name)
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
				fmt.Println("OKKKKKKKKKKKKK")
				panic(err)
			}
			err = os.Rename(path+"/"+name, path+"/"+newDir+"/"+name)
			if err != nil {
				fmt.Println("permession ????")
				panic(err)
			}
		}
	}
}

func main() {
	//slc := []string{"a", "a", "b", "c", "d", "b", "a", "b", "c", "c", "c", "c", "b"}
	//u := uniqitem(slc)
	//fmt.Println(u)
	path := os.Args[1]
	readCurrentDir(path)
}

// uniqitem filters multy items to uniq item
func uniqitem(slc []string) []string {
	var isUniq bool
	var uslc []string
	for k, v := range slc {
		isUniq = true
		for i := k + 1; i < len(slc); i++ {

			if v == slc[i] {
				isUniq = false
				continue
			}
		}
		if isUniq == true {
			uslc = append(uslc, v)
		}

	}
	return uslc
}