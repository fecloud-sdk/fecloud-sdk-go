package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSharedTagsRequest struct {
	ContentType string `json:"Content-Type"`

	ShareId string `json:"share_id"`
}

func (o ShowSharedTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSharedTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowSharedTagsRequest", string(data)}, " ")
}
