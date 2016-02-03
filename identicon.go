package main

import (
    "crypto/md5"
    "flag"
    "fmt"
)

// Hashed string
var hashString string
func init() {
    flag.StringVar(&hashString, "hs", "Hello", "String to hash")
}

func main() {

    // Get string to hash
    flag.Parse()

    hashed := md5.New()
    hashed.Write([]byte(hashString))
    byteString := hashed.Sum(nil)

    fmt.Println("The string to hash: ", hashString)
    fmt.Printf("Hash in hex: %x\n", byteString)
    fmt.Println("Hash in bytes: ", byteString)
}
