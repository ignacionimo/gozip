package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func main() {
	src, err := os.Open("world-cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dest, err := os.Create("next.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	zdest := zip.NewWriter(dest)
	wr, err := zdest.Create("_world-cities.csv")
	if _, err := io.Copy(wr, src); err != nil {
		log.Fatal(err)
	}
	err = zdest.Close()
	log.Println("ERR: ", err)
}
