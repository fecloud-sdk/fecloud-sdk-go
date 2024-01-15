package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSqlResultResponse struct {
	Id *string `json:"id,omitempty"`

	Message *string `json:"message,omitempty"`

	Statement *string `json:"statement,omitempty"`

	Status *string `json:"status,omitempty"`

	ResultLocation *string `json:"result_location,omitempty"`

	Content        *[][]string `json:"content,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowSqlResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlResultResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlResultResponse", string(data)}, " ")
}
