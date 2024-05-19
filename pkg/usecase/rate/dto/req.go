package user_dto

type CreateRateDto struct {
	ExchangeDate string  `json:"exchangedate"`
	R030         int     `json:"r030"`
	Cc           string  `json:"cc"`
	Txt          string  `json:"txt"`
	Enname       string  `json:"enname"`
	Rate         float64 `json:"rate"`
	Units        int     `json:"units"`
	RatePerUnit  float64 `json:"rate_per_unit"`
	Group        string  `json:"group"`
	Calcdate     string  `json:"calcdate"`
}

type GetRateByValCodeDto struct {
	RateValCode
}

type DeleteRateByValCodeDto struct {
	RateValCode
}

type RateValCode struct {
	ValCode string `json:"valCode"`
}
