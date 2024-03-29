package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchTagActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchTagActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionResponse struct{}"
	}

	return strings.Join([]string{"BatchTagActionResponse", string(data)}, " ")
}
