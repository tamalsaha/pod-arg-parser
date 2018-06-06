package main

import (
	"io/ioutil"
	"log"
	"github.com/ghodss/yaml"
	"fmt"
)

func main() {
	data, err := ioutil.ReadFile("pod.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	var pod Pod
	err = yaml.Unmarshal(data, &pod)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(pod.Kind)
}
