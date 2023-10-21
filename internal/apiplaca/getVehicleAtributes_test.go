package apiplaca

import (
	"reflect"
	"testing"
)

func TestGetVehiclesAtributes(t *testing.T) {
	vehicle, err := GetVehicleAttributesByPlate("INT8C36")
	vehicleExpected := VehicleAttributesAPI{
		Plate:        "INT8C36",
		Year:         2007,
		ModelYear:    2007,
		Manufacturer: "VW",
		VehicleModel: "CROSSFOX",
		SubModel:     "CROSSFOX",
		Version:      "CROSSFOX",
		Color:        "Prata",
		Uf:           "RS",
		City:         "São Leopoldo",
		Origin:       "NACIONAL",
	}

	if err != nil {
		t.Error("Error on getting the atributes of the plate")
	}
	if !reflect.DeepEqual(vehicle, vehicleExpected) {
		t.Error("Vehicle Different from th expected")
	}
}
