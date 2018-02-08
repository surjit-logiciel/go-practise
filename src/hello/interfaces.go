package main

import (
	"fmt"
)

type VowelFinder interface {
	FindVowel() []rune
}