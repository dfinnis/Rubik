package rubik

import (
	"os"
	"io/ioutil"
	"fmt"
)

func tableGenerator() {
	if _, err := os.Stat("tables/G0.txt"); os.IsNotExist(err) {
		err := ioutil.WriteFile("tables/G0.txt", []byte("Hello"), 0644)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}	
	}
}