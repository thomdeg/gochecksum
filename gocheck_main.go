package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	//"crypto/sha1"
	"hash"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/md5"
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

	// array with hashes and hash-names
	hashDescArr := [...]struct{ 
		desc string
		h hash.Hash 
	}{
		{"md5   :", md5.New()},
		{"sha1  :", sha1.New()},
		{"sha224:", sha256.New224()},
		{"sha256:", sha256.New()},
		{"sha384:", sha512.New384()},
		{"sha512:", sha512.New()},
	}

	// we need a io.Writer-Slice because the 3dot-operator
	// needs the correct type to unpack the slice later
	wrAr    := make( []io.Writer, 0, len(hashDescArr)  )
	for i:=0; i < len(hashDescArr); i++ {
		wrAr = append( wrAr,  hashDescArr[i].h  )
	}

	// copy infile to all the hashes:
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println("file:", filename)

	// setup multiWriter,  3dots instead of: io.MultiWriter(h224, h256, h512)
	mw := io.MultiWriter( wrAr... )
	
	// copy input file to multiwriter 
	if _, err := io.Copy(mw, f); err != nil {
		log.Fatal(err)
	}
	
	// Results:
	le := len(hashDescArr)
	resultArr := make( [][]byte, 0, le )
	for i:=0; i<le; i++ {
		resultArr = append( resultArr,  hashDescArr[i].h.Sum(nil) )
	}

	for i,_ := range resultArr  {
		fmt.Println( hashDescArr[i].desc , 
			hex.EncodeToString( resultArr[i] ))
	}

	fmt.Println("-------------")
	// we don't need our test
	//mySha256Test()
	//fmt.Println("-------------")
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