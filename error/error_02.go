/*
# @Time : 2021/6/5 11:24
# @Author : team_go
# @File : error_02.go
# @Software: GoLand
*/
package main

import (
    "fmt"
)

type errorSting02 struct {
    s string
}

func (e errorSting02) Error() string {
    return e.s
}

func NewError(txt string) error {
    return errorSting02{txt}
}

var ErrorType = NewError("EOF")

func Positive(n int) bool {
    if n == 0 {
        panic("error")
    }
    return n > -1
}

func main() {
    defer func() {
        if recover() != nil {
            fmt.Println("00")
        }
    }()
    // 返回的是结构体~  类型和值全部相等  就满足相等
    if ErrorType == NewError("EOF") {
        fmt.Println("error Name error :", ErrorType)
    }
    if Positive(0){
        fmt.Println("0")
    }else {
        fmt.Println("6")
    }
}
