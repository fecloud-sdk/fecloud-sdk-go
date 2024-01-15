package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchDeletePublicipTagsResponse Response Object
type BatchDeletePublicipTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeletePublicipTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicipTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicipTagsResponse", string(data)}, " ")
}
