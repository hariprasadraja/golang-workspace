package main

import (
	//"time"
	"log"
	//"net/http"
	//"encoding/json"
	//"io/ioutil"
	//"time"
	"time"
	//"net/http"
	//"io/ioutil"
	//"encoding/json"
	//"errors"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {

	//res := make(map[string]interface{})

	//var ok <-chan time.Time
	//data := time.Now()
	//ok <- data
	for range time.Tick(time.Second * 10) {
		resp, err := http.Get(url)
		if err != nil {
			log.Print(err.Error())
		}
		if resp != nil {
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}

			res := make(map[string]interface{})
			json.Unmarshal(data, res)
			for key, val := range res {
				log.Println(key, ":", val)
			}

		} else {
			log.Println("empty")
		}

	}
}

//}
