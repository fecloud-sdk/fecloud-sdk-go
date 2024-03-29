package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowPublicipResponse Response Object
type ShowPublicipResponse struct {
	Publicip       *PublicipShowResp `json:"publicip,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowPublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicipResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicipResponse", string(data)}, " ")
}
