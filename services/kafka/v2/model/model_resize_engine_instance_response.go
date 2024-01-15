package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResizeEngineInstanceResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeEngineInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeEngineInstanceResponse struct{}"
	}

	return strings.Join([]string{"ResizeEngineInstanceResponse", string(data)}, " ")
}