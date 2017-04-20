package main

import (
	"log"
	"sort"
	"strings"
)

func main()  {
	Names := []string{
		"vinoth",
		"hari",
		"kerishna",
		"Vinoth",
		"hari",
		"balaji",
		"xyz",
		"12345",
		"poipu",
		"5456565",
		"hari",
		"prasanth",
		"rajaesh",
		"xyz",
		"prasanth",
	}
	var same []string
	var diff []string
	sort.Strings(Names)
	 for _,Pointer := range Names{
		 count :=0
          for _,value := range Names {
			  if strings.EqualFold(Pointer, value){
                count = count + 1
			  }
		  }
		 if(count > 1){
			 diff = append(diff,Pointer)
		 }else{
			 same = append(same,Pointer)
		 }
       log.Println(count)

	 }
	log.Println(diff)
	log.Println(same)
	}
