package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateEndpointServiceRequestBody struct {
	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`

	ServiceName *string `json:"service_name,omitempty"`

	Ports *[]PortList `json:"ports,omitempty"`

	PortId *string `json:"port_id,omitempty"`

	TcpProxy *UpdateEndpointServiceRequestBodyTcpProxy `json:"tcp_proxy,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o UpdateEndpointServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceRequestBody", string(data)}, " ")
}

type UpdateEndpointServiceRequestBodyTcpProxy struct {
	value string
}

type UpdateEndpointServiceRequestBodyTcpProxyEnum struct {
	CLOSE      UpdateEndpointServiceRequestBodyTcpProxy
	TOA_OPEN   UpdateEndpointServiceRequestBodyTcpProxy
	PROXY_OPEN UpdateEndpointServiceRequestBodyTcpProxy
	OPEN       UpdateEndpointServiceRequestBodyTcpProxy
	PROXY_VNI  UpdateEndpointServiceRequestBodyTcpProxy
}

func GetUpdateEndpointServiceRequestBodyTcpProxyEnum() UpdateEndpointServiceRequestBodyTcpProxyEnum {
	return UpdateEndpointServiceRequestBodyTcpProxyEnum{
		CLOSE: UpdateEndpointServiceRequestBodyTcpProxy{
			value: "close",
		},
		TOA_OPEN: UpdateEndpointServiceRequestBodyTcpProxy{
			value: "toa_open",
		},
		PROXY_OPEN: UpdateEndpointServiceRequestBodyTcpProxy{
			value: "proxy_open",
		},
		OPEN: UpdateEndpointServiceRequestBodyTcpProxy{
			value: "open",
		},
		PROXY_VNI: UpdateEndpointServiceRequestBodyTcpProxy{
			value: "proxy_vni",
		},
	}
}

func (c UpdateEndpointServiceRequestBodyTcpProxy) Value() string {
	return c.value
}

func (c UpdateEndpointServiceRequestBodyTcpProxy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointServiceRequestBodyTcpProxy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
