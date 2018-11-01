package main

import (
	"compress/zlib"
	"io"
	"log"
	"os"
)

func main() {
	src, err := os.Open("examplefiles/source.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dest, err := os.Create("examplefiles/results/next.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()
	zdest := zlib.NewWriter(dest)
	defer zdest.Close()
	if _, err := io.Copy(zdest, src); err != nil {
		log.Fatal(err)
	}
}
