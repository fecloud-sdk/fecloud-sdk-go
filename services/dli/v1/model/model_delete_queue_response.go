package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteQueueResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueueResponse struct{}"
	}

	return strings.Join([]string{"DeleteQueueResponse", string(data)}, " ")
}
