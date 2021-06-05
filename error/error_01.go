/*
# @Time : 2021/6/5 11:14
# @Author : team_go
# @File : error_01.go
# @Software: GoLand
*/

package main

import (
    "errors"
    "fmt"
)

type errorSting string

func (e errorSting) Error() string {
    return string(e)
}

func New(txt string) error {
    return errorSting(txt)
}

var ErrorNameType = New("EOF")
var ErrorStructType = errors.New("EOF")

func main() {
    if ErrorNameType == New("EOF") {
        fmt.Println("error Name error")
    }
    if ErrorStructType.Error() ==errors.New("EOF").Error() {
        fmt.Println("error Struct error")
    }
    if ErrorNameType == ErrorStructType {
        fmt.Println("==")
    }
}
