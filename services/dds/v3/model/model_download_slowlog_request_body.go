package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DownloadSlowlogRequestBody struct {
	FileNameList *[]string `json:"file_name_list,omitempty"`

	NodeIdList *[]string `json:"node_id_list,omitempty"`
}

func (o DownloadSlowlogRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSlowlogRequestBody struct{}"
	}

	return strings.Join([]string{"DownloadSlowlogRequestBody", string(data)}, " ")
}
