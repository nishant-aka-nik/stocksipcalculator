package model

type Rule struct {
	StockRules []StockRule `json:"stockrule" bson:"stockrule"`
	Email      string      `json:"email" bson:"email"`
}

type StockRule struct {
	StockName        string  `json:"stockname" bson:"stockname"`
	TargetPercentage float64 `json:"targetpercentage" bson:"targetpercentage"`
	TargetPrice      float64 `json:"targetprice" bson:"targetprice"`
}

type DBRule struct {
	ID          string  `bson:"id"`
	StockName   string  `json:"stockname" bson:"stockname"`
	TargetPrice float64 `json:"targetprice" bson:"targetprice"`
	UPBelow     string  `bson:"upbelow"`
	Alerted     bool    `bson:"alerted"`
	Email       string  `json:"email" bson:"email"`
}
