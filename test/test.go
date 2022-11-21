package main

import (
	"fmt"
	"os"
	"parp"
)

func main() {
	params := parp.Process(os.Args[1:len(os.Args)], []string{"-n", "-p"}, []string{"--trace"})
	fmt.Println(params)
	n := params["-n"].Int()
	p := params["-p"].String()
	trace := params["--trace"].Bool()
	fmt.Printf("n: %v\np: %v\ntrace: %v\n", n, p, trace)
}
