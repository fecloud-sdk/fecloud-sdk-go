package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCoordinatorsRespCoordinators struct {
	GroupId *string `json:"group_id,omitempty"`

	Id *int32 `json:"id,omitempty"`

	Host *string `json:"host,omitempty"`

	Port *int32 `json:"port,omitempty"`
}

func (o ShowCoordinatorsRespCoordinators) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCoordinatorsRespCoordinators struct{}"
	}

	return strings.Join([]string{"ShowCoordinatorsRespCoordinators", string(data)}, " ")
}
