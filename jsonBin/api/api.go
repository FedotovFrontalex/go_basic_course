package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"jsonBin/bins"
	"jsonBin/config"
	"jsonBin/file"
	"net/http"
)

type Api struct {
	Key       string
	AccessKey string
}

type CreateApiResponse struct {
	Metadata bins.Bin `json:"metadata"`
}

type ErrorBinResponse struct {
	Message string `json:"message"`
}

var ErrHttpRequest = errors.New("error request")

func Init(config *config.Config) *Api {
	return &Api{
		Key:       config.Key,
		AccessKey: config.AccessKey,
	}
}

func (api *Api) Create(filename string, binName string) (*bins.Bin, error) {
	data, err := file.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	resp, _ := http.NewRequest(
		http.MethodPost,
		"https://api.jsonbin.io/v3/b",
		bytes.NewBuffer(data),
	)

	resp.Header.Add("Content-Type", "application/json")
	resp.Header.Add("X-Master-Key", api.Key)
	resp.Header.Add("X-Access-Key", api.AccessKey)
	resp.Header.Add("X-Bin-Name", binName)

	response, err := http.DefaultClient.Do(resp)

	if err != nil {
		return nil, ErrHttpRequest
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		var responseData CreateApiResponse
		json.Unmarshal(body, &responseData)

		return &responseData.Metadata, nil
	}

	var errorResponse ErrorBinResponse
	json.Unmarshal(body, &errorResponse)

	return nil, errors.New(errorResponse.Message)

}

func (api *Api) Update(filename string, id string) error {
	data, err := file.ReadFile(filename)

	if err != nil {
		return err
	}

	resp, err := http.NewRequest(
		http.MethodPut,
		"https://api.jsonbin.io/v3/b/"+id,
		bytes.NewBuffer(data),
	)

	resp.Header.Add("Content-Type", "application/json")
	resp.Header.Add("X-Master-Key", api.Key)
	resp.Header.Add("X-Access-Key", api.AccessKey)

	response, err := http.DefaultClient.Do(resp)

	if err != nil {
		return ErrHttpRequest
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	if response.StatusCode == 200 {
		return nil
	}

	var errorResponse ErrorBinResponse
	json.Unmarshal(body, &errorResponse)

	return errors.New(errorResponse.Message)
}

func (api *Api) Get(id string) (string, error) {
	resp, err := http.NewRequest(
		http.MethodGet,
		"https://api.jsonbin.io/v3/b/"+id,
		nil,
	)

	resp.Header.Add("Content-Type", "application/json")
	resp.Header.Add("X-Master-Key", api.Key)
	resp.Header.Add("X-Access-Key", api.AccessKey)
	resp.Header.Add("X-Bin-Meta", "false")

	response, err := http.DefaultClient.Do(resp)

	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	if response.StatusCode == 200 {
		return string(body), nil
	}

	var errorResponse ErrorBinResponse
	json.Unmarshal(body, &errorResponse)

	return "", errors.New(errorResponse.Message)
}

func (api *Api) Delete(id string) error {
	resp, err := http.NewRequest(
		http.MethodDelete,
		"https://api.jsonbin.io/v3/b/"+id,
		nil,
	)

	resp.Header.Add("Content-Type", "application/json")
	resp.Header.Add("X-Master-Key", api.Key)
	resp.Header.Add("X-Access-Key", api.AccessKey)

	response, err := http.DefaultClient.Do(resp)

	if err != nil {
		return err
	}

	if response.StatusCode == 200 {
		return nil
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	var errorResponse ErrorBinResponse
	json.Unmarshal(body, &errorResponse)

	return errors.New(errorResponse.Message)
}
