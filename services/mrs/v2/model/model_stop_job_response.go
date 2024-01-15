package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type StopJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobResponse struct{}"
	}

	return strings.Join([]string{"StopJobResponse", string(data)}, " ")
}
