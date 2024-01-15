package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteBackgroundTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBackgroundTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackgroundTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackgroundTaskResponse", string(data)}, " ")
}
