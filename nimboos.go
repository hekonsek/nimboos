package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"os"
)

func main() {
	if os.Args[1] == "env" && os.Args[2] == "create" {
		cacheAnsibleCommand("env-create")

		env := os.Args[3]
		cmd := exec.Command("ansible-playbook","/tmp/nimboos-env-create.yml", fmt.Sprintf("-e vpc_name=%s", env))
		stdoutStderr, err := cmd.CombinedOutput()
		exitOnError(err, string(stdoutStderr))
		fmt.Println(string(stdoutStderr))
	} else if os.Args[1] == "env" && os.Args[2] == "remove" {
		cacheAnsibleCommand("env-remove")

		env := os.Args[3]
		cmd := exec.Command("ansible-playbook","/tmp/nimboos-env-remove.yml", fmt.Sprintf("-e vpc_name=%s", env))
		stdoutStderr, err := cmd.CombinedOutput()
		exitOnError(err, string(stdoutStderr))
		fmt.Println(string(stdoutStderr))
	} else {
		fmt.Println("Unknown command.")
	}
}

// Ansible

func cacheAnsibleCommand(command string)  {
	yml, err := Asset(fmt.Sprintf("data/%s.yml", command))
	exitOnError(err)
	err = ioutil.WriteFile(fmt.Sprintf("/tmp/nimboos-%s.yml", command), yml, 0644)
	exitOnError(err)
}

// Utils

func exitOnError(err error, context... interface{}) {
	if err != nil {
		fmt.Println(err)
		if len(context) > 0 {
			fmt.Println(context)
		}
		os.Exit(-1)
	}
}