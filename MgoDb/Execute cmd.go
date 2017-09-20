package MgoDb

import (
	"log"
	"os/exec"
)

func main() {
	command := "mongorestore --db Linga LingaDb"
	command1 := "cd dump"
	cmd := exec.Command(command1)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	cmd = exec.Command(command)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}

}
