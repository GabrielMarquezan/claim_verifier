package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func StartClaimAnalysis(claim []byte, address string) error {
	response, err := http.Post(address+"/claims", "application/json", bytes.NewBuffer(claim))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Println("Status code:", response.Status)
	fmt.Println("Response:", string(body))

	return nil
}
