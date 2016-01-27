package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func printPaths(v interface{}, key, path string) {
	switch v := v.(type) {
	case map[string]interface{}:
		for mk, mv := range v {
			p := path + "." + mk
			if mk == key {
				fmt.Println(p)
			}
			printPaths(mv, key, p)
		}
	case []interface{}:
		for i, sv := range v {
			printPaths(sv, key, fmt.Sprintf("%s[%d]", path, i))
		}
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
		os.Exit(1)
	}
	key := flag.Arg(0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var v interface{}
		b := scanner.Bytes()
		if len(b) == 0 {
			continue
		}
		if err := json.Unmarshal(scanner.Bytes(), &v); err != nil {
			fmt.Fprintf(os.Stderr, "JSON parsing error: %s\n", err)
			continue
		}
		printPaths(v, key, "")
	}
	if err := scanner.Err(); err != nil {
		fatalln(err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `usage:
  %s key
where key is the key to search for in JSON structures passed to standard input.
`, os.Args[0])
}

func fatalln(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}
