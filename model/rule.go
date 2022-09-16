package model

type Rule struct {
	StockRules []StockRule `json:"stockrule" bson:"stockrule"`
	Email      string      `json:"email" bson:"email"`
}

type StockRule struct {
	ID               string  `bson:"id"`
	StockName        string  `json:"stockname" bson:"stockname"`
	TargetPercentage float32 `json:"targetpercentage" bson:"targetpercentage"`
	TargetPrice      float32 `json:"targetprice" bson:"targetprice"`
	Alerted          bool    `bson:"alerted"`
}
