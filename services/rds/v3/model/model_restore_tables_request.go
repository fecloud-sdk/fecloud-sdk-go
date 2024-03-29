package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreTablesRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *RestoreTablesRequestBody `json:"body,omitempty"`
}

func (o RestoreTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTablesRequest struct{}"
	}

	return strings.Join([]string{"RestoreTablesRequest", string(data)}, " ")
}
