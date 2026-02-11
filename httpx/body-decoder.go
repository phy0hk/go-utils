package httpx

import (
	"encoding/json"
	"fmt"
)

func BodyDecoder[T any](body *[]byte, json_writer *T) error {
	if len(*body) == 0 {
		return fmt.Errorf("no request body")
	}
	if err := json.Unmarshal(*body, json_writer); err != nil {
		return fmt.Errorf("failed to decode json : %s", err.Error())
	}
	return nil
}
