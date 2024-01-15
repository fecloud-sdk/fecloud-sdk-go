package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListDrRelationsResponse struct {
	InstanceDrRelations *[]InstanceDrRelation `json:"instance_dr_relations,omitempty"`
	HttpStatusCode      int                   `json:"-"`
}

func (o ListDrRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDrRelationsResponse struct{}"
	}

	return strings.Join([]string{"ListDrRelationsResponse", string(data)}, " ")
}
