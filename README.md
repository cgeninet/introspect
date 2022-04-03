# Go Introspection
Yaml file (test.yaml) :
    
    ---
    a: 1

    vars:
      hello: world
      number: 2

    1.2:
      - str1a: "string 1 a"
        str1b: "string 1 b"
      -
        - a
        - b
        - c
      - str2a: "string 2 a"
        str2b:
      - "string 3"
   
Code :

    package main

    import (
        "fmt"
        "io/ioutil"
	    "gopkg.in/yaml.v3"

        "github.com/cgeninet/introspect"
    )

    func main() {
        // Load Yaml file
        yamlFile, _ := ioutil.ReadFile("test.yaml")
        data := make(map[interface{}]interface{})
        yaml.Unmarshal(yamlFile, &data)

        is := introspect.NewStruct(data)
        for _, k := range is.Keys() {
            fmt.Println("path", k, "has value :", is.Value(k))
        }
    }

Result :

    path /1.2/2/str2a has value : string 2 a
    path /1.2/3 has value : string 3
    path /1.2/1/0 has value : a
    path /1.2/1/1 has value : b
    path /vars/number has value : 2
    path /1.2/0/str1a has value : string 1 a
    path /1.2/0/str1b has value : string 1 b
    path /1.2/1/2 has value : c
    path /a has value : 1
    path /vars/hello has value : world
