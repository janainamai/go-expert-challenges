package shared

import (
	"encoding/json"
	"fmt"

	"github.com/janainamai/go-expert-challenges/internal/shared/dto"
)

func Encode(value interface{}) ([]byte, *dto.Error) {
	jsonResult, err := json.Marshal(value)
	if err != nil {
		message := fmt.Sprintf("Error during encoding JSON: %v", err.Error())
		return nil, dto.InitError().WithDetail(message)
	}

	return jsonResult, nil
}
