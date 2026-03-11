package types

import "encoding/json"

type MunicipalityCoordinatesFeatureCollection struct {
	Features []MunicipalityCoordinatesFeature `json:"features"`
}

type MunicipalityCoordinatesFeature struct {
	ID         json.RawMessage                `json:"_id"`
	Geometry   MunicipalityCoordinatesGeometry `json:"geometry"`
	Properties MunicipalityCoordinatesProps    `json:"properties"`
}

type MunicipalityCoordinatesProps struct {
	Name       string `json:"name"`
	AreaHa     any    `json:"area_ha"`
	DistrictID any    `json:"district_id"`
}

type MunicipalityCoordinatesGeometry struct {
	Type        string          `json:"type"`
	Coordinates json.RawMessage `json:"coordinates"`
}
