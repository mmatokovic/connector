package cdm

type City struct {
	CityZipCode string `json:"city_zip_code"`
	IdCounty    int    `json:"id_county"`
	Name        string `json:"name"`
}

type County struct {
	IdCounty int    `json:"id_county"`
	Name     string `json:"name"`
}
