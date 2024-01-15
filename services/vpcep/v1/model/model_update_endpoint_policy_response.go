package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateEndpointPolicyResponse struct {
	Id *string `json:"id,omitempty"`

	ServiceType *UpdateEndpointPolicyResponseServiceType `json:"service_type,omitempty"`

	Status *UpdateEndpointPolicyResponseStatus `json:"status,omitempty"`

	ActiveStatus *[]string `json:"active_status,omitempty"`

	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	MarkerId *int32 `json:"marker_id,omitempty"`

	EndpointServiceId *string `json:"endpoint_service_id,omitempty"`

	EnableDns *bool `json:"enable_dns,omitempty"`

	DnsNames *[]string `json:"dns_names,omitempty"`

	Ip *string `json:"ip,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	Error *[]QueryError `json:"error,omitempty"`

	Whitelist *[]string `json:"whitelist,omitempty"`

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	Routetables *[]string `json:"routetables,omitempty"`

	Description *string `json:"description,omitempty"`

	PolicyStatement *[]PolicyStatement `json:"policy_statement,omitempty"`

	EndpointPoolId *string `json:"endpoint_pool_id,omitempty"`

	PublicBorderGroup *string `json:"public_border_group,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o UpdateEndpointPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointPolicyResponse", string(data)}, " ")
}

type UpdateEndpointPolicyResponseServiceType struct {
	value string
}

type UpdateEndpointPolicyResponseServiceTypeEnum struct {
	INTERFACE UpdateEndpointPolicyResponseServiceType
	GATEWAY   UpdateEndpointPolicyResponseServiceType
}

func GetUpdateEndpointPolicyResponseServiceTypeEnum() UpdateEndpointPolicyResponseServiceTypeEnum {
	return UpdateEndpointPolicyResponseServiceTypeEnum{
		INTERFACE: UpdateEndpointPolicyResponseServiceType{
			value: "interface",
		},
		GATEWAY: UpdateEndpointPolicyResponseServiceType{
			value: "gateway",
		},
	}
}

func (c UpdateEndpointPolicyResponseServiceType) Value() string {
	return c.value
}

func (c UpdateEndpointPolicyResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointPolicyResponseServiceType) UnmarshalJSON(b []byte) error {
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

type UpdateEndpointPolicyResponseStatus struct {
	value string
}

type UpdateEndpointPolicyResponseStatusEnum struct {
	PENDING_ACCEPTANCE UpdateEndpointPolicyResponseStatus
	CREATING           UpdateEndpointPolicyResponseStatus
	ACCEPTED           UpdateEndpointPolicyResponseStatus
	REJECTED           UpdateEndpointPolicyResponseStatus
	FAILED             UpdateEndpointPolicyResponseStatus
	DELETING           UpdateEndpointPolicyResponseStatus
}

func GetUpdateEndpointPolicyResponseStatusEnum() UpdateEndpointPolicyResponseStatusEnum {
	return UpdateEndpointPolicyResponseStatusEnum{
		PENDING_ACCEPTANCE: UpdateEndpointPolicyResponseStatus{
			value: "pendingAcceptance",
		},
		CREATING: UpdateEndpointPolicyResponseStatus{
			value: "creating",
		},
		ACCEPTED: UpdateEndpointPolicyResponseStatus{
			value: "accepted",
		},
		REJECTED: UpdateEndpointPolicyResponseStatus{
			value: "rejected",
		},
		FAILED: UpdateEndpointPolicyResponseStatus{
			value: "failed",
		},
		DELETING: UpdateEndpointPolicyResponseStatus{
			value: "deleting",
		},
	}
}

func (c UpdateEndpointPolicyResponseStatus) Value() string {
	return c.value
}

func (c UpdateEndpointPolicyResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointPolicyResponseStatus) UnmarshalJSON(b []byte) error {
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
