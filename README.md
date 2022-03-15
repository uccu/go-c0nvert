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
c0nvert.ToString(value interface{}) string

// @param isUpper 是否是大写
c0nvert.ToString(value bool, isUpper bool) string

// @param decimals 小数的位数
c0nvert.ToString(value float64|float32, decimals int) string

// @param sep 分割符号
c0nvert.ToString(value []interface{}, sep string) string
c0nvert.ToString(value [][]interface{}, sep1 string, sep2 string) string
...
// tips. []byte、[]rune如果没有设置sep则默认按string处理
```

