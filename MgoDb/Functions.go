package MgoDb

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	//"net/http"
	//"encoding/json"
	//"io/ioutil"
	//"github.com/julienschmidt/httprouter"
	//"net/http"
)

type Data struct {
	Id    string          `bson:"_id"`
	Value bson.JavaScript `bson:"value"`
}

func main() {

	res := make(map[string]interface{})
	//router := httprouter.New()
	//router.Handle("POST","/check",handler)

	//if err := http.ListenAndServe(":1234",router); err != nil {
	//	log.Fatal("ListenAndServe:", err)
	//}
	//js := mongoNow()
	DB := connectDatabase()
	//data := Data{"add",js}

	//DB.C("StackOverFlow").Insert( struct{Value interface{}}{Value: mongoNow()})
	// DB.C("system.js").Insert(&data)

	err := DB.Run(bson.M{"eval": "add(5,5);"}, res)
	if err != nil {
		log.Println(err)
	} else {
		for key, val := range res {
			log.Println(key, " : ", val)
		}
	}
}

func mongoNow() bson.JavaScript {

	return bson.JavaScript{
		// place your function in here in string
		Code: "function add(a,b){return a+b}",
	}
}

func connectDatabase() *mgo.Database {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB("test")
}
