package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)


func test(b string) {
	fmt.Println(b)
}

func getVersion(base string) {
	nodesDir := path.Join(base, "nodes")
	fmt.Println(nodesDir)
	files, err := ioutil.ReadDir(nodesDir)
}

func main() {
	var argsWithoutProg = os.Args[1:]

	fmt.Println(argsWithoutProg[0])

	test("blah")
	getVersion(argsWithoutProg[0])

}
