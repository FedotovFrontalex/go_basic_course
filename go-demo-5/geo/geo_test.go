package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestLocation(t *testing.T) {
	city := "London"
	expected := geo.GeoData{
			City: "London",
	}

	got, err := geo.GetLocation(city)
	if err != nil {
		t.Error("Не удалось проверить город")
	}

	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}
