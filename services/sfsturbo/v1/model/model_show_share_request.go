package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowShareRequest struct {
	ContentType string `json:"Content-Type"`

	ShareId string `json:"share_id"`
}

func (o ShowShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShareRequest struct{}"
	}

	return strings.Join([]string{"ShowShareRequest", string(data)}, " ")
}
