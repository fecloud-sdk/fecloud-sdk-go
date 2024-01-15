package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteSharedTagRequest struct {
	ContentType string `json:"Content-Type"`

	ShareId string `json:"share_id"`

	Key string `json:"key"`
}

func (o DeleteSharedTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSharedTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteSharedTagRequest", string(data)}, " ")
}
