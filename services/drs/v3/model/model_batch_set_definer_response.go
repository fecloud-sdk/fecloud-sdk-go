package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchSetDefinerResponse struct {
	Count *int32 `json:"count,omitempty"`

	Results        *[]ModifyJobResp `json:"results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchSetDefinerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetDefinerResponse struct{}"
	}

	return strings.Join([]string{"BatchSetDefinerResponse", string(data)}, " ")
}
