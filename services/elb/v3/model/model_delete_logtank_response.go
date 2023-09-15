package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteLogtankResponse Response Object
type DeleteLogtankResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogtankResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogtankResponse", string(data)}, " ")
}
