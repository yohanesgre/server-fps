package main

import (
	"fmt"
	"github.com/yohanesgre/server-fps/core"
)

func main(){
	builder:=InitBuilder(10)
	fmt.Println(builder.position)
}