package rubik

import (
	"os"
	// "io/ioutil"
	"fmt"
)

func tableGenerator() {
	if _, err := os.Stat("tables/G0.txt"); os.IsNotExist(err) {
		var table [4096]uint8
		table[1] = 1//
		
		f, err := os.Create("tables/G0.txt")
		if err != nil {
			fmt.Printf("error creating file: %v", err)
			return
		}
		defer f.Close()
		for i := 0; i < len(table); i++ {
			_, err = f.WriteString(fmt.Sprintf("%d", table[i]))
			if err != nil {
				fmt.Printf("error writing to file: %v", err)
			}
		}
	}
}