package MgoDb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	//"lingapos/server/db"
	"log"

	"fmt"
)

var customer = map[string]interface{}{
"name" : "customer 1",
"firstName" : "",
"lastName" : "",
"gender" : "",
"phoneNumber" : "8098994302",
"emailId" : "hariprasad@benseron.com",
"imageAvailable" : false,
"imageVersion" : 0,
"gateCode" : "",
"loyaltyPoints" : 0,
"account" : bson.ObjectIdHex("5936895978362e42ac5e02ce"),
"activeStatus" : true,
"customerAddress" : []string{},
"dateCreated" : time.Now(),
"lastUpdated" : time.Now(),
"customerId" : "",
"cardDetails" : []string{},
"taxExempt" : false,
"notes" : "",
"stores" : []bson.ObjectId{bson.ObjectIdHex("5936895978362e42ac5e02cc")},
}

func main(){
	start := time.Now()
	for i:=1;i<=3000;i++{
        customer["firstName"] = fmt.Sprintf("%d thCustomer",i)
		customer["emailId"] = fmt.Sprintf(" %d thCustomer@Customer.com",i)
		number :=   i + 1000000000
		customer["phoneNumber"] = fmt.Sprint(number)
		customer["_id"] = bson.NewObjectId()
		//err := db.Obj.C("customer").Insert(customer)
		//if err != nil {
		//	log.Println("err",err.Error())
		//}

	}
	end := time.Now()
	log.Println(end.Sub(start).Seconds())
	log.Println("complete")


}
