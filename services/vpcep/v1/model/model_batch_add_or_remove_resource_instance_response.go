package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchAddOrRemoveResourceInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddOrRemoveResourceInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrRemoveResourceInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchAddOrRemoveResourceInstanceResponse", string(data)}, " ")
}
