package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DatastoreResult struct {
	NodeType string `json:"node_type"`

	Version string `json:"version"`
}

func (o DatastoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreResult struct{}"
	}

	return strings.Join([]string{"DatastoreResult", string(data)}, " ")
}
