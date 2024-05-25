package types

type Distance struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obuID"`
	Unix  int64   `json:"unix"`
}

type OBUData struct {
	OBUID int     `json:"OBUID"`
	Lat   float64 `json:"Lat"`
	Long  float64 `json:"Long"`
}
