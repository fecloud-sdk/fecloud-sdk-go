package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DownloadSlowlogRequest struct {
	InstanceId string `json:"instance_id"`

	Body *DownloadSlowlogRequestBody `json:"body,omitempty"`
}

func (o DownloadSlowlogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSlowlogRequest struct{}"
	}

	return strings.Join([]string{"DownloadSlowlogRequest", string(data)}, " ")
}
