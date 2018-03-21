package main

import (
	//"zenpepper.com/zenpepper/server/controllers"
	"encoding/json"
	"log"
	//"zenpepper.com/zenpepper/server/services"
	"zenpepper.com/zenpepper/server/models"
)

func main() {

	//orders := controllers.OrderHistory{}
	//orders := services.CustomerLoyalty{}
	//orders := models.Customer{}
	orders := models.WebMultiStoreData{}
	deliveryZone := models.DeliveryZone{}
	deliveryZone.ZoneDetails = append(deliveryZone.ZoneDetails,models.ZoneDetail{})
	orders.DeliveryZones = append(orders.DeliveryZones,deliveryZone)



	//orders.Address = append(orders.Address,models.Address{})

	res, err:= json.Marshal(&orders)
    if err != nil {
		log.Print(err)
	}

	log.Printf("Response: %s\n", res)
}
