package apparel

import "fmt"

type APIError struct {
	ErrorCode   string `xml:"ErrorCode"`
	Description string `xml:"Description"`
}

func (e APIError) Error() string {
	return fmt.Sprintf(`{"code":"%s", "description": "%s"}`, e.ErrorCode, e.Description)
}
