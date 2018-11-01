package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	path := "examplefiles/"
	fileName :=  "world-cities.csv" //os.Args[1]
	content, err := ioutil.ReadFile(path +fileName)
	if err != nil {
		log.Fatal(err)
	}

	zipFileName := path + "results/" + fileName[:len(fileName)-4] + ".zip"
	fmt.Printf("Creating zip file: %s \n", zipFileName)

	f, err := os.Create(zipFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Writing content\n")
	_, err = f.Write([]byte(content))
	if err != nil {
		log.Fatal(err)
	}

	// Make sure to check the error on Close.
	if err = f.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Zip file %s was created\n", zipFileName)

}
