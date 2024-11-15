package req

import (
	"bytes"
	"io"
	"net/http"
)

func POST(url string, str string) (string, error) {
	data := []byte(str)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
