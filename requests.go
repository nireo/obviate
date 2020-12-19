package obviate

import (
	"io"
	"net/http"
)

func Get(url string) (*http.Response, error) {
	response, err := http.Get("http://www.example.com")
	if err != nil {
		return nil, err
	}

	return response, nil
}

func Delete(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return request.Response, nil
}

func Patch(url string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}

	return request.Response, nil
}

func Post(url string, body io.Reader, contentType string) (*http.Response, error) {
	response, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func Put(url string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	return request.Response, nil
}
