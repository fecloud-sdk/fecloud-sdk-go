package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowClusterDetailsResponse struct {
	Cluster        *Cluster `json:"cluster,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowClusterDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailsResponse", string(data)}, " ")
}
