package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateRedislogDownloadLinkResponse Response Object
type CreateRedislogDownloadLinkResponse struct {

	// 日志id
	Id *string `json:"id,omitempty"`

	// 日志下载链接，默认有效时间为24小时
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
