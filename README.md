# mengpo（孟婆）

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
- 支持`环境变量`

## 为什么要叫孟婆

按照本人的一惯作风，所有项目都在`中国古代神话人物或者历史名人`寻找和项目`意义相近`的`神话人物或者历史人`来做作为项目的名称，原因

- 去TMD`崇洋媚外`
- 致敬`中华民族的先贤`

`孟婆`作为阴司掌管生死轮回的大神，有将一切人物还原到最被的状态，这和`设置默认值`不谋而和，故而使用`孟婆`来命名项目是合适的

## 使用方法

### 安装

安装非常简单，推荐使用`go.mod`来使用`孟婆`

```go
package main

import `github.com/storezhang/mengpo`

func main() {
    // xxx
}
```

或者

```shell
go get github.com/storezhang/mengpo
```

### 简单使用

使用非常简单，只需要调用`mengpo.Set`就可以了

```go
package main

import `github.com/storezhang/mengpo`

type testByNormal struct {
    Addr string `default:"127.0.0.1"`
    Port int    `default:"80"`
}

func main() {
    normal := new(testByNormal)
    if err := mengpo.Set(normal); nil != err {
        panic(err)
    }
}
```

### 配置标签

可以很方便的使用其它`标签`，方法

```go
package main

import `github.com/storezhang/mengpo`

type testByTag struct {
    Addr string `test:"127.0.0.1"`
    Port int    `test:"80"`
}

func main() {
    tag := new(testByTag)
    if err := mengpo.Set(tag, mengpo.Tag(`test`)); nil != err {
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

import `github.com/storezhang/mengpo`

type testByJson struct {
    Orders []string `default:"['mqtts', 'mqtt', 'wss', 'ws']"`
    // 同样支持这样写
    // Orders []string `default:"[\"mqtts\", \"mqtt\", \"wss\", \"ws\"]"`
}

func main() {
    json := new(testByJson)
    if err := mengpo.Set(json); nil != err {
        panic(err)
    }
}
```

### 使用环境变量

`孟婆`支持使用环境变量来配置默认值

```go
package main

import `github.com/storezhang/mengpo`

type testByEnv struct {
    Order string `default:"${ORDER}"`
    // 同样支持这种写法
    // Order string `default:"$ORDER"`
}

func main() {
    env := new(testByEnv)
    if err := mengpo.Set(env); nil != err {
        panic(err)
    }
}
```
