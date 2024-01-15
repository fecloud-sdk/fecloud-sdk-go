package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GetBackupDownloadLinkResponseBodyFiles struct {
	Name string `json:"name"`

	Size int64 `json:"size"`

	DownloadLink string `json:"download_link"`

	LinkExpiredTime string `json:"link_expired_time"`
}

func (o GetBackupDownloadLinkResponseBodyFiles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetBackupDownloadLinkResponseBodyFiles struct{}"
	}

	return strings.Join([]string{"GetBackupDownloadLinkResponseBodyFiles", string(data)}, " ")
}
