package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowGroupsResponse struct {
	Group          *ShowGroupsRespGroup `json:"group,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupsResponse", string(data)}, " ")
}
