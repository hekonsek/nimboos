package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"os"
)

func main() {
	envCreateYml, err := Asset("data/env-create.yml")
	exitOnError(err)

	err = ioutil.WriteFile("/tmp/nimboos-env-create.yml", envCreateYml, 0644)
	exitOnError(err)

	cmd := exec.Command("ansible-playbook", "/tmp/nimboos-env-create.yml", "-e vpc_name=nimboos_test1")
	stdoutStderr, err := cmd.CombinedOutput()
	exitOnError(err)
	fmt.Println(string(stdoutStderr))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}