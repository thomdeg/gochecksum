package main

import (
	"fmt"
	"os"
	//"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"io"
	"log"
	"encoding/hex"
)


func main() {

	filename := ""
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if filename == "" {
		fmt.Println("filename is missing.")
		os.Exit(0)
	}

	// init hashes:
	h224 := sha256.New224()
	h256 := sha256.New()
	h512 := sha512.New()

	// copy infile to all the hashes:
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Copy file to a multiWriter
	mw := io.MultiWriter(h224, h256, h512)
	if _, err := io.Copy(mw, f); err != nil {
		log.Fatal(err)
	}

	
	// Results:
	h224sum := h224.Sum(nil)
	h256sum := h256.Sum(nil)
	h512sum := h512.Sum(nil)


	//fmt.Println("sha244:", h224sum)
	//fmt.Println("sha256:", h256sum)
	fmt.Println("sha244:", hex.EncodeToString(h224sum))
	fmt.Println("sha256:", hex.EncodeToString(h256sum))
	fmt.Println("sha512:", hex.EncodeToString(h512sum))

	fmt.Println("-------------")
	mySha256Test()
	fmt.Println("-------------")
}

func mySha256Test() {
	
	f, err := os.Open(os.Args[1])
	if err != nil {
	  log.Fatal(err)
	}
	defer f.Close()
  
	h256 := sha256.New()
	if _, err := io.Copy(h256, f); err != nil {
	  log.Fatal(err)
	}
  
	fmt.Printf("%x \n" , h256.Sum(nil))	
}
