package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowEngineInstanceExtendProductInfoResponse struct {
	Engine *string `json:"engine,omitempty"`

	Versions *[]string `json:"versions,omitempty"`

	Products       *[]ExtendProductInfoEntity `json:"products,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowEngineInstanceExtendProductInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineInstanceExtendProductInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowEngineInstanceExtendProductInfoResponse", string(data)}, " ")
}
