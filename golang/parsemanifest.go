package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
	"encoding/json"
)

// convertMapI2MapS a recursive function to parse multi depth yaml file
func convertMapI2MapS(i interface{}) interface{} {
    switch x := i.(type) {
    case map[interface{}]interface{}:
        m2 := map[string]interface{}{}
        for k, v := range x {
            m2[k.(string)] = convertMapI2MapS(v)
        }
        return m2
    case []interface{}:
        for i, v := range x {
            x[i] = convertMapI2MapS(v)
        }
    }
    return i
}

type Testcases struct {
	Description string `json:"description"`
	Tests []string `json:"testcases"`
}

func ParseManifestAndLaunchRunTestCases(manifest string) (error) {
	// body interface to store first level yaml
	var body interface{}
	var testlist Testcases

    yamlFile, err := ioutil.ReadFile(manifest)
    if err != nil {
        log.Printf("yamlFile :%s parsing erro #%v ", manifest, err)
    }
    err = yaml.Unmarshal(yamlFile, &body)
    if err != nil {
        log.Fatalf("Manifest yamlfile: %s unmarshal: %v", manifest, err)
    }

	body = convertMapI2MapS(body)
	//fmt.Printf("body=%s\n", body)

    if b, err := json.Marshal(body); err != nil {
        panic(err)
    } else {
		//fmt.Printf("Output:%s\n", b)
		json.Unmarshal(b, &testlist)
		//fmt.Printf("json.Unmarshal=: %s\n", testlist)
		fmt.Printf("Test.Description:%s\n", testlist.Description)
		fmt.Printf("Total number of Test cases:%d\n", len(testlist.Tests))
		for v, t := range testlist.Tests {
			fmt.Printf("test cases dirs: %d->%s\n",v, t)
			if v > 0 {
				fmt.Printf("Prev Test case: %d->%s\n",v, testlist.Tests[v -1])
			}
		}
        return nil
    }
}

func main() {
	fmt.Printf("Input: %s\n", "manifest.yml")
    _ =ParseManifestAndLaunchRunTestCases("manifest.yml")
}
