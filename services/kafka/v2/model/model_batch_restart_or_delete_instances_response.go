package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchRestartOrDeleteInstancesResponse struct {
	Results        *[]BatchRestartOrDeleteInstanceRespResults `json:"results,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o BatchRestartOrDeleteInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestartOrDeleteInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchRestartOrDeleteInstancesResponse", string(data)}, " ")
}
