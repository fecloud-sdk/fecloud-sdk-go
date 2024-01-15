package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DownloadSlowlogResult struct {
	NodeName string `json:"node_name"`

	FileName string `json:"file_name"`

	Status string `json:"status"`

	FileSize string `json:"file_size"`

	FileLink string `json:"file_link"`

	UpdateAt int64 `json:"update_at"`
}

func (o DownloadSlowlogResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSlowlogResult struct{}"
	}

	return strings.Join([]string{"DownloadSlowlogResult", string(data)}, " ")
}
