package main

import (

	"path/filepath"
	"runtime"
	"time"
	"reflect"
	"log"
)

func main() {
	getFileName(1)
	time.Sleep(time.Hour)
}

func getFileName(shift int) {
	go func() {
		pc, file, line, ok := runtime.Caller(shift)
		log.Println(reflect.TypeOf(pc))
		var s []uintptr
		s = append(s,pc)
       log.Println(reflect.TypeOf(file))
		log.Println(file)
		log.Println(reflect.TypeOf(line))
		log.Println(line)
		if !ok {
			file = "???"
			line = 0
		} else {
			file = filepath.Base(file)
		}
        function :=runtime.FuncForPC(pc)
		log.Println("functionname",function.Name())
		log.Println()

		//fmt.Printf("%s:%d", file, line)
		//frames :=runtime.CallersFrames(s)


	}()
}

