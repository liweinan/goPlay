# goPlay

Go语言学习和实验项目，包含各种Go语言特性的演示和实验。

## 功能特性

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

## 输出示例

程序会输出以下内容：
1. 接口类型演示
2. 类型断言必要性演示
3. 类型元信息获取方法演示

每个演示都包含详细的说明和实际运行结果，帮助理解Go语言中类型系统的各种特性。