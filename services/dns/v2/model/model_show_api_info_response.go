package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowApiInfoResponse struct {
	Version        *VersionItem `json:"version,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowApiInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowApiInfoResponse", string(data)}, " ")
}
