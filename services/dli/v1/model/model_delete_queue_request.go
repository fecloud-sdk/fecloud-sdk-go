package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteQueueRequest struct {
	QueueName string `json:"queue_name"`
}

func (o DeleteQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueueRequest struct{}"
	}

	return strings.Join([]string{"DeleteQueueRequest", string(data)}, " ")
}
