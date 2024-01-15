package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateClusterTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateClusterTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterTagResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterTagResponse", string(data)}, " ")
}
