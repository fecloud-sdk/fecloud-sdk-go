package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteInstancesResponse struct {
	Results        *[]BatchOpsResult `json:"results,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o BatchDeleteInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstancesResponse", string(data)}, " ")
}
