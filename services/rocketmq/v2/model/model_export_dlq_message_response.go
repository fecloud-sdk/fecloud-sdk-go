package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExportDlqMessageResponse struct {
	Body           *[]Message `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ExportDlqMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDlqMessageResponse struct{}"
	}

	return strings.Join([]string{"ExportDlqMessageResponse", string(data)}, " ")
}
