package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowApiInfoRequest struct {
	Version string `json:"version"`
}

func (o ShowApiInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowApiInfoRequest", string(data)}, " ")
}
