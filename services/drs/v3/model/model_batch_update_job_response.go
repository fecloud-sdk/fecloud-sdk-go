package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateJobResponse struct {
	Count *int32 `json:"count,omitempty"`

	Results        *[]ModifyJobResp `json:"results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchUpdateJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateJobResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateJobResponse", string(data)}, " ")
}
