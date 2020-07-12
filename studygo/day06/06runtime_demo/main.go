package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller()

func f1()  {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funName := runtime.FuncForPC(pc).Name()
	fmt.Println(funName)
	fmt.Println(file) //06runtime_demo/main.go
	fmt.Println(path.Base(file))
	fmt.Println(line) //11
}

func main()  {
	f1()
}
