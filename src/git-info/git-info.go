package main

import (
  "fmt"
  "os"
  "path/filepath"
)

func main() {
	fmt.Println("git-info!")
  dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
  fmt.Println(dir)
}
