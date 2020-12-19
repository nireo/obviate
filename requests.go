package obviate

import "net/http"

func Get(url string) (*http.Response, error) {
	response, err := http.Get("http://www.example.com")
	if err != nil {
		return nil, err
	}

	return response, nil
}
