package delivery

import (
	"github.com/mitchellh/mapstructure"
)

func wheneverOrderPlaced_StartDelivery(data map[string]interface{}){
	
	event := NewOrderPlaced()
	mapstructure.Decode(data,&event)
	delivery := &Delivery{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	deliveryrepository.save(delivery)
	
}

func wheneverTest123123_CancelDelivery(data map[string]interface{}){
	
	event := NewTest123123()
	mapstructure.Decode(data,&event)
	delivery := &Delivery{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	deliveryrepository.save(delivery)
	
}

