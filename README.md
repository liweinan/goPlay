# goPlay

Go语言学习和实验项目，包含各种Go语言特性的演示和实验。

## 功能特性

### 空接口 (`interface{}`) 演示 (`playInterface`)
- 展示空接口的基本概念和用法
- 演示如何将不同类型的值存储到 `[]interface{}` 切片中
- 说明空接口是Go语言中实现多态的基础
- 展示空接口的局限性（需要类型断言才能访问具体类型的方法）

### 类型断言演示 (`demonstrateTypeAssertion`)
- 演示类型断言的必要性
- 展示不使用类型断言时的局限性
- 展示使用类型断言后的强大功能
- 演示安全和不安全的类型断言

### 类型元信息获取演示 (`demonstrateTypeMetadata`)
- `reflect.TypeOf(v)` - 获取完整类型信息
- `reflect.ValueOf(v)` - 获取值信息
- `fmt.Printf("%T", v)` - 格式化输出类型
- `reflect.TypeOf(v).Kind()` - 获取基本类型种类
- 类型断言的不同形式
- 特殊类型信息（切片、映射、指针、结构体）

## 运行示例

```bash
# 编译并运行主程序
go run main.go

# 编译特定文件
go build standalone/rwLockExp.go
```

## Go语言中的 `interface{}`

### 什么是空接口？
`interface{}` 是Go语言中的空接口，它是所有类型的超类型。任何值都可以赋值给 `interface{}` 类型的变量。

### 基本用法
```go
var list []interface{}
list = append(list, 42, "hello", 3.14)
fmt.Println(list) // 输出: [42 hello 3.14]
```

### 空接口的特点
- **类型安全**: 可以存储任何类型的值
- **多态性**: 实现Go语言的多态特性
- **局限性**: 只能调用空接口本身的方法（几乎没有）
- **类型断言**: 需要通过类型断言访问具体类型的值和方法

### 使用场景
- **通用容器**: 存储不同类型的值
- **函数参数**: 接受任意类型的参数
- **JSON处理**: 处理动态JSON数据
- **反射**: 与反射包配合使用

## 输出示例

程序会输出以下内容：
1. 空接口演示 (`playInterface`)
2. 类型断言必要性演示 (`demonstrateTypeAssertion`)
3. 类型元信息获取方法演示 (`demonstrateTypeMetadata`)

每个演示都包含详细的说明和实际运行结果，帮助理解Go语言中类型系统的各种特性。