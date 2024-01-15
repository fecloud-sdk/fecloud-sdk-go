package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchListRposAndRtosResponse struct {
	Count *int32 `json:"count,omitempty"`

	Results        *[]QueryRpoAndRtoResp `json:"results,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchListRposAndRtosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListRposAndRtosResponse struct{}"
	}

	return strings.Join([]string{"BatchListRposAndRtosResponse", string(data)}, " ")
}
