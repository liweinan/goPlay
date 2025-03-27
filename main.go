package main

import (
	"bufio"
	"fmt"
	"goPlay/server"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"reflect"
)

type Liters float64
type Gallons float64

//type Config struct {
//	Name    string `yaml:"name"`
//	Version string `yaml:"version"`
//	Port    int    `yaml:"port"`
//	Debug   bool   `yaml:"debug"`
//	Servers []struct {
//		Host string `yaml:"host"`
//		Port int    `yaml:"port"`
//	} `yaml:"servers"`
//}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Name    string   `yaml:"name"`
	Version string   `yaml:"version"`
	Port    int      `yaml:"port"`
	Debug   bool     `yaml:"debug"`
	Servers []Server `yaml:"servers"`
}

func main() {
	yamlData := `
name: MyApp
version: 1.0.0
port: 8080
debug: true
servers:
  - host: server1.example.com
    port: 8000
  - host: server2.example.com
    port: 8001
`

	var config Config
	err := yaml.Unmarshal([]byte(yamlData), &config)
	if err != nil {
		log.Fatalf("error parsing YAML: %v", err)
	}

	fmt.Printf("Parsed config: %+v\n", config)

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

	resp1 := make(chan string)

	go opFile(resp1)

	println("resp: ", <-resp1)

	playSlices()

	playEmptyInterfaceAsString("strstr")
	playEmptyInterfaceAsString("42")

	println("main exit")

	server.Serve()
}

func playEmptyInterfaceAsString(incomeVal interface{}) {
	if reflect.TypeOf(incomeVal).Kind() == reflect.String {
		println("incomeVal: ", incomeVal.(string))
	} else {
		println("Not supported type!")
	}
}

// https://medium.com/@tucnak/why-go-is-a-poorly-designed-language-1cc04e5daf2
func playSlices() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println("nums: ", nums)
	log.Println("nums[2:] ", nums[2:])

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
