package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DownloadErrorlogResponse struct {
	List *[]DownloadSlowlogResult `json:"list,omitempty"`

	Status *string `json:"status,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DownloadErrorlogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadErrorlogResponse struct{}"
	}

	return strings.Join([]string{"DownloadErrorlogResponse", string(data)}, " ")
}
