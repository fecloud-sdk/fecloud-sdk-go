package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteShareRequest struct {
	ContentType string `json:"Content-Type"`

	ShareId string `json:"share_id"`
}

func (o DeleteShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareRequest struct{}"
	}

	return strings.Join([]string{"DeleteShareRequest", string(data)}, " ")
}
