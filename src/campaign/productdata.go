package main

type ProductLocation struct {
	District string `form:"district,omitempty" json:"district,omitempty" xml:"district,omitempty"`
	Province string `form:"province,omitempty" json:"province,omitempty" xml:"province,omitempty"`
}
type ProductCriteria struct {
	ComparisonType string `form:"comparisonType,omitempty" json:"comparisonType,omitempty" xml:"comparisonType,omitempty"`
	MaxValue       int    `form:"maxValue,omitempty" json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue       int    `form:"minValue,omitempty" json:"minValue,omitempty" xml:"minValue,omitempty"`
	Variable       string `form:"variable,omitempty" json:"variable,omitempty" xml:"variable,omitempty"`
}
type ProductPayload struct {
	AvailableLocations []ProductLocation `form:"availableLocations,omitempty" json:"availableLocations,omitempty" xml:"availableLocations,omitempty"`
	ClientCode         string            `form:"clientCode,omitempty" json:"clientCode,omitempty" xml:"clientCode,omitempty"`
	Criteria           []ProductCriteria `form:"criteria,omitempty" json:"criteria,omitempty" xml:"criteria,omitempty"`
	DailyVolume        int               `form:"dailyVolume,omitempty" json:"dailyVolume,omitempty" xml:"dailyVolume,omitempty"`
	ProductCode        string            `form:"productCode,omitempty" json:"productCode,omitempty" xml:"productCode,omitempty"`
	ProductType        string            `form:"productType,omitempty" json:"productType,omitempty" xml:"productType,omitempty"`
}

// productLocation user type.

var productMap = make(map[string]ProductPayload)

func(p *ProductPayload) init() {
	productMap["1"] = ProductPayload{ProductCode: "CLWDB1", ProductType: "cash_loan", ClientCode: "home_credit", DailyVolume: 1000, AvailableLocations: []ProductLocation{{"q1", "hcmc"}, {"q6", "hanoi"}}, Criteria: []ProductCriteria{{Variable: "credit_score", ComparisonType: "between", MinValue: 400, MaxValue: 500}}}
	productMap["2"] = ProductPayload{ProductCode: "CLWDB2", ProductType: "insurance", ClientCode: "insure", DailyVolume: 800, AvailableLocations: []ProductLocation{{"q6", "hanoi"}, {"q7", "tan sun"}}, Criteria: []ProductCriteria{{Variable: "credit_score", ComparisonType: "between", MinValue: 550, MaxValue: 700}}}

}