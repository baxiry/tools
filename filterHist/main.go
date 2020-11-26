package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	list := []string{}

	user := os.Getenv("USER")
	file, err := os.Open("/home/" + user + "/.bash_history") //
	if err != nil {
		fmt.Printf("we have an error: %s", err)
	}
	defer file.Close()

	data := make([]byte, 1024*1024)
	lenData, err := file.Read(data)
	if err != nil {
		fmt.Println("error is ", err)
	}
	mydata := string(data[:lenData])
	cmds := strings.Split(mydata, "\n")

	for _, cmd := range cmds {
		// some cmd not importent/ like cd ls mv rm ..
		if len(cmd) < 4 {
			continue
		}
		if cmd[:3] == "git" || cmd[:2] == "rm" || cmd[:2] == "mv" || cmd[:2] == "cd" {
			continue
		}
		list = append(list, cmd)
	}

	res := filter(list)
	fdata := ""
	for _, cmd := range res {
		fdata += cmd
		fdata += "\n"
	}
	_, err = file.WriteAt([]byte(fdata), int64(lenData)) // Write at last file
	if err != nil {
		fmt.Printf("failed writing to file: %s", err)
	}
	//fmt.Println(fdata)
	err = ioutil.WriteFile("/home/"+user+"/.bash_history", []byte(fdata), 0644)
	if err != nil {
		fmt.Printf("could not write to cmds file %v", err)
	}

	//fmt.Println("don")

}

// filter commande line for uniquet it
func filter(slc []string) []string {
	res := []string{}
	for i := 0; i < len(slc); i++ {
		exit := true
		for j := i + 1; j < len(slc); j++ {
			if slc[i] == slc[j] {
				exit = false
			}
		}
		if exit {
			res = append(res, slc[i])

		}
	}
	return res
}
