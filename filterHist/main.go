package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	list := []string{}

	file, err := os.Open("/home/fedora/.bash_history") //
	if err != nil {
		fmt.Errorf("we have an error: %s", err)
	}
	defer file.Close()

	data := make([]byte, 1024*1024)
	i, err := file.Read(data)
	if err != nil {
		fmt.Println("error is ", err)
	}
	mydata := string(data[:i])
	cmds := strings.Split(mydata, "\n")

	for _, cmd := range cmds {
		if len(cmd) < 2 {
			continue
		}
		if cmd[:2] != "cd" && cmd[:2] != "ls" {
			list = append(list, cmd)
		}
	}

	res := filter(list)
	for i, cmd := range res {
		fmt.Println(i, cmd)
	}
	test := []string{"a", "b", "2", "3", "3", "4", "1", "a", "c"}
	fmt.Println(filter(test))
}

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
