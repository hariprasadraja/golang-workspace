package main

import (
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("sh", "run.sh")
	res, err := cmd.Output()
	log.Printf("Res : %+v, Error : %+v", string(res), err)

}
