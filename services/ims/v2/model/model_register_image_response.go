package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RegisterImageResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterImageResponse struct{}"
	}

	return strings.Join([]string{"RegisterImageResponse", string(data)}, " ")
}
