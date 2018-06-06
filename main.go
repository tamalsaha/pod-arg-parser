package main

import (
	"io/ioutil"
	"log"
	"github.com/ghodss/yaml"
	"fmt"
	"strings"
	"os"
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

	cmds := pod.Spec.Containers[0].Command

	for _, c := range cmds {
		if strings.Contains(c, "kube-apiserver") {
			fields := strings.Fields(c)

			var idx int
			for i, w := range fields {
				if strings.HasSuffix(w, "kube-apiserver") {
					idx = i
					break
				}
			}

			var args []string

			for _, w := range fields[idx+1:] {



			}

			
			fmt.Println(strings.Join(fields, "\n"))
			os.Exit(0)
		}
	}

	fmt.Println(pod.Spec.Containers[0].Command)

	// kube-apiserver
}
