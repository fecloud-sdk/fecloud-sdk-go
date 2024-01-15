package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchValidateClustersConnectionsResponse struct {
	Results *[]CheckJobResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchValidateClustersConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchValidateClustersConnectionsResponse struct{}"
	}

	return strings.Join([]string{"BatchValidateClustersConnectionsResponse", string(data)}, " ")
}
