package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPostgresqlParamValueResponse struct {
	Name *string `json:"name,omitempty"`

	Value *string `json:"value,omitempty"`

	RestartRequired *bool `json:"restart_required,omitempty"`

	ValueRange *string `json:"value_range,omitempty"`

	Type *string `json:"type,omitempty"`

	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPostgresqlParamValueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPostgresqlParamValueResponse struct{}"
	}

	return strings.Join([]string{"ShowPostgresqlParamValueResponse", string(data)}, " ")
}
