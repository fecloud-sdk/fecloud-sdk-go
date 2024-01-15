package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateTopicReplicaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTopicReplicaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicReplicaResponse struct{}"
	}

	return strings.Join([]string{"UpdateTopicReplicaResponse", string(data)}, " ")
}
