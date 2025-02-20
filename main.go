package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Liters float64
type Gallons float64

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	s = "changed"
	fmt.Printf("Hello and welcome, %s!\n", s)

	var x int
	x = 42
	fmt.Println(x)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}

	notes := [7]string{"a", "b", "c", "d", "e", "f", "g"}
	for idx, note := range notes {
		fmt.Println(idx, note)
	}

	var nums []int
	nums = append(nums, 0)
	nums = append(nums, 100)
	fmt.Println(nums)

	carFuel := Gallons(Liters(40.0) * 0.264)
	fmt.Printf("carFule: %0.1f\n", carFuel)
	fmt.Printf("carFule: %s\n", carFuel)

	item := Item{key: "theAnswer", val: 42}
	item.sayHi()

	var v2 NumInt
	v2 = IntType(5)
	fmt.Println("v2 =", v2.M1)

	foo()
	bar()

	resp := make(chan string)

	go opFile(resp)

	println("resp: ", <-resp)
	println("main exit")
}

func opFile(resp chan string) {
	// https://gobyexample.com/temporary-files-and-directories
	outFile, err := os.CreateTemp("", "foo")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(outFile.Name())

	println("filename", outFile.Name())

	contents := "Hello world!"

	os.WriteFile(outFile.Name(), []byte(contents), 0666)

	inFile, err := os.Open(outFile.Name())

	if err != nil {
		log.Fatal(err)
	}

	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	var outStr string

	for scanner.Scan() {
		println("line: ", scanner.Text())
		outStr += scanner.Text() + "\n"
	}

	resp <- outStr
}

func foo() {
	defer fmt.Println("foo exit")
	println("foo")
}

func bar() {
	println("bar")
	defer func() {
		if r := recover(); r != nil {

		}
	}()
	panic("PANIC")
}

type NumInt interface {
	M1()
	M2()
	M3()
}

type IntType int

func (t IntType) M1() {
	fmt.Println("M1")
}
func (t IntType) M2() {
	fmt.Println("M2")
}
func (t IntType) M3() {
	fmt.Println("M3")
}

type Item struct {
	key string
	val int
}

func (item Item) sayHi() {
	fmt.Printf("sayHi: %s\n", item)
}
