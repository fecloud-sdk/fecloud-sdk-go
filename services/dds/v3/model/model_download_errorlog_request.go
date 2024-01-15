package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DownloadErrorlogRequest struct {
	InstanceId string `json:"instance_id"`

	Body *DownloadErrorlogRequestBody `json:"body,omitempty"`
}

func (o DownloadErrorlogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadErrorlogRequest struct{}"
	}

	return strings.Join([]string{"DownloadErrorlogRequest", string(data)}, " ")
}
