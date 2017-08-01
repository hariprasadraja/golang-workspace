package main

import (
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("sh", "/home/parvathavarthinik/kpm/goworkspace/src/Hariprasad/Commands/run.sh")
	res, err := cmd.Output()
	log.Printf("Res : %+v, Error : %+v", string(res), err)

}
