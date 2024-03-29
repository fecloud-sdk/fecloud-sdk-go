package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowVersionRequest struct {
	Version string `json:"version"`
}

func (o ShowVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowVersionRequest", string(data)}, " ")
}
