package main

import (
	"demo/weather/geo"
	"demo/weather/msg"
	"demo/weather/weather"
	"flag"
)

func main() {
	city := flag.String("city", "", "city of user")
	format := flag.Int("format", 1, "output format")
	flag.Parse()

	geo, err := geo.GetLocation(*city)

	if err != nil {
		msg.Error(err)
		return
	}

	weatherResult := weather.GetWeather(*geo, *format)
	msg.Success(weatherResult)
}
