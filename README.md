[![CircleCI](https://circleci.com/gh/uccu/go-c0nvert/tree/master.svg?style=svg)](https://circleci.com/gh/uccu/go-c0nvert/tree/master)
[![Maintainability](https://api.codeclimate.com/v1/badges/90b7fb370c2f9926877e/maintainability)](https://codeclimate.com/github/uccu/go-c0nvert/maintainability)
[![codecov](https://codecov.io/gh/uccu/go-c0nvert/branch/master/graph/badge.svg?token=MU57MLHI6Z)](https://codecov.io/gh/uccu/go-c0nvert)
[![GitHub issues](https://img.shields.io/github/issues/uccu/go-c0nvert)](https://github.com/uccu/go-c0nvert/issues)
![GitHub](https://img.shields.io/github/license/uccu/go-c0nvert)

# go-c0nvert
 Arbitrary Conversion

### 安装

```go
go get -u github.com/uccu/go-c0nvert
```

### 使用


```go
// any => string
c0nvert.To[string](value any) string

// bool => string
// @param isUpper 是否是大写
c0nvert.To[string](value bool, isUpper bool) string

// float => string
// @param decimals 小数的位数
c0nvert.To[string](value float64|float32, decimals int) string

// slice => string
// @param sep 分割符号
c0nvert.To[string](value []any, sep string) string
c0nvert.To[string](value [][]any, sep1 string, sep2 string) string
...
// tips. []byte、[]rune如果没有设置sep则默认按string处理

// any => int
c0nvert.To[int](value any) int

// any => float
// @param decimals 小数的位数
c0nvert.To[float32](value any, decimals int) float32

// any => bool
c0nvert.To[bool](value any) bool

// string => bool
// @param fuzzy true开启模糊模式,"0"和"false"（无视大小写）也会判定为false
c0nvert.To[bool](value string, fuzzy bool) bool

// int32 => bool 0为false其他为true
c0nvert.To[bool](value int32) bool

// 分割字符串到切片中
// string => []int
// @param sep 分割的字符串
c0nvert.Split[[]int](value[, sep] string) []int

```

### 说明

#####  1. `To[bool](src any[, fuzzy bool]) bool`

任意类型转换为`bool`类型，当传入值类型为零值时返回`false`，非零值时返回`true`

- `src` **`any`** 任意类型值
- `fuzzy` **`bool`** 模糊模式，当类型为`string`、`[]byte`、`[]rune`，值为空字符串、`false`（无视大小写）、`0`也判断为`false`

```go
c0nvert.To[bool](1) // true
c0nvert.To[bool](0) // false
c0nvert.To[bool]("") // false
c0nvert.To[bool]("false") // true
c0nvert.To[bool]("false", true) // false
c0nvert.To[bool]([]byte("")) // true
c0nvert.To[bool]([]byte(""), true) // false
c0nvert.To[bool]([]byte(nil)) // false
c0nvert.To[bool]([]byte(nil), true) // false

```



### 以后的规划

- `Split` 支持分割多级切片
- `To[string]([]T[, ...string]) string` 合并字符串功能单独分出来
