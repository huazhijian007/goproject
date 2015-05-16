package main

/*
   #include <stdio.h>
   #include <stdlib.h>
   void hello() {
        printf("Hello, World!\n");
} */
import "C"
import "fmt"
import (
	"time"
)

type G interface {
	set() int
}
type Context struct {
	name     string
	id       uint32
	callName func(n string) int
}

func (this *Context) CallName(n string) bool {
	fmt.Printf("hello world\n")
	return true
}

func exec() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	for i := 0; i < 1; i++ {
		time.Sleep(1 * time.Second)
	}

	//fmt.Printf("hello\n")
	var a *int
	a = nil
	*a = 10
	//fmt.Printf("world\n")
}

func main() {
	//C.hello()
	//fmt.Printf("%d\n", c())
	go exec()

	time.Sleep(10 * time.Second)
}
