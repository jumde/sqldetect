package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("SQLi Detector is running...")
	r := bufio.NewReader(os.Stdin)
	for {
		line, _, err := r.ReadLine()
		// process buf
		if err != nil && err != io.EOF {
			//			log.Println(s)
			log.Fatal(err)
		}
		// s is the sql statement passed in from tshark
		fp := fingerprintSQL(string(line))
		fmt.Println(fp.StatementFP)
	}
}
