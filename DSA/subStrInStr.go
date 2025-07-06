package main

import "fmt"

func main() {
	str1 := "akhilesh"
	str2 := "sha"
	// str1 += str1 // for circular check
	// [:]
	for i := 0; i <= len(str1)-len(str2); i++ {
		if str2 == str1[i:len(str2)+i] { // sliding window technique
			fmt.Println("substr found")
			break
		}
	}

}


func main1() {
    str1 := "akhilesh"
    str2 := "les"

    // Create a map to store substrings of length len(str2) from str1
    substrMap := make(map[string]bool)

    // Populate the map with substrings from str1
    for i := 0; i <= len(str1)-len(str2); i++ {
        substrMap[str1[i:i+len(str2)]] = true
    }

    // Check if str2 exists in the map
    if substrMap[str2] {
        fmt.Println("Continuous substr found")
    } else {
        fmt.Println("Continuous substr not found")
    }
}
