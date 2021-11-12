# god

Golang默认值设置工具，支持功能

- 支持近乎所有的内置类型
    - `int`
    - `int8`
    - `int32`
    - `int64`
    - `uint`
    - `uint8`
    - `uint32`
    - `uint64`
    - `bool`
    - `float32`
    - `float64`
- 支持复杂类型
    - `slice`
    - `map`
    - `struct`
- 支持配置
    - 配置所使用的`tag`名称
- 支持复杂结构初始化值
    - 支持`json`初始化值设置
    - 支持更易于书写的初始化值
- 支持类型别名
- 支持指针类型

## 使用方法

### 简单使用

使用非常简单，只需要调用`god.Set`就可以了

```go
package main

import `github.com/storezhang/god`

type test struct {
    Addr string `default:"127.0.0.1"`
    Port int    `default:"80"`
}

func main() {
    _test := new(test)
    if err := god.Set(_test); nil != err {
        panic(err)
    }
}
```

### 配置标签

可以很方便的使用其它`标签`，方法

```go
package main

import `github.com/storezhang/god`

type testByTag struct {
    Addr string `test:"127.0.0.1"`
    Port int    `test:"80"`
}

func main() {
    tag := new(testByTag)
    if err := god.Set(tag, god.Tag(`test`)); nil != err {
        panic(err)
    }
}
```

### 配置复杂类型

可以使用`json`来配置复杂的类型，比如

- `map`
- `slice`
- `struct`

`json`写法支持使用单引号`'`来替换转义字符`\"`，这样在书写默认值`json`更容易

```go
package main

import `github.com/storezhang/god`

type testByJson struct {
    Orders []string `default:"['mqtts', 'mqtt', 'wss', 'ws']"`
    // 同样支持这样写
    // Orders []string `default:"[\"mqtts\", \"mqtt\", \"wss\", \"ws\"]"`
}

func main() {
    json := new(testByJson)
    if err := god.Set(json); nil != err {
        panic(err)
    }
}
```
