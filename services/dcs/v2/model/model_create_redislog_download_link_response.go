package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateRedislogDownloadLinkResponse struct {
	Id *string `json:"id,omitempty"`

	Link           *string `json:"link,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRedislogDownloadLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedislogDownloadLinkResponse struct{}"
	}

	return strings.Join([]string{"CreateRedislogDownloadLinkResponse", string(data)}, " ")
}
