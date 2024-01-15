package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListGroupReplicationInfoResponse struct {
	GroupList *[]InstanceGroupListInfo `json:"group_list,omitempty"`

	GroupCount     *int32 `json:"group_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListGroupReplicationInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupReplicationInfoResponse struct{}"
	}

	return strings.Join([]string{"ListGroupReplicationInfoResponse", string(data)}, " ")
}
