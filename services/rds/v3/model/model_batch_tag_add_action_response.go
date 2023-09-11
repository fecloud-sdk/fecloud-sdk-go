package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchTagAddActionResponse Response Object
type BatchTagAddActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchTagAddActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagAddActionResponse struct{}"
	}

	return strings.Join([]string{"BatchTagAddActionResponse", string(data)}, " ")
}
