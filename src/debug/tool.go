package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	//"path"
	//"path/filepath"
)

func openZip(file string) error {
	logFileRegExp := regexp.MustCompile(`([a-z]{5}/[a-z]{5}/\d+)/logs/.*.log`)
	schemaFileRegExp := regexp.MustCompile(`([a-z]{5}\/[a-z]{6}\/[[:word:]]*)`)

	// Open a zip archive for reading.
	r, err := zip.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()


	//This is tracking whether we found one log file
	//One log file is sufficient to return the version
	var foundLog bool = false

	for _, f := range r.File {
		//if f.FileInfo().IsDir(){
		//	fmt.Println(f.Name)
		//}
		//fmt.Println(f.Name)

		switch {
			//This case is to find one log file to print the version and cluster ID
			case logFileRegExp.MatchString(f.Name):
				if !foundLog {
					if err := getVersion(f); err != nil {
						return err
					}
					if err := getClusterID(f); err != nil {
						return err
					}
					foundLog = true
				}

				if err := getStartFlags(f); err != nil {
					return err
				}
			case schemaFileRegExp.MatchString(f.Name):
				if !strings.HasPrefix(f.Name,"debug/schema/system") && !strings.Contains(f.Name,"@")   {
					fmt.Printf("database: %s:\n", strings.Split(f.Name,"/")[2])
					if err := getDDL(f); err != nil {
						return err
					}
				}
			case f.Name == "debug/crdb_internal.jobs.txt":
				if err := getJobs(f); err != nil {
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

func getJobs(f *zip.File) error {
	fmt.Println("NON SUCCEEDED JOBS:")
	re := regexp.MustCompile(`succeeded`)

	rff, err := f.Open()

	if err != nil {
		return err
	}
	defer rff.Close()

	scanner := bufio.NewScanner(rff)

	var gotCols bool = false
	var gotDashes bool = false

	for scanner.Scan() {
		if !gotCols {
			var cols = strings.Split(scanner.Text(),"|")
			fmt.Println(cols[0] + "|" + cols[1] + "|" + cols[2] + "|" + cols[6] + "|" + cols[14] + "\n")
			gotCols = true
		} else if !gotDashes {
			var dashes = strings.Split(scanner.Text(),"+")
			fmt.Println(dashes[0] + "+" + dashes[1] + "+" + dashes[2] + "+" + dashes[3] + "+" + dashes[7] + "+" + dashes[15] + "\n")
			gotDashes = true
		} else {
			if !re.MatchString(scanner.Text()) {
				data := strings.Split(scanner.Text(), "|")
				//Received a "panic: runtime error: index out of range" when reaching the last line of this file
				//Workaround is checking the length of data, if it's 1 then EOF
				if len(data) != 1 {
					fmt.Println(data[0] + "|" + data[1] + "|" + data[2] + "|" + data[6] + "|" + data[14] + "\n")
				}
			}



		}


	}


	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func getDDL(f *zip.File) error {
	//fmt.Println(f.Name)
	rff, err := f.Open()

	if err != nil {
		return err
	}

	contents, err := ioutil.ReadAll(rff)

	if err != nil {
		return err
	}


	var m map[string]interface{}
	err = json.Unmarshal(contents, &m)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", m["create_table_statement"])


	return nil

}


func getClusterID(f *zip.File) error {
	re := regexp.MustCompile(`clusterID: [0-9a-f]{8}\-[0-9a-f]{4}\-[0-9a-f]{4}\-[0-9a-f]{4}\-[0-9a-f]{12}`)
	rff, err := f.Open()

	if err != nil {
		return err
	}
	defer rff.Close()

	scanner := bufio.NewScanner(rff)
	for scanner.Scan() {
		if len(re.FindString(scanner.Text())) != 0 {
			fmt.Printf("%s\n",re.FindString(scanner.Text()))
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func getStartFlags(f *zip.File) error {
	re := regexp.MustCompile(`arguments: \[(.*?)]`)
	rff, err := f.Open()

	if err != nil {
		return err
	}
	defer rff.Close()

	scanner := bufio.NewScanner(rff)
	for scanner.Scan() {
		if len(re.FindString(scanner.Text())) != 0 {
			fmt.Printf("node %s %s\n",strings.Split(f.Name,"/")[2],re.FindString(scanner.Text()))
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func getVersion(f *zip.File ) error {
	re := regexp.MustCompile(`v[0-9]+.[0-9].[0-9]`)

	rff, err := f.Open()

	if err != nil {
		return err
	}
	defer rff.Close()

	scanner := bufio.NewScanner(rff)
	for scanner.Scan() {
		if len(re.FindString(scanner.Text())) != 0 {
			fmt.Printf("CockroachDB %s\n",re.FindString(scanner.Text()))
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

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


}
