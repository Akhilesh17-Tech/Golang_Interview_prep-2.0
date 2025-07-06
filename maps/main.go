package main

import "fmt"

func main() {

	//map in go :  mao is built in data type to store keyvalue pair
	var maps = make(map[string]string)

	maps["name"] = "akhilesh"
	maps["age"] = "20"
	maps["city"] = "indore"

	fmt.Println(maps["city"])

	// check if exist
	value, exist := maps["citys"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("not exist")
	}

	// delete key from map
	delete(maps, "city")

	fmt.Println(maps)

	// create and initialze map
	newMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	// maps are undordered
	fmt.Println(newMap)

}
