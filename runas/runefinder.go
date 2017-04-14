package main

import (
    "strconv"
    "strings"
)

func AnalisarLinha(ucdLine string) (rune, string) {
    campos := strings.Split(ucdLine, ";")
    código, _ := strconv.ParseInt(campos[0], 16, 32)
    return rune(código), campos[1]
}
