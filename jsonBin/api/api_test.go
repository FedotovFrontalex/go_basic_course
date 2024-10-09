package api_test

import (
	"encoding/json"
	"errors"
	"jsonBin/api"
	"jsonBin/bins"
	"jsonBin/config"
	"jsonBin/file"
	"jsonBin/testHelpers"
	"net/http"
	"os"
	"testing"
	"time"
)

const KEY = "$2a$10$be9Pc0WryjC6kd0hMxPI4"
const ACCESS_KEY = "access key"


func TestInit(t *testing.T) {

	expected := &api.Api{
		Key:       KEY,
		AccessKey: ACCESS_KEY,
	}

	conf := &config.Config{
		Key:       KEY,
		AccessKey: ACCESS_KEY,
	}

	got := api.Init(conf)

	if got.Key != expected.Key || got.AccessKey != expected.AccessKey {
		t.Errorf("Получено %v, ожидалось %v", got, expected)
	}
}

func TestCreate(t *testing.T) {
	t.Run("testSuccess", func(t *testing.T) {
		filename, err := testHelpers.CreateMockFile("json")
		defer os.Remove(filename)
		binName := "binName"

		if err != nil {
			t.Error("Can't create mock file")
		}

		successResponse := &api.CreateApiResponse{
			Metadata: bins.Bin{
				Id:        "1",
				Name:      binName,
				Private:   true,
				CreatedAt: time.Now(),
			},
		}

		expected := successResponse.Metadata

		respBytes, _ := json.Marshal(successResponse)

		fakeTransport := testHelpers.CreateFakeTransport(respBytes, nil, 200)

		oldTransport := http.DefaultTransport
		
		defer func() {
			http.DefaultTransport = oldTransport
		}()
		
		http.DefaultTransport = fakeTransport

		binApi := &api.Api{
			Key:       KEY,
			AccessKey: ACCESS_KEY,
		}

		data, err := binApi.Create(filename, binName)

		if data.Id != expected.Id || data.Private != expected.Private || data.CreatedAt.UTC() != expected.CreatedAt.UTC() || data.Name != expected.Name {
			t.Errorf("Ожидалось: %v, Получено %v", expected, data)
		}
	})

	t.Run("file is not valid json", func(t *testing.T) {
		filename, err := testHelpers.CreateMockFile("txt")
		defer os.Remove(filename)
		binName := "binName"

		if err != nil {
			t.Error("Can't create mock file")
		}

		binApi := &api.Api{
			Key:       KEY,
			AccessKey: ACCESS_KEY,
		}

		data, err := binApi.Create(filename, binName)

		if data != nil {
			t.Errorf("Ожидалось: %v, Получено %v", nil, data)
		}

		if err != file.ErrNotJson {
			t.Errorf("Ожидалось: %v, Получено %v", file.ErrNotJson, err)
		}
	})

	t.Run("error http request", func(t *testing.T) {
		filename, err := testHelpers.CreateMockFile("json")
		defer os.Remove(filename)
		binName := "binName"

		if err != nil {
			t.Error("Can't create mock file")
		}
		
		error := api.ErrHttpRequest
		fakeTransport := testHelpers.CreateFakeTransport(nil, error, 500)

		oldTransport := http.DefaultTransport
		
		defer func() {
			http.DefaultTransport = oldTransport
		}()
		
		http.DefaultTransport = fakeTransport

		binApi := &api.Api{
			Key:       KEY,
			AccessKey: ACCESS_KEY,
		}

		data, err := binApi.Create(filename, binName)

		if data != nil {
			t.Errorf("Ожидалось: %v, Получено %v", nil, data)
		}

		if err != error {
			t.Errorf("Ожидалось: %v, Получено %v", error, err)
		}
	})
	
	t.Run("response with no 200", func(t *testing.T) {
		filename, err := testHelpers.CreateMockFile("json")
		defer os.Remove(filename)
		binName := "binName"

		if err != nil {
			t.Error("Can't create mock file")
		}

		expected := errors.New("some message from server")

		response := &api.ErrorBinResponse{
				Message: expected.Error(),
		}

		respBytes, _ := json.Marshal(response)

		fakeTransport := testHelpers.CreateFakeTransport(respBytes, nil, 400)

		oldTransport := http.DefaultTransport
	
		defer func() {
			http.DefaultTransport = oldTransport
		}()
	
		http.DefaultTransport = fakeTransport

		binApi := &api.Api{
			Key:       KEY,
			AccessKey: ACCESS_KEY,
		}

		_, err = binApi.Create(filename, binName)

		 if err.Error() != expected.Error() {
			t.Errorf("Ожидалось: %v, Получено %v", expected, err)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		filename, err := testHelpers.CreateMockFile("json")
		defer os.Remove(filename)
		id := "1"

		if err != nil {
			t.Error("Can't create mock file")
		}

		successResponse := "OK"

		fakeTransport := testHelpers.CreateFakeTransport([]byte(successResponse), nil, 200)

		oldTransport := http.DefaultTransport
		
		defer func() {
			http.DefaultTransport = oldTransport
		}()
		
		http.DefaultTransport = fakeTransport

		binApi := &api.Api{
			Key:       KEY,
			AccessKey: ACCESS_KEY,
		}

		err = binApi.Update(filename, id)

		if err != nil {
			t.Errorf("Ожидалось: %v, Получено %v", nil, err)
		}
	})

	t.Run("file is not valid json", func(t *testing.T) {
		filename, err := testHelpers.CreateMockFile("txt")
		defer os.Remove(filename)
		id := "1"

		if err != nil {
			t.Error("Can't create mock file")
		}

		binApi := &api.Api{
			Key:       KEY,
			AccessKey: ACCESS_KEY,
		}

		err = binApi.Update(filename, id)

		if err != file.ErrNotJson {
			t.Errorf("Ожидалось: %v, Получено %v", file.ErrNotJson, err)
		}
	})

	t.Run("error http request", func(t *testing.T) {
		filename, err := testHelpers.CreateMockFile("json")
		defer os.Remove(filename)
		id := "1"

		if err != nil {
			t.Error("Can't create mock file")
		}
		
		error := api.ErrHttpRequest
		fakeTransport := testHelpers.CreateFakeTransport(nil, error, 500)

		oldTransport := http.DefaultTransport
		
		defer func() {
			http.DefaultTransport = oldTransport
		}()
		
		http.DefaultTransport = fakeTransport

		binApi := &api.Api{
			Key:       KEY,
			AccessKey: ACCESS_KEY,
		}

		err = binApi.Update(filename, id)

		if err != error {
			t.Errorf("Ожидалось: %v, Получено %v", error, err)
		}
	})
	
	t.Run("response with no 200", func(t *testing.T) {
		filename, err := testHelpers.CreateMockFile("json")
		defer os.Remove(filename)
		id := "binName"

		if err != nil {
			t.Error("Can't create mock file")
		}

		expected := errors.New("some message from server")

		response := &api.ErrorBinResponse{
				Message: expected.Error(),
		}

		respBytes, _ := json.Marshal(response)

		fakeTransport := testHelpers.CreateFakeTransport(respBytes, nil, 400)

		oldTransport := http.DefaultTransport
	
		defer func() {
			http.DefaultTransport = oldTransport
		}()
	
		http.DefaultTransport = fakeTransport

		binApi := &api.Api{
			Key:       KEY,
			AccessKey: ACCESS_KEY,
		}

		err = binApi.Update(filename, id)

		 if err.Error() != expected.Error() {
			t.Errorf("Ожидалось: %v, Получено %v", expected, err)
		}
	})
}
