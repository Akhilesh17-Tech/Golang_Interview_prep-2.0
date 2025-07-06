package main

import "fmt"

func main() {

	fmt.Println(reverseString("akhilesh"))
}

func reverseString(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result

}
