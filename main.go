package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"reflect"
	"strings"
	"unsafe"

	"gopkg.in/yaml.v3"
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

	playInterface()
	fmt.Println("----------------------------------------")
	checkInterface(42)
	checkInterface("hello")
	checkInterface([]int{1, 2, 3})

	// 演示类型断言的必要性
	fmt.Println("----------------------------------------")
	demonstrateTypeAssertion()

	// 演示各种获取类型元信息的方法
	fmt.Println("----------------------------------------")
	demonstrateTypeMetadata()

	//parseJson()

	//walkDir()
	//
	//callYaml()
	//
	//useMap()
	//
	//playUnsafe()
	//
	//playStructTags()

	//f := Form{Name: ""}
	//err := ValidateStruct(f) // 返回错误："Name is required"
	//println(err.Error())
	//all()

	//server.Serve()
}

func checkInterface(v interface{}) {
	switch value := v.(type) {
	case int:
		fmt.Printf("整数: %d\n", value)
	case string:
		fmt.Printf("字符串: %s\n", value)
	default:
		fmt.Println("未知类型")
	}
}

// demonstrateTypeAssertion 演示类型断言的必要性
func demonstrateTypeAssertion() {
	fmt.Println("=== 类型断言必要性演示 ===")

	// 创建一个包含不同类型值的切片
	var values []interface{}
	values = append(values, 42, "hello", 3.14, true, []int{1, 2, 3})

	fmt.Println("原始值:", values)

	// 1. 不使用类型断言 - 只能进行有限的操作
	fmt.Println("\n1. 不使用类型断言:")
	for i, v := range values {
		fmt.Printf("  索引 %d: 类型=%T, 值=%v\n", i, v, v)
		// 只能调用空接口支持的方法，几乎没有任何有用的操作
	}

	// 2. 使用类型断言 - 可以访问具体类型的值和方法
	fmt.Println("\n2. 使用类型断言:")
	for i, v := range values {
		switch value := v.(type) {
		case int:
			fmt.Printf("  索引 %d: 整数 %d, 可以进行数学运算: %d + 10 = %d\n",
				i, value, value, value+10)
		case string:
			fmt.Printf("  索引 %d: 字符串 \"%s\", 长度: %d, 大写: \"%s\"\n",
				i, value, len(value), strings.ToUpper(value))
		case float64:
			fmt.Printf("  索引 %d: 浮点数 %f, 平方: %f\n",
				i, value, value*value)
		case bool:
			fmt.Printf("  索引 %d: 布尔值 %t, 取反: %t\n",
				i, value, !value)
		case []int:
			fmt.Printf("  索引 %d: 整数切片 %v, 长度: %d, 求和: %d\n",
				i, value, len(value), sumSlice(value))
		default:
			fmt.Printf("  索引 %d: 未知类型 %T, 值: %v\n", i, v, v)
		}
	}

	// 3. 演示类型断言失败的情况
	fmt.Println("\n3. 类型断言失败处理:")
	var mixed interface{} = "hello"

	// 安全的类型断言
	if intValue, ok := mixed.(int); ok {
		fmt.Printf("   安全断言成功: %d\n", intValue)
	} else {
		fmt.Printf("   安全断言失败: 值不是 int 类型\n")
	}

	// 不安全的类型断言（会 panic）
	fmt.Println("   尝试不安全的类型断言...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("   捕获到 panic: %v\n", r)
		}
	}()

	// 这行会 panic，但被 defer 捕获
	// intValue := mixed.(int) // 这行被注释掉，避免实际 panic
	fmt.Println("   注意: 不安全的类型断言会导致 panic")
}

// sumSlice 计算整数切片的总和
func sumSlice(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// demonstrateTypeMetadata 演示各种获取类型元信息的方法
func demonstrateTypeMetadata() {
	fmt.Println("=== 类型元信息获取方法演示 ===")

	// 测试不同类型的值
	testValues := []interface{}{
		42,                     // int
		"hello",                // string
		3.14,                   // float64
		true,                   // bool
		[]int{1, 2, 3},         // []int
		map[string]int{"a": 1}, // map[string]int
		&struct{}{},            // *struct{}
	}

	for i, v := range testValues {
		fmt.Printf("\n--- 测试值 %d: %v ---\n", i+1, v)

		// 1. reflect.TypeOf(v) - 获取类型信息
		t := reflect.TypeOf(v)
		fmt.Printf("1. reflect.TypeOf:\n")
		fmt.Printf("   类型: %T\n", t)
		fmt.Printf("   类型名: %s\n", t.Name())
		fmt.Printf("   种类: %s\n", t.Kind())
		fmt.Printf("   字符串表示: %s\n", t.String())
		fmt.Printf("   大小: %d bytes\n", t.Size())
		fmt.Printf("   对齐: %d\n", t.Align())

		// 2. reflect.ValueOf(v) - 获取值信息
		val := reflect.ValueOf(v)
		fmt.Printf("2. reflect.ValueOf:\n")
		fmt.Printf("   值: %v\n", val.Interface())
		fmt.Printf("   类型: %s\n", val.Type())
		fmt.Printf("   种类: %s\n", val.Kind())
		fmt.Printf("   是否为零值: %t\n", val.IsZero())

		// 3. fmt.Printf("%T", v) - 格式化输出类型
		fmt.Printf("3. fmt.Printf(\"%%T\"):\n")
		fmt.Printf("   类型: %T\n", v)

		// 4. reflect.TypeOf(v).Kind() - 获取基本类型种类
		kind := reflect.TypeOf(v).Kind()
		fmt.Printf("4. reflect.TypeOf(v).Kind():\n")
		fmt.Printf("   基本类型: %s\n", kind)

		// 5. 类型断言的不同形式
		fmt.Printf("5. 类型断言:\n")

		// 安全断言
		if intVal, ok := v.(int); ok {
			fmt.Printf("   安全断言为int: %d\n", intVal)
		}
		if strVal, ok := v.(string); ok {
			fmt.Printf("   安全断言为string: %s\n", strVal)
		}
		if floatVal, ok := v.(float64); ok {
			fmt.Printf("   安全断言为float64: %f\n", floatVal)
		}
		if boolVal, ok := v.(bool); ok {
			fmt.Printf("   安全断言为bool: %t\n", boolVal)
		}
		if sliceVal, ok := v.([]int); ok {
			fmt.Printf("   安全断言为[]int: %v\n", sliceVal)
		}
		if mapVal, ok := v.(map[string]int); ok {
			fmt.Printf("   安全断言为map[string]int: %v\n", mapVal)
		}

		// 6. 特殊类型信息
		fmt.Printf("6. 特殊类型信息:\n")
		switch t.Kind() {
		case reflect.Slice:
			fmt.Printf("   是切片类型，元素类型: %s\n", t.Elem())
		case reflect.Map:
			fmt.Printf("   是映射类型，键类型: %s, 值类型: %s\n", t.Key(), t.Elem())
		case reflect.Ptr:
			fmt.Printf("   是指针类型，指向: %s\n", t.Elem())
		case reflect.Struct:
			fmt.Printf("   是结构体类型，字段数: %d\n", t.NumField())
		}
	}

	// 演示结构体类型的详细信息
	fmt.Printf("\n--- 结构体类型详细信息 ---\n")
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "Alice", Age: 30}
	t := reflect.TypeOf(p)

	fmt.Printf("结构体类型: %s\n", t.Name())
	fmt.Printf("字段数: %d\n", t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("字段 %d: %s (类型: %s, 标签: %s)\n",
			i, field.Name, field.Type, field.Tag)
	}
}

func playInterface() {
	var list []interface{}
	list = append(list, 42, "hello", 3.14)
	fmt.Println(list) // 输出: [42 hello 3.14]
}

type Pod struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
}

type Metadata struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
}

func parseJson() {
	obj := map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "Pod",
		"metadata": map[string]interface{}{
			"name":      "example-pod",
			"namespace": "default",
			"labels": map[string]interface{}{
				"app": "example",
			},
		},
	}

	name := obj["metadata"].(map[string]interface{})["name"].(string)

	fmt.Println(name)
}

func walkDir() {
	root := "/tmp"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
}

type Form struct {
	Name string `validate:"required"`
}

func ValidateStruct(s interface{}) error {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	fmt.Println("type:", t.Name())
	fmt.Println("value:", v)

	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("field:%v\n", t.Field(i))
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
	v2.M1() // Call M1() separately since it doesn't return a value
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
