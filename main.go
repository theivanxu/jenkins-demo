package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println("Hello, Kubernetes!I'm from Jenkins CI!")
	fmt.Println("BRANCH_NAME:", os.Getenv("branch"))
}