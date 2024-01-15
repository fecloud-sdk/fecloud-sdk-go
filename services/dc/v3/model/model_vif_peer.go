package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type VifPeer struct {
	Id *string `json:"id,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	AddressFamily *string `json:"address_family,omitempty"`

	LocalGatewayIp *string `json:"local_gateway_ip,omitempty"`

	RemoteGatewayIp *string `json:"remote_gateway_ip,omitempty"`

	RouteMode *VifPeerRouteMode `json:"route_mode,omitempty"`

	BgpAsn *int32 `json:"bgp_asn,omitempty"`

	BgpMd5 *string `json:"bgp_md5,omitempty"`

	RemoteEpGroup *[]string `json:"remote_ep_group,omitempty"`

	ServiceEpGroup *[]string `json:"service_ep_group,omitempty"`

	DeviceId *string `json:"device_id,omitempty"`

	BgpRouteLimit *int32 `json:"bgp_route_limit,omitempty"`

	BgpStatus *string `json:"bgp_status,omitempty"`

	Status *string `json:"status,omitempty"`

	VifId *string `json:"vif_id,omitempty"`

	ReceiveRouteNum *int32 `json:"receive_route_num,omitempty"`
}

func (o VifPeer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VifPeer struct{}"
	}

	return strings.Join([]string{"VifPeer", string(data)}, " ")
}

type VifPeerRouteMode struct {
	value string
}

type VifPeerRouteModeEnum struct {
	BGP    VifPeerRouteMode
	STATIC VifPeerRouteMode
}

func GetVifPeerRouteModeEnum() VifPeerRouteModeEnum {
	return VifPeerRouteModeEnum{
		BGP: VifPeerRouteMode{
			value: "bgp",
		},
		STATIC: VifPeerRouteMode{
			value: "static",
		},
	}
}

func (c VifPeerRouteMode) Value() string {
	return c.value
}

func (c VifPeerRouteMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VifPeerRouteMode) UnmarshalJSON(b []byte) error {
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
