package test

import (
	"fmt"
	"strings"
)

func main() {
	slice_url := strings.Split("/parts/", "/")

	fmt.Println(len(slice_url))
	fmt.Println(slice_url[0])
	fmt.Println(slice_url[1])
	fmt.Println(slice_url[2]) /*

		//  /parts ou /parts/
		if len(slice_url) == 2 ||
			(len(slice_url) == 3 && slice_url[2] != "") {
	}*/
}
