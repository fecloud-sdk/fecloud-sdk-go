package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPasswordComplexityResponse struct {
	TotalNum *int64 `json:"total_num,omitempty"`

	DataList       *[]PwdPolicyInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListPasswordComplexityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPasswordComplexityResponse struct{}"
	}

	return strings.Join([]string{"ListPasswordComplexityResponse", string(data)}, " ")
}
