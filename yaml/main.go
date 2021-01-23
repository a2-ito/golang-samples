package main

import (
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
	"fmt"
  pretty "github.com/kr/pretty"
)

const conf = "./config.yaml"

func main() {

  configbuf, err := ioutil.ReadFile(conf)
  if err != nil {
    fmt.Println(err)
    return
  }

  //data, err := ReadOnSliceMap(configbuf)
  yamldata := make(map[interface{}]interface{})
  err = yaml.Unmarshal(configbuf, &yamldata)
  if err != nil {
    fmt.Println(err)
    return
  }
	pretty.Print("--- yamldata:\n%# v\n\n", yamldata)

	fmt.Println()
	//fmt.Println(yamldata["hoge"].(string))
	roles := []string{}
  for _, value := range yamldata["a"].(map[interface {}]interface {})["b"].(map[interface {}]interface {})["roles"].([]interface {}) {
		roles = append(roles, value.(string))
	}
	//fmt.Println(yamldata["roles"])
	fmt.Println(roles)

}
