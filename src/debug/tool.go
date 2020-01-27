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
	//Map to keep track of visited node files so it doesn't print start flags for every log file
	nodeVisited := make(map[string]bool)

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

				//Currently this is printing start flags for every log file

				if nodeVisited[strings.Split(f.Name,"/")[2]] == false {
					if err := getStartFlags(f); err != nil {
						return err
					}
					nodeVisited[strings.Split(f.Name,"/")[2]] = true
				}



			case schemaFileRegExp.MatchString(f.Name):
				if !strings.HasPrefix(f.Name,"debug/schema/system") && !strings.Contains(f.Name,"@")   {
					//fmt.Printf("database: %s:\n", strings.Split(f.Name,"/")[2])
					//fmt.Printf("%s","blah")
					//if err := getDDL(f); err != nil {
					//	return err
					//}
				}
			//case f.Name == "debug/crdb_internal.jobs.txt":
			//	if err := getJobs(f); err != nil {
			//		return err
			//	}
			//case f.Name == "debug/events.json":
			//	if err := getEvents(f); err != nil {
			//	return err
			//	}
			case f.Name == "debug/settings.json":
				settings,err := getClusterSettingsAndCompare(f)
				if err != nil {
					return err
				}

				if len(settings) != 0 {
					fmt.Println("Changed Cluster Settings from debug.zip:")
					//Search for these settings in the log files
					for _,file := range r.File {
						if logFileRegExp.MatchString(file.Name){
							//fmt.Println(file.Name)
							err := searchForChangedClusterSettings(settings,file)
							if err != nil {
								return err
							}
						}
					}
				}

			default:

		}



	}
	return nil
}

func searchForChangedClusterSettings(clusterSettings []string,f *zip.File) error {

	settingStringForRegex := ""
	bar := ""
	for i := range clusterSettings {
		settingStringForRegex += bar + clusterSettings[i]
		bar = "|"
	}

	//fmt.Println(settingStringForRegex)

	re := regexp.MustCompile(settingStringForRegex)
	rff, err := f.Open()

	if err != nil {
		return err
	}
	defer rff.Close()

	scanner := bufio.NewScanner(rff)
	for scanner.Scan() {
		if len(re.FindString(scanner.Text())) != 0 {
			fmt.Printf("Found Cluster Setting Changed: %s\n\n",scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}





type localClusterSettings struct {
	KeyValues map[string]setting `json:"key_values"`
}

type setting struct {
	Value string `json:"value"`
	Type string `json:"type"`
	Description string `json:"description"`
}

func getClusterSettingsAndCompare (f *zip.File) ([]string,error) {
	rff, err := f.Open()

	if err != nil {
		return nil,err
	}

	contents, err := ioutil.ReadAll(rff)

	if err != nil {
		return nil,err
	}


	var sets localClusterSettings
	err = json.Unmarshal(contents, &sets)

	if err != nil {
		fmt.Println(err)
	}


	var baseSettings localClusterSettings
	err = json.Unmarshal([]byte(DefaultClusterSettings19_2_2), &baseSettings)

	var setArr []string
	for baseSetting, baseValue := range baseSettings.KeyValues {
		userValue, exists := sets.KeyValues[baseSetting]

		if !exists {
			fmt.Printf("%s is not present in the settings.json!!!\n", baseSetting)
			continue
		}
		if userValue.Value != baseValue.Value {
			//fmt.Printf("setting: %s\n", baseSetting)
			setArr = append(setArr, baseSetting)
			fmt.Printf("%s is set to %s, default is %s\n", baseSetting, userValue.Value, baseValue.Value)

		}
	}
	//fmt.Println(f.Name)
	//for i := range setArr {
	//	fmt.Println(setArr[i])
	//
	//}

	for userSetting, userValue := range sets.KeyValues {
		_, exists := sets.KeyValues[userSetting]
		if !exists {
			fmt.Printf("%s is a setting that shouldn't exist!!!!, it's value is %+v\n", userSetting, userValue)
			continue
		}
	}

	return setArr,nil
}


type eventJson struct {
	Events []event `json:"events"`
}
type event struct {
	Timestamp string `json:"timestamp"`
	Event_type string `json:"event_type"`
	Target_id int `json:"target_id"`
	Reporting_id int `json:"reporting_id"`
	Info string `json:"info"`
	Unique_id string `json:"unique_id"`

}
func getEvents(f *zip.File) error {
	rff, err := f.Open()

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(rff)

	if err != nil {
		return err
	}


	var e eventJson
	err = json.Unmarshal(data, &e)

	if err != nil {
		fmt.Println(err)
	}
	var e1 = "node_join"
	var e2 = "node_restart"
	var e3 = "node_decommissioned"
	var e4 = "node_recommissioned"
	fmt.Println("Node related events from events.json:")

	for _, event := range e.Events {

		if (event.Event_type == e1) || (event.Event_type == e2) || (event.Event_type == e3) || (event.Event_type == e4) {
			fmt.Printf("%s n%d at %s %s\n",event.Event_type,event.Target_id,event.Timestamp[:10], event.Timestamp[12:])
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



func main()  {
	var argsWithoutProg = os.Args[1:]

	fmt.Println(argsWithoutProg[0])
	err := openZip(argsWithoutProg[0])
	if err != nil {
		log.Fatal(err)
	}

}
