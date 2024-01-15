package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateEndpointWhiteResponse struct {
	Id *string `json:"id,omitempty"`

	ServiceType *UpdateEndpointWhiteResponseServiceType `json:"service_type,omitempty"`

	Status *UpdateEndpointWhiteResponseStatus `json:"status,omitempty"`

	Ip *string `json:"ip,omitempty"`

	ActiveStatus *[]string `json:"active_status,omitempty"`

	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	MarkerId *int32 `json:"marker_id,omitempty"`

	EndpointServiceId *string `json:"endpoint_service_id,omitempty"`

	EnableDns *bool `json:"enable_dns,omitempty"`

	DnsNames *[]string `json:"dns_names,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	Whitelist *[]string `json:"whitelist,omitempty"`

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o UpdateEndpointWhiteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointWhiteResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointWhiteResponse", string(data)}, " ")
}

type UpdateEndpointWhiteResponseServiceType struct {
	value string
}

type UpdateEndpointWhiteResponseServiceTypeEnum struct {
	INTERFACE UpdateEndpointWhiteResponseServiceType
	GATEWAY   UpdateEndpointWhiteResponseServiceType
}

func GetUpdateEndpointWhiteResponseServiceTypeEnum() UpdateEndpointWhiteResponseServiceTypeEnum {
	return UpdateEndpointWhiteResponseServiceTypeEnum{
		INTERFACE: UpdateEndpointWhiteResponseServiceType{
			value: "interface",
		},
		GATEWAY: UpdateEndpointWhiteResponseServiceType{
			value: "gateway",
		},
	}
}

func (c UpdateEndpointWhiteResponseServiceType) Value() string {
	return c.value
}

func (c UpdateEndpointWhiteResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointWhiteResponseServiceType) UnmarshalJSON(b []byte) error {
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

type UpdateEndpointWhiteResponseStatus struct {
	value string
}

type UpdateEndpointWhiteResponseStatusEnum struct {
	PENDING_ACCEPTANCE UpdateEndpointWhiteResponseStatus
	CREATING           UpdateEndpointWhiteResponseStatus
	ACCEPTED           UpdateEndpointWhiteResponseStatus
	REJECTED           UpdateEndpointWhiteResponseStatus
	FAILED             UpdateEndpointWhiteResponseStatus
	DELETING           UpdateEndpointWhiteResponseStatus
}

func GetUpdateEndpointWhiteResponseStatusEnum() UpdateEndpointWhiteResponseStatusEnum {
	return UpdateEndpointWhiteResponseStatusEnum{
		PENDING_ACCEPTANCE: UpdateEndpointWhiteResponseStatus{
			value: "pendingAcceptance",
		},
		CREATING: UpdateEndpointWhiteResponseStatus{
			value: "creating",
		},
		ACCEPTED: UpdateEndpointWhiteResponseStatus{
			value: "accepted",
		},
		REJECTED: UpdateEndpointWhiteResponseStatus{
			value: "rejected",
		},
		FAILED: UpdateEndpointWhiteResponseStatus{
			value: "failed",
		},
		DELETING: UpdateEndpointWhiteResponseStatus{
			value: "deleting",
		},
	}
}

func (c UpdateEndpointWhiteResponseStatus) Value() string {
	return c.value
}

func (c UpdateEndpointWhiteResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointWhiteResponseStatus) UnmarshalJSON(b []byte) error {
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
