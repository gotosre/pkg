package main

import (
	"fmt"

	"github.com/gotosre/pkg/file"
	"github.com/gotosre/pkg/runner"
	"github.com/gotosre/pkg/slice"
)

func main() {
	a := []int64{1, 3, 5, 56, 67, 7}
	b := []int64{3, 5, 6, 7}
	fmt.Println(slice.SubInt64(a, b))

	runner.Init()
	fmt.Println(runner.Hostname)
	fmt.Println(runner.Cwd)

	fmt.Println(file.FilesUnder(runner.Cwd))
}
