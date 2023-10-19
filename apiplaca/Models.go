package apiplaca

type VehicleAttributesAPI struct {
	Plate        string `json:"placa"`
	Year         uint32 `json:"ano,string"`
	ModelYear    uint32 `json:"anoModelo,string"`
	Manufacturer string `json:"MARCA"`
	VehicleModel string `json:"MODELO"`
	SubModel     string `json:"SUBMODELO"`
	Version      string `json:"VERSAO"`
	Color        string `json:"cor"`
	Uf           string `json:"uf"`
	City         string `json:"municipio"`
	Origin       string `json:"Origem"`
}
