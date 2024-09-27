package shared

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
)

func Decode(r io.ReadCloser, request interface{}) *dto.Error {
	err := json.NewDecoder(r).Decode(request)

	if err != nil {
		message := fmt.Sprintf("Error during decoding JSON: %v", err.Error())
		return dto.InitError().WithDetail(message)
	}

	return nil
}
