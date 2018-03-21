package main

import (
	"crypto/sha1"
	"fmt"
	"encoding/json"
)

func main()  {

		s := make(map[string]string)

	    s["name"] = "hari"
	    s["roll"] = "dev"



	    h := sha2.New()
	    data1,_ := json.Marshal(s)
		h.Write(data1)
		bs := h.Sum(nil)
		fmt.Printf("%x\n", bs)

	    d := make(map[string]string)
		d["roll"] = "dev"
	    d["name"] = "hari"
	    data2,_ := json.Marshal(d)
		h = sha1.New()
		h.Write(data2)
		ss := h.Sum(nil)
		fmt.Printf("%x\n", ss)




}