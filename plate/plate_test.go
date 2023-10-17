package plate_test

import (
	"testing"

	plate "github.com/kisukegremory/plateapi/plate"
)

func TestPlateValidate(t *testing.T) {
	match, _ := plate.PlateValidate("AAA0A00")
	if !match {
		t.Error("Should Match")
	}

	match, _ = plate.PlateValidate("AA00A00")
	if match {
		t.Error("Should not match")
	}

	match, _ = plate.PlateValidate("000A00")
	if match {
		t.Error("Should not match")
	}
}
