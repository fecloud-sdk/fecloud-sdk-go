package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestartOrFlushInstancesResponse struct {
	Results        *[]BatchOpsResult `json:"results,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o RestartOrFlushInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartOrFlushInstancesResponse struct{}"
	}

	return strings.Join([]string{"RestartOrFlushInstancesResponse", string(data)}, " ")
}
