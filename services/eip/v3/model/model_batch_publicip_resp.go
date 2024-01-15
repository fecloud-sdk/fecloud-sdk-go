package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchPublicipResp struct {

	// 响应码
	StatusCode int32 `json:"statusCode"`

	Publicip *PublicipResp `json:"publicip"`
}

func (o BatchPublicipResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPublicipResp struct{}"
	}

	return strings.Join([]string{"BatchPublicipResp", string(data)}, " ")
}
