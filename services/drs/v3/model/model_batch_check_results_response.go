package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCheckResultsResponse struct {
	Results *[]QueryPreCheckResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCheckResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckResultsResponse struct{}"
	}

	return strings.Join([]string{"BatchCheckResultsResponse", string(data)}, " ")
}
