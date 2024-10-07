package weather

import (
	"demo/weather/geo"
	"demo/weather/msg"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.GeoData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)

	if err != nil {
		msg.Error(err)
		return ""
	}

	params := url.Values{}
	params.Set("format", fmt.Sprint(format))

	baseUrl.RawQuery = params.Encode()

	res, err := http.Get(baseUrl.String())
	if err != nil {
		msg.Error(err)
		return ""
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		msg.Error(err)
		return ""
	}
	return string(body)
}
