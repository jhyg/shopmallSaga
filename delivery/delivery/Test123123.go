package delivery

import (
	"time"
)

type Test123123 struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	
}

func NewTest123123() *Test123123{
	event := &Test123123{EventType:"Test123123", TimeStamp:time.Now().String()}

	return event
}
