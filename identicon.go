package main

import (
    "crypto/sha1"
    "flag"
    "fmt"
)

var hashString string
func init() {
    flag.StringVar(&hashString, "hashs", "Hello", "String to hash")
}

func main() {

    flag.Parse()

    hashed := sha1.New()
    hashed.Write([]byte(hashString))
    byteString := hashed.Sum(nil)
    fmt.Println("The string to hash: ", hashString)
    fmt.Printf("Hashed: %x\n", byteString)
}
