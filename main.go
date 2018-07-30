package main

import (
	"fmt"
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	nBytes, nChunks := int64(0), int64(0)
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)
	for {
		n, err := r.Read(buf[:cap(buf)])
		s := string(buf[:n])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nChunks++
		nBytes += int64(len(buf))
		// process buf
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		fmt.Println(s)
	}
	log.Println("Bytes:", nBytes, "Chunks:", nChunks)
}
