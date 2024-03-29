package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AwakeClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AwakeClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AwakeClusterResponse struct{}"
	}

	return strings.Join([]string{"AwakeClusterResponse", string(data)}, " ")
}
