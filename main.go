// this program will create a hash of stdin
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var (
	// Command line flags.
	s384    bool
	s512    bool
	example bool
	version = "devel" // for -v flag, updated during the release process with -ldflags=-X=main.version=...
)

func init() {
	flag.BoolVar(&s384, "s384", false, "Create sha384 hash")
	flag.BoolVar(&s512, "s512", false, "Create sha512 hash")
	flag.BoolVar(&example, "example", false, "Print example usage")
	flag.Usage = usage
}

func printExampleUsage() {
	fmt.Println("url-pinger https://www.google.com")
	fmt.Println("url-pinger -delay 2 https://www.google.com")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] string\n\n", os.Args[0])
	fmt.Println("***Default hash: sha256")
	fmt.Fprintln(os.Stderr, "OPTIONS:")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if example {
		printExampleUsage()
		os.Exit(0)
	}

	if isFlagPassed("s384") {
		input := os.Args[2]
		hash := sha512.Sum384([]byte(input))
		pass := string(hash[:])
		fmt.Printf("%s: %x\n", input, pass)
	} else if isFlagPassed("s512") {
		input := os.Args[2]
		hash := sha512.Sum512([]byte(input))
		pass := string(hash[:])
		fmt.Printf("%s: %x\n", input, pass)
	} else {
		input := os.Args[1]
		hash := sha256.Sum256([]byte(input))
		pass := string(hash[:])
		fmt.Printf("%s: %x\n", input, pass)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
