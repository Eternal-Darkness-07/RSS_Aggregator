package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World");
	PortString := os.Getenv("PORT");
	
	if len(PortString) == 0 {
		log.Fatal("PORT env not found...!");
	}
	fmt.Println("PORT: ", PortString);
}