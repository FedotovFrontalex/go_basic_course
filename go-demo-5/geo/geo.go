package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CheckCityResponse struct {
	Error bool `json:"error"`
}

func GetLocation(location string) (*GeoData, error) {
	if location != "" {
		isCity, err := checkCity(location)
		if err != nil {
			return nil, err
		}

		if !isCity {
			return nil, errors.New("city is not found")
		}

		return &GeoData{
			City: location,
		}, nil
	}
	res, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("request error")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var geo GeoData

	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, errors.New("Can't convert data response")
	}

	return &geo, nil
}

func checkCity(city string) (bool, error) {
	data, _ := json.Marshal(map[string]string{
		"city": city,
	})

	res, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	var checkResp CheckCityResponse
	json.Unmarshal(body, &checkResp)

	return !checkResp.Error, nil
}
