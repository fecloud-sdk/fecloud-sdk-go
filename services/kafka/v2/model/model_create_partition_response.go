package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePartitionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePartitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePartitionResponse struct{}"
	}

	return strings.Join([]string{"CreatePartitionResponse", string(data)}, " ")
}
