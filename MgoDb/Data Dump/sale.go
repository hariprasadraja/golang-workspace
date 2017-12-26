package main

import (
	"time"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
)
var modifier = []map[string]interface{}{}
var tax =[]map[string]interface{}{
	{
		"id":             bson.ObjectIdHex("56e25d8f08811a14c835b76f"),
		"name":           "Service tax",
		"taxRate":        1500,
		"totalTaxAmount": 450,
		"inclusive":      false,
		"taxes":          []map[string]interface{}{},
	},
}
var alltax = []map[string]interface{}{}
var applieddiscount = []map[string]interface{}{}
var tax1 =[]map[string]interface{}{
	{
		"id":                bson.ObjectIdHex("56e25d8f08811a14c835b76f"),
		"name":              "Service tax",
		"taxRate":           1500,
		"taxRateStr":        "",
		"totalTaxAmount":    4800,
		"totalTaxAmountStr": "",
		"inclusive":         false,
	},
}
var applieddiscount1 = []map[string]interface{}{}
var appliedcheckdiscount = []map[string]interface{}{}
var payment = []map[string]interface{}{}
var freeitem = []map[string]interface{}{}
var giftcard = []map[string]interface{}{}
var refund =[]map[string]interface{}{}
var ordres = []map[string]interface{}{
	{

		"_id": bson.ObjectIdHex("56e2ce6008811a2d09715088"),
		"isNew": false,
		"seat": "",
		"orderNo": int(0),
		"menuType": "",
		"measureType": "",
		"menuItem":               bson.ObjectIdHex("5968c5a148238129caac01ed"),
		"menuName":               "soft vada",
		"category":               bson.ObjectIdHex("584a4e1744c9cb1aa77a1a46"),
		"subCategory":            bson.ObjectIdHex("5884961b44c9cb0ef48ab812"),
		"categoryLevel":          2,
		"department":             bson.ObjectIdHex("584a4e1744c9cb1aa77a1a34"),
		"course":                 bson.ObjectIdHex("584a4e1744c9cb1aa77a1a2a"),
		"servingSize":            bson.ObjectIdHex("584a4e1744c9cb1aa77a1a3a"),
		"giftCardNumber":         "",
		"giftCardReference":      "",
		"haAccountNumber":        "",
		"quantity":               int(1),
		"displayQuantity":        "1",
		"modifiers":              modifier,
		"isVoid":                 true,
		"isItemToGo":             false,
		"voidError":              "Not good",
		"openItem":               false,
		"employee":               bson.ObjectIdHex("56e1c14208811a14c835b73c"),
		"sliceName":              "",
		"netSales":               0,
		"grossAmount":            3000,
		"totalPrice":             3000,
		"totalGrossAmount":       3000,
		"totalTaxAmount":         450,
		"modifierTotalTaxAmount": 0,
		"taxes":                  tax,
		"allTaxes":               alltax,
		"appliedDiscounts":       applieddiscount,
		"checkDiscountAmount":    0,
		"checkDiscountTax":       0,
		"itemDiscountAmount":     0,
		"itemDiscountTax":        0,
		"totalDiscountAmount":    0,
		"totalDiscountTax": 0,
		"totalDiscounts": 0,
		"totalAmount": 3450,
		"totalInclusiveTaxAmount": 0,
		"menuInclusiveDiscount": 0,
		"modifierInclusiveDiscount": 0,
		"cutAndModify": false,
		"fraction": 1,
		"cutFraction": 1,
		"splitQuantity": 0,
		"originalGrossAmount": 0,
		"lastSeatOrderSplit": false,
	},

}

var sale = map[string]interface{}{

	/* 1 */

	"New object Id" : bson.ObjectIdHex("56e2ce6008811a2d09715087"), // New Object Id
	"saleNo" : "A000-",    // A00 - Index
	"orders" : ordres,
	"grossSales" : 100,           // 100
	"grossSalesStr" : "",
	"grandSales" : 200,               //200
	"grandSalesStr" : "",
	"grossReceipt" : 300,             //300
	"grossReceiptStr" : "",
	"totalTaxAmount" : 25,           // 25
	"totalTaxAmountStr" : "",
	"inclusiveTaxAmount" : 50,       // 50
	"inclusiveTaxAmountStr" : "",
	"taxExempt" : false,           // 25
	"store" : bson.ObjectIdHex("56e1c14208811a14c835b73b"),
	"serviceCharge" : 10,   //10
	"serviceChargeStr" : "",
	"netSales" : 5000,     // 5000
	"netSalesStr" : "",
	"tipPercentage" : 10,   //10
	"tipPercentageStr" : "",
	"tipTotalAmount" : 500,    // 500
	"tipTotalAmountStr" : "",
	"grossVoid" : 3000,    // 3000
	"grossVoidStr" : "",
	"netVoid" : 2000,      // 2000
	"netVoidStr" : "",
	"taxVoid" : 4000,       // 4000
	"taxVoidStr" : "",
	"voidDiscounts" : 0,
	"voidDiscountsStr" : "",
	"voidDiscountTax" : 0,
	"voidDiscountTaxStr" : "",
	"discounts" : 0,
	"discountsStr" : "",
	"discountTax" : 0,
	"discountTaxStr" : "",
	"checkDiscountAmount" : 0,
	"checkDiscountAmountStr" : "",
	"checkDiscountTax" : 0,
	"checkDiscountTaxStr" : "",
	"itemDiscountAmount" : 0,
	"itemDiscountAmountStr" : "",
	"itemDiscountTax" : 0,
	"itemDiscountTaxStr" : "",
	"paidAmount" : 5000,          //5000
	"paidAmountStr" : "",
	"balanceAmount" : 300,       // 300
	"balanceAmountStr" : "",
	"taxExemptAmount" : 0,
	"taxExemptAmountStr" : "",
	"taxes" : tax1,
	"appliedDiscounts" : applieddiscount1,
	"appliedCheckDiscounts" : appliedcheckdiscount,
	"payments" : payment,
	"giftCardAmount" : 0,
	"giftCardAmountStr" : "",
	"guestCount" : 1,
	"cashAmount" : 0,
	"cashAmountStr" : "",
	"sideCC" : 0,
	"sideCCStr" : "",
	"creditCardAmount" : 0,
	"creditCardAmountStr" : "",
	"giftCardSoldAmount" : 0,
	"giftCardSoldAmountStr" : "",
	"employee" : bson.ObjectIdHex("56e1c14208811a14c835b73c"),
	"isNewCustomer" : false,
	"saleDate" : time.Now(),
	"saleTime" : "19:22",
	"saleHour" : 19,
	"saleMinute" : 22,
	"saleOpenTime" : "19:16",
	"saleOpenHour" : 19,
	"saleOpenMinute" : 0,
	"createdBy" : bson.ObjectIdHex("56e1c14208811a14c835b73c"),
	"dateCreated" : time.Now(),
	"node" : "A000",
	"startDate" : time.Now().Add(-2*time.Hour),
	"closeDate" : time.Now(),
	"syncWithInventory" : false,
	"feedback" : "No feedback",
	"freeItems" : freeitem,
	"totalTaxAmountWithOutDiscounts" : 0,
	"totalTaxAmountWithOutDiscountsStr" : "",
	"tableNo" : "QSR",
	"seatNo" : "1",
	"isVoidCheck" : true,
	"giftCards" : giftcard,
	"changeDue" : 0,
	"changeDueStr" : "",
	"totalInclusiveDiscountTax" : 0,
	"totalInclusiveDiscountTaxStr" : "",
	"totalInclusiveCheckDiscountTax" : 0,
	"totalInclusiveCheckDiscountTaxStr" : "",
	"syncWithCashier" : false,
	"refunds" : refund,
	"ticketNo" : "00-001",
	"closeDay" : true,
	"totalVoidInclusiveTax" : 0,
	"totalVoidInclusiveTaxStr" : "",

}
func main() {

	sale = map[string]interface{}{
		//"_id" : bson.ObjectIdHex("592c1eb3482381351656c54e"),
		"account" : bson.ObjectIdHex("584a4dee44c9cb1aa77a1a12"),
		"store" : bson.ObjectIdHex("584a4dee44c9cb1aa77a1a10"),
		"appversion" : "2.0.3(0.1)Go Local Test",
		"employee" : bson.ObjectIdHex("584a4dee44c9cb1aa77a1a11"),
		"reopenVersion" : 0,
		"saleUniqueId" : "f40ef83b06f74ec0a55b54171f292a85",
		"saleNo" : "A01-290517013",
		"ticketNo" : "1-013",
		"node" : "A01",
		"floorId" : "5884ac2244c9cb0ef48ab817",
		"tableId" : "5884ac2244c9cb0ef48ab818",
		"seatIds" : []string{"1"},
		"tableNo" : "T1",
		"seatNo" : "1",
		"guestCount" : 1,
		"customer" : bson.ObjectIdHex("5927c3744823810f1696a2d2"),
		"isNewCustomer" : true,
		"customerName" : "parvatham k",
		"feedback" : "No feedback",
		"textReceipt" : true,
		"receiptMobileNumber" : "8807712908",
		"emailReceipt" : true,
		"receiptMailId" : "parvatha@arthika.com",
		"service" : "TABLESERVICE",
		"totalTaxAmount" : 861,
		"inclusiveTaxAmount" : 861,
		"taxExempt" : false,
		"taxExemptAmount" : 0,
		"tipPercentage" : 0,
		"tipTotalAmount" : 0,
		"grossVoid" : 0,
		"netVoid" : 0,
		"taxVoid" : 0,
		"appliedSaleDiscounts" : []string{},
		"discounts" : 0,
		"discountTax" : 0,
		"checkDiscountAmount" : 0,
		"checkDiscountTax" : 0,
		"itemDiscountAmount" : 0,
		"itemDiscountTax" : 0,
		"totalTaxAmountWithOutDiscounts" : 861,
		"totalInclusiveDiscountTax" : 0,
		"totalInclusiveCheckDiscountTax" : 0,
		"totalVoidInclusiveTax" : 0,
		"voidDiscounts" : 0,
		"voidDiscountTax" : 0,
		"deliveryCharge" : 0,
		"giftCardAmount" : 0,
		"cashAmount" : 9110,
		"sideCC" : 0,
		"creditCardAmount" : 0,
		"giftCardSoldAmount" : 0,
		"castleGoAmount" : 0,
		"othersPayment" : 0,
		"loyaltyAmount" : 0,
		"paidAmount" : 9110,
		"changeDue" : 0,
		"refunds" : []string{},
		"saleOpenDate" :time.Now(),
		"isVoidCheck" : false,
		"forceClose" : false,
		"syncWithCashier" : true,
		"closeDay" : true,
		"syncWithInventory" : true,
		"lastUpdated" : time.Now(),
		"cashierOutDate" : time.Now(),
		"orders" : ordres,
		"taxes" : tax1,
		"appliedDiscounts" : applieddiscount1,
		"appliedCheckDiscounts" : appliedcheckdiscount,
		"payments" : payment,
		"freeItems" : freeitem,
		"giftCards" : giftcard,
		"saleDate" : time.Now(),
		"saleTime" : "19:22",
		"saleHour" : 19,
		"saleMinute" : 22,
		"saleOpenTime" : "19:16",
		"saleOpenHour" : 19,
		"saleOpenMinute" : 0,
		"createdBy" : bson.ObjectIdHex("56e1c14208811a14c835b73c"),
		"dateCreated" : time.Now(),
		"startDate" : time.Now().Add(-2*time.Hour),
		"closeDate" : time.Now(),
	}


	sales:= make([]map[string]interface{},0)


	//store := bson.ObjectIdHex("56e1c14208811a14c835b73c")
	//employee := bson.ObjectIdHex("56e1c14208811a14c835b73c")
	//createdBy := bson.ObjectIdHex("56e1c14208811a14c835b73c")
	customer := bson.ObjectIdHex("5927c3744823810f1696a2d2")
	isNewCustomer := true
	customerName := "parvatham k"


	for i := 1; i <= 1; i++ {
		//sale["store"] = store
		//sale["employee"] = employee
		//sale["createdBy"] = createdBy
		sale["saleNo"] =i+1
		sale["grossSales"]=100
		sale["grandSales"]=200
		sale["grossReceipt"]= 1000
		sale["totalTaxAmount"]=25
		sale["inclusiveTaxAmount"]=50
		sale["serviceCharge"]=10
		sale["netSales"] =5000
		sale["tipPercentage"]=10
		sale["tipTotalAmount"]=500
		sale["grossVoid"]=3000
		sale["netVoid"]=2000
		sale["taxVoid"]=4000
		sale["paidAmount"]=5000
		sale["balanceAmount"]=300
		sale["customer"] = customer
		sale["isNewCustomer"] = isNewCustomer
		sale["customerName"] = customerName
		sale["saleUniqueId"] = rand.Int()

		//for _, order := range sale["orders"].([]map[string]interface{}) {
		//	modifier := order["modifiers"].([]map[string]interface{})
		//	log.Println(modifier)
		//	tax := order["taxes"].([]map[string]interface{})
		//	log.Println(tax)
		//	alltax := order["allTaxes"].([]map[string]interface{})
		//	log.Println(alltax)
		//	applieddiscount := order["appliedDiscounts"].([]map[string]interface{})
		//	log.Println(applieddiscount)
		//
		//	modifier1 := order["modifiers"].([]map[string]interface{})
		//	log.Println(modifier1)
		//	tax1 := order["taxes"].([]map[string]interface{})
		//	log.Println(tax1)
		//	alltax1 := order["allTaxes"].([]map[string]interface{})
		//	log.Println(alltax1)
		//	applieddiscount1 := order["appliedDiscounts"].([]map[string]interface{})
		//	log.Println(applieddiscount1)
		//
		//	modifier2 := order["modifiers"].([]map[string]interface{})
		//	log.Println(modifier2)
		//	tax2 := order["taxes"].([]map[string]interface{})
		//	log.Println(tax2)
		//	alltax2 := order["allTaxes"].([]map[string]interface{})
		//	log.Println(alltax2)
		//	applieddiscount2 := order["appliedDiscounts"].([]map[string]interface{})
		//	log.Println(applieddiscount2)
		//
		//	modifier3 := order["modifiers"].([]map[string]interface{})
		//	log.Println(modifier3)
		//	tax3 := order["taxes"].([]map[string]interface{})
		//	log.Println(tax3)
		//	alltax3 := order["allTaxes"].([]map[string]interface{})
		//	log.Println(alltax3)
		//	applieddiscount3 := order["appliedDiscounts"].([]map[string]interface{})
		//	log.Println(applieddiscount3)
		//
		//	modifier4 := order["modifiers"].([]map[string]interface{})
		//	log.Println(modifier4)
		//	tax4 := order["taxes"].([]map[string]interface{})
		//	log.Println(tax4)
		//	alltax4 := order["allTaxes"].([]map[string]interface{})
		//	log.Println(alltax4)
		//	applieddiscount4 := order["appliedDiscounts"].([]map[string]interface{})
		//	log.Println(applieddiscount4)
		//
		//	modifier5 := order["modifiers"].([]map[string]interface{})
		//	log.Println(modifier5)
		//	tax5 := order["taxes"].([]map[string]interface{})
		//	log.Println(tax5)
		//	alltax5 := order["allTaxes"].([]map[string]interface{})
		//	log.Println(alltax5)
		//	applieddiscount5 := order["appliedDiscounts"].([]map[string]interface{})
		//	log.Println(applieddiscount5)
		//
		//	modifier6 := order["modifiers"].([]map[string]interface{})
		//	log.Println(modifier6)
		//	tax6 := order["taxes"].([]map[string]interface{})
		//	log.Println(tax6)
		//	alltax6 := order["allTaxes"].([]map[string]interface{})
		//	log.Println(alltax6)
		//	applieddiscount6 := order["appliedDiscounts"].([]map[string]interface{})
		//	log.Println(applieddiscount6)
		//
		//	modifier7 := order["modifiers"].([]map[string]interface{})
		//	log.Println(modifier7)
		//	tax7 := order["taxes"].([]map[string]interface{})
		//	log.Println(tax7)
		//	alltax7 := order["allTaxes"].([]map[string]interface{})
		//	log.Println(alltax7)
		//	applieddiscount7 := order["appliedDiscounts"].([]map[string]interface{})
		//	log.Println(applieddiscount7)
		//}
		//tax8 := sale["taxes"].([]map[string]interface{})
		//log.Println(tax8)
		//
		//applieddiscount8 := sale["appliedDiscounts"].([]map[string]interface{})
		//log.Println(applieddiscount8)
		//
		//appliedcheckdiscount := sale["appliedCheckDiscounts"].([]map[string]interface{})
		//log.Println(appliedcheckdiscount)
		//
		//payment := sale["payments"].([]map[string]interface{})
		//log.Println(payment)
		//
		//freeitem := sale["freeItems"].([]map[string]interface{})
		//log.Println(freeitem)
		//
		//giftcard := sale["giftCards"].([]map[string]interface{})
		//log.Println(giftcard)
		//
		//refund := sale["refunds"].([]map[string]interface{})
		//log.Println(refund)
		sales  = append(sales,sale)
	}

	db := connectDB()
	for _,sale := range sales {
		if err := db.C("sale").Insert(sale); err != nil {
			log.Print("Error - ", err.Error())
		}

	}

	log.Println("complete")


}

func connectDB() *mgo.Database {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB("Linga")
}
