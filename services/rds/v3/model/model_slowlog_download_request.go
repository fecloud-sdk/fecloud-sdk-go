package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SlowlogDownloadRequest struct {
	FileName *string `json:"file_name,omitempty"`
}

func (o SlowlogDownloadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogDownloadRequest struct{}"
	}

	return strings.Join([]string{"SlowlogDownloadRequest", string(data)}, " ")
}
