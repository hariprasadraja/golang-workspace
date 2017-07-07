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
)

//const(
//	GraveDevil  =	30101
//    stagingAdmin = 58eb24a49230ca6124a58699
//)

func main() {

	//res := make(map[string]interface{})

	//grave Devil production
	//url := "http://admin.zenpepper.com/store/30101/placedOrders?token=3fe24a4e-5e9c-453d-9f70-42550d82b840" // grave Devil Production
	//url := "https://stagingadmin.zenpepper.com/store/58eb24a49230ca6124a58699/placedOrders?token=170e30de-163e-427e-978c-91a53f5bcf2c"

	//url := "https://stagingadmin.zenpepper.com/store/58eb24a49230ca6124a58699/placedOrders?token=170e30de-163e-427e-978c-91a53f5bcf2c"

	var ok <-chan time.Time
	data := time.Now()
	ok <- data
	for ok {

		log.Println("check", ok)
		ok = Tick(time.Minute * 10)

		//resp, _ := http.Get(url)
		//if (resp != nil) {
		//	data, err := ioutil.ReadAll(resp.Body)
		//	if err != nil {
		//		log.Println(err)
		//	}
		//
		//	json.Unmarshal(data, res)
		//	for key, val := range res {
		//		log.Println(key, ":", val)
		//	}
		//
		//} else {
		//	log.Println("empty")
		//}
		//
	}
}

//}
