package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Print Hello, World!
	fmt.Println("Hello, World!")

	// Get Go dependencies using `go list -m all`
	cmd := exec.Command("go", "list", "-m", "all")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error fetching dependencies:", err)
		return
	}

	// Print dependencies line by line
	fmt.Println("\nGo Dependencies:")
	dependencies := strings.Split(string(output), "\n")
	for _, dep := range dependencies {
		if dep != "" {
			fmt.Println(dep)
		}
	}
}
