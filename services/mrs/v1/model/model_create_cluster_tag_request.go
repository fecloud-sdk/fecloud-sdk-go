package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateClusterTagRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *CreateTagReq `json:"body,omitempty"`
}

func (o CreateClusterTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterTagRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterTagRequest", string(data)}, " ")
}
