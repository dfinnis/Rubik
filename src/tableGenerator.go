package rubik

import (
	"os"
	"io/ioutil"
	"fmt"
)

func tableGenerator() {
	if _, err := os.Stat("tables/G0.txt"); os.IsNotExist(err) {
		tableG0 := []byte{115, 111, 109, 101, 10}

		err := ioutil.WriteFile("tables/G0.txt", tableG0, 0644)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}	
	}
}