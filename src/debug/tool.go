package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"path"
	//"path/filepath"
)

func openZip(file string) error {
	// Open a zip archive for reading.
	r, err := zip.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	//close this later at some point, when im done with the actual contents of the zip file
	//defer r.Close()

	for _, f := range r.File {
		//if f.FileInfo().IsDir(){
		//	fmt.Println(f.Name)
		//}
		fmt.Println(f.Name)

		switch {
			case f.Name == "root/nodes/1/log.txt":
				if err := readFromFile(f); err != nil {
					return err
		  		}
		  	case f.Name == "root/nodes/2/log.txt":
				if err := readFromFile(f); err != nil {
					return err
				}
			case f.Name == "root/b.json":
				if err := readFromJson(f); err != nil {
					return err
				}
			default:

		}



	}
	return nil
}

func readFromJson(f *zip.File) error {
	rff, err := f.Open()

	if err != nil {
		return err
	}

	contents, err := ioutil.ReadAll(rff)

	if err != nil {
		return err
	}

	//fmt.Println(string(contents))

	//type model struct {
	//	KEY string `json:"key"`
	//	KEY2 string `json:"key2"`
	//}
	//
	//// json data
	//var obj model
	//
	//err = json.Unmarshal(contents, &obj)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//fmt.Printf("%s: %s", obj.KEY)
	//fmt.Printf("%s: %s", obj.KEY)

	var m map[string]interface{}
	err = json.Unmarshal(contents, &m)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("from file: %s\n",f.Name)
	for k,v := range m {
		fmt.Printf("%s : %s\n",k,v)
	}

	//can also index into the map instead of for loop and checking if k == "<mykey>"

	return nil
}

func readFromFile(f *zip.File ) error {
	rff, err := f.Open()

	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(rff)

	if err != nil {
		return err
	}

	fmt.Printf("from file: %s\n",f.Name)
	fmt.Println(string(buf))

	defer rff.Close()

return nil
}

//func getText(base string) {
//	nodesDir := path.Join(base, "nodes")
//	//fmt.Println(nodesDir)
//	files, err := ioutil.ReadDir(nodesDir)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, nodeNumber := range files {
//		if nodeNumber.Name() != ".DS_Store" {
//			nodeFolder := path.Join(nodesDir, nodeNumber.Name())
//			logFiles, err := ioutil.ReadDir(nodeFolder)
//			if err != nil {
//				log.Fatal(err)
//			}
//			for _, logfile := range logFiles {
//				fmt.Println(logfile.Name())
//				fullPathToLogFile := path.Join(nodeFolder, logfile.Name())
//				data, err := ioutil.ReadFile(fullPathToLogFile)
//				if err != nil {
//					log.Fatal(err)
//				}
//				fmt.Println(string(data))
//			}
//			//fmt.Println(nodeFolder)
//			//fmt.Println(f.Name())
//		}
//
//	}
//
//}

func main() {
	var argsWithoutProg = os.Args[1:]

	fmt.Println(argsWithoutProg[0])
	openZip(argsWithoutProg[0])
	//unzipped := openZip(argsWithoutProg[0])

	//dir, err := ioutil.TempDir("", "example")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//defer os.RemoveAll(dir) // clean up


	//fmt.Println(dir)

	// Iterate through the files in the archive,
	// printing some of their contents.
	//for _, f := range unzipped.File {
	//	//if (f.Name != ".DS_Store") {
	//	//
	//	//}
	//	fmt.Printf("%s:\n", f.Name)
	//	reader, err := f.Open()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	contents, err := ioutil.ReadAll(reader)
	//	tmpfn := path.Join(dir, f.Name)
	//	base := filepath.Dir(tmpfn)
	//
	//	if _, err := os.Stat(base); os.IsNotExist(err) {
	//		fmt.Println("made it here")
	//
	//		if err := os.MkdirAll(base, 0777); err != nil {
	//
	//			log.Fatal(err)
	//		}
	//	}
	//	if err := ioutil.WriteFile(tmpfn, contents, 0777); err != nil {
	//		log.Fatal(err)
	//	}
	//}

	//fmt.Println(dir)
	//debugDir := path.Join(dir, "debug")
	//getText(debugDir)

}
