package redsocial

import (
	"strings"
)

func ConvertirIngreso(input string) []string {
	vectorIngreso := strings.Split(input, " ")
	if len(vectorIngreso) >= 3 {
		parametro := strings.Join(vectorIngreso[1:], " ")
		comando := vectorIngreso[0]
		vectorIngreso = make([]string, 2)
		vectorIngreso[0], vectorIngreso[1] = comando, parametro
	}
	return vectorIngreso
}

func cmpstring(str1, str2 string) int {
	if str1 > str2 {
		return 1
	} else if str1 < str2 {
		return -1
	}
	return 0
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
