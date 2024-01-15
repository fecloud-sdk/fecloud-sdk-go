package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type DeleteEndpointPolicyResponse struct {
	Id *string `json:"id,omitempty"`

	ServiceType *DeleteEndpointPolicyResponseServiceType `json:"service_type,omitempty"`

	Status *DeleteEndpointPolicyResponseStatus `json:"status,omitempty"`

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

func (o DeleteEndpointPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteEndpointPolicyResponse", string(data)}, " ")
}

type DeleteEndpointPolicyResponseServiceType struct {
	value string
}

type DeleteEndpointPolicyResponseServiceTypeEnum struct {
	INTERFACE DeleteEndpointPolicyResponseServiceType
	GATEWAY   DeleteEndpointPolicyResponseServiceType
}

func GetDeleteEndpointPolicyResponseServiceTypeEnum() DeleteEndpointPolicyResponseServiceTypeEnum {
	return DeleteEndpointPolicyResponseServiceTypeEnum{
		INTERFACE: DeleteEndpointPolicyResponseServiceType{
			value: "interface",
		},
		GATEWAY: DeleteEndpointPolicyResponseServiceType{
			value: "gateway",
		},
	}
}

func (c DeleteEndpointPolicyResponseServiceType) Value() string {
	return c.value
}

func (c DeleteEndpointPolicyResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteEndpointPolicyResponseServiceType) UnmarshalJSON(b []byte) error {
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

type DeleteEndpointPolicyResponseStatus struct {
	value string
}

type DeleteEndpointPolicyResponseStatusEnum struct {
	PENDING_ACCEPTANCE DeleteEndpointPolicyResponseStatus
	CREATING           DeleteEndpointPolicyResponseStatus
	ACCEPTED           DeleteEndpointPolicyResponseStatus
	REJECTED           DeleteEndpointPolicyResponseStatus
	FAILED             DeleteEndpointPolicyResponseStatus
	DELETING           DeleteEndpointPolicyResponseStatus
}

func GetDeleteEndpointPolicyResponseStatusEnum() DeleteEndpointPolicyResponseStatusEnum {
	return DeleteEndpointPolicyResponseStatusEnum{
		PENDING_ACCEPTANCE: DeleteEndpointPolicyResponseStatus{
			value: "pendingAcceptance",
		},
		CREATING: DeleteEndpointPolicyResponseStatus{
			value: "creating",
		},
		ACCEPTED: DeleteEndpointPolicyResponseStatus{
			value: "accepted",
		},
		REJECTED: DeleteEndpointPolicyResponseStatus{
			value: "rejected",
		},
		FAILED: DeleteEndpointPolicyResponseStatus{
			value: "failed",
		},
		DELETING: DeleteEndpointPolicyResponseStatus{
			value: "deleting",
		},
	}
}

func (c DeleteEndpointPolicyResponseStatus) Value() string {
	return c.value
}

func (c DeleteEndpointPolicyResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteEndpointPolicyResponseStatus) UnmarshalJSON(b []byte) error {
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
