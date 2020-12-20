package main

import (
	"fmt"
	"github.com/dybrkr/wiki_action_test/calc"
)

var (
	temp = calc.Calc{}
)

func main(){
	fmt.Println(temp.Add(1,2))
}