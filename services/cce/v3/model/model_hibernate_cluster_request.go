package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HibernateClusterRequest struct {
	ClusterId string `json:"cluster_id"`
}

func (o HibernateClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HibernateClusterRequest struct{}"
	}

	return strings.Join([]string{"HibernateClusterRequest", string(data)}, " ")
}
