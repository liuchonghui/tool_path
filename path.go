package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main()  {

	Log("start path")
	args := os.Args
	if args == nil || len(args) < 2 || len(args) > 3 {
		Usage()
		os.Exit(0)
	}
	input := args[1]
	abstractPath, err := filepath.Abs(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(abstractPath)
	}
}
var Log = func(msg... string) {
	//fmt.Println(msg)
}

var Usage = func() {
	fmt.Println("Usage of path:");
	fmt.Println("	path [option] source");
}