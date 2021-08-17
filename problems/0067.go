package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	var c int = 0

	res := ""

	for i >= 0 || j >= 0 || c == 1 {
		if i >= 0 {
			c += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			c += int(b[j] - '0')
			j--
		}
		res = string('0'+c%2) + res
		c /= 2
	}

	return res
}
