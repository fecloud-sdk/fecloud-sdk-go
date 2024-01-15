package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteListenerTagsResponse", string(data)}, " ")
}
