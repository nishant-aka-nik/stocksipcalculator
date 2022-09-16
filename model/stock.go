package model

type Stock struct {
	Name string  `json:"name" bson:"name"`
	LTP  float32 `json:"ltp" bson:"ltp"`
}


