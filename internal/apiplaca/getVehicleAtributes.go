package apiplaca

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

var apiToken = os.Getenv("APIPLACA_TOKEN")

func GetVehicleAttributesByPlate(plateString string) (vehicle VehicleAttributesAPI, err error) {
	url := fmt.Sprintf("https://wdapi2.com.br/consulta/%s/%s", plateString, apiToken)
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		slog.Error("Problems on request: ", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Error on reading body: ", err)
		return
	}
	err = json.Unmarshal(body, &vehicle)
	if err != nil {
		slog.Error("Error on Parsing body: ", err)
		return
	}
	return
}
