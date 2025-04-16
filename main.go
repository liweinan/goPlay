package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"log"
	"os"
	"reflect"
	"unsafe"
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
	root := "/tmp"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})

	//callYaml()

	//useMap()

	//playUnsafe()

	//playStructTags()

	//f := Form{Name: ""}
	//err := ValidateStruct(f) // 返回错误："Name is required"
	//println(err.Error())
	//all()

	//server.Serve()
}

type Form struct {
	Name string `validate:"required"`
}

func ValidateStruct(s interface{}) error {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("validate")
		value := v.Field(i).String()

		if tag == "required" && value == "" {
			return fmt.Errorf("%s is required", field.Name)
		}
	}
	return nil
}

type User struct {
	ID    int    `json:"id" db:"user_id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email,omitempty"` // omitempty 表示零值时忽略
}

func playStructTags() {
	u := User{}
	t := reflect.TypeOf(u) // 获取结构体类型信息

	fmt.Println("" +
		"Struct Tags:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // 获取字段信息
		tag := field.Tag    // 获取标签
		fmt.Printf("Field: %s, Tag: %s\n", field.Name, tag)
	}
}

// https://medium.com/lyonas/go-type-casting-starter-guide-a9c1811670c5
func playUnsafe() {
	var i uint64 = 10
	var ptr *uint64
	ptr = (*uint64)(unsafe.Pointer(&i))
	*ptr = 20
	fmt.Printf("sizeof i: %d\n", unsafe.Sizeof(i))
}

func all() {
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
	fmt.Printf("carFule: %v\n", carFuel)

	item := Item{key: "theAnswer", val: 42}
	item.sayHi()

	var v2 NumInt
	v2 = IntType(5)
	v2.M1()  // Call M1() separately since it doesn't return a value
	fmt.Println("v2 =", v2)

	foo()
	bar()

	resp1 := make(chan string)

	go opFile(resp1)

	println("resp: ", <-resp1)

	playSlices()

	playEmptyInterfaceAsString("strstr")
	playEmptyInterfaceAsString("42")

	println("main exit")
}

func useMap() {
	ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	fmt.Println(ranks["gold"])
}

func callYaml() {
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
	// https://gosamples.dev/print-type/
	fmt.Println(yamlData)
	fmt.Printf("type of yamlData: %s\n", reflect.TypeOf(yamlData).Kind().String())
	fmt.Printf("type of yamlData: %T\n", yamlData)

	var config Config
	err := yaml.Unmarshal([]byte(yamlData), &config)
	if err != nil {
		log.Fatalf("error parsing YAML: %v", err)
	}

	fmt.Printf("Parsed config: %+v\n", config)
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
	fmt.Printf("Hi from item: %v\n", item)
}
