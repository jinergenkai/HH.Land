package models

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Polygon struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type LandParcelRaw struct {
	Status     int     `json:"status"`
	Points     []Point `json:"points"`
	Address    string  `json:"address"`
	Id         int     `json:"id"`
	ProvinceId string  `json:"province_id"`
	DistrictId string  `json:"district_id"`
	WardId     string  `json:"ward_id"`
	Html       string  `json:"html"`
	Ext        []any   `json:"ext"`
}

type LandParcel struct {
	Status     int         `json:"status"`
	Location   interface{} `json:"location"`
	Address    string      `json:"address"`
	Id         int         `json:"id"`
	ProvinceId string      `json:"province_id"`
	DistrictId string      `json:"district_id"`
	WardId     string      `json:"ward_id"`
	Html       string      `json:"html"`
	Ext        []any       `json:"ext"`
}

// ConvertToGeoJSON chuyển đổi LandParcelRaw thành LandParcel đúng chuẩn GeoJSON
func ConvertToGeoJSON(raw LandParcelRaw) LandParcel {
	var coordinates [][]float64

	// Chuyển đổi mảng points thành GeoJSON coordinates
	for _, p := range raw.Points {
		coordinates = append(coordinates, []float64{p.Lng, p.Lat})
	}

	// Tạo đối tượng GeoJSON
	geoJSON := Polygon{
		Type:        "Polygon",
		Coordinates: [][][]float64{coordinates}, // Đưa vào mảng cấp 3
	}

	return LandParcel{
		Status:     raw.Status,
		Location:   geoJSON,
		Address:    raw.Address,
		Id:         raw.Id,
		ProvinceId: raw.ProvinceId,
		DistrictId: raw.DistrictId,
		WardId:     raw.WardId,
		Html:       raw.Html,
		Ext:        raw.Ext,
	}
}
