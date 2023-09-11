package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchCreateDeletePrivateNatTagsResponse Response Object
type BatchCreateDeletePrivateNatTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateDeletePrivateNatTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeletePrivateNatTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDeletePrivateNatTagsResponse", string(data)}, " ")
}
