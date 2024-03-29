package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteInstanceResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceResponse", string(data)}, " ")
}
