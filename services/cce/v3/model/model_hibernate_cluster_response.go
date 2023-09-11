package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HibernateClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o HibernateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HibernateClusterResponse struct{}"
	}

	return strings.Join([]string{"HibernateClusterResponse", string(data)}, " ")
}
