package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchCreatePublicipTagsRequest Request Object
type BatchCreatePublicipTagsRequest struct {

	// 资源ID
	PublicipId string `json:"publicip_id"`

	Body *BatchCreatePublicipTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreatePublicipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicipTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicipTagsRequest", string(data)}, " ")
}
