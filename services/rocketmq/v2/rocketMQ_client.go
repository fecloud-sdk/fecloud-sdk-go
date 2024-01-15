package v2

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/rocketmq/v2/model"
)

type RocketMQClient struct {
	HcClient *http_client.HcHttpClient
}

func NewRocketMQClient(hcClient *http_client.HcHttpClient) *RocketMQClient {
	return &RocketMQClient{HcClient: hcClient}
}

func RocketMQClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *RocketMQClient) BatchCreateOrDeleteRocketmqTag(request *model.BatchCreateOrDeleteRocketmqTagRequest) (*model.BatchCreateOrDeleteRocketmqTagResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteRocketmqTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteRocketmqTagResponse), nil
	}
}

func (c *RocketMQClient) BatchCreateOrDeleteRocketmqTagInvoker(request *model.BatchCreateOrDeleteRocketmqTagRequest) *BatchCreateOrDeleteRocketmqTagInvoker {
	requestDef := GenReqDefForBatchCreateOrDeleteRocketmqTag()
	return &BatchCreateOrDeleteRocketmqTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) BatchDeleteInstances(request *model.BatchDeleteInstancesRequest) (*model.BatchDeleteInstancesResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstancesResponse), nil
	}
}

func (c *RocketMQClient) BatchDeleteInstancesInvoker(request *model.BatchDeleteInstancesRequest) *BatchDeleteInstancesInvoker {
	requestDef := GenReqDefForBatchDeleteInstances()
	return &BatchDeleteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) BatchUpdateConsumerGroup(request *model.BatchUpdateConsumerGroupRequest) (*model.BatchUpdateConsumerGroupResponse, error) {
	requestDef := GenReqDefForBatchUpdateConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateConsumerGroupResponse), nil
	}
}

func (c *RocketMQClient) BatchUpdateConsumerGroupInvoker(request *model.BatchUpdateConsumerGroupRequest) *BatchUpdateConsumerGroupInvoker {
	requestDef := GenReqDefForBatchUpdateConsumerGroup()
	return &BatchUpdateConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) CreateConsumerGroupOrBatchDeleteConsumerGroup(request *model.CreateConsumerGroupOrBatchDeleteConsumerGroupRequest) (*model.CreateConsumerGroupOrBatchDeleteConsumerGroupResponse, error) {
	requestDef := GenReqDefForCreateConsumerGroupOrBatchDeleteConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConsumerGroupOrBatchDeleteConsumerGroupResponse), nil
	}
}

func (c *RocketMQClient) CreateConsumerGroupOrBatchDeleteConsumerGroupInvoker(request *model.CreateConsumerGroupOrBatchDeleteConsumerGroupRequest) *CreateConsumerGroupOrBatchDeleteConsumerGroupInvoker {
	requestDef := GenReqDefForCreateConsumerGroupOrBatchDeleteConsumerGroup()
	return &CreateConsumerGroupOrBatchDeleteConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) CreateInstanceByEngine(request *model.CreateInstanceByEngineRequest) (*model.CreateInstanceByEngineResponse, error) {
	requestDef := GenReqDefForCreateInstanceByEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceByEngineResponse), nil
	}
}

func (c *RocketMQClient) CreateInstanceByEngineInvoker(request *model.CreateInstanceByEngineRequest) *CreateInstanceByEngineInvoker {
	requestDef := GenReqDefForCreateInstanceByEngine()
	return &CreateInstanceByEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) CreatePostPaidInstance(request *model.CreatePostPaidInstanceRequest) (*model.CreatePostPaidInstanceResponse, error) {
	requestDef := GenReqDefForCreatePostPaidInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidInstanceResponse), nil
	}
}

func (c *RocketMQClient) CreatePostPaidInstanceInvoker(request *model.CreatePostPaidInstanceRequest) *CreatePostPaidInstanceInvoker {
	requestDef := GenReqDefForCreatePostPaidInstance()
	return &CreatePostPaidInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) CreateRocketMqMigrationTask(request *model.CreateRocketMqMigrationTaskRequest) (*model.CreateRocketMqMigrationTaskResponse, error) {
	requestDef := GenReqDefForCreateRocketMqMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRocketMqMigrationTaskResponse), nil
	}
}

func (c *RocketMQClient) CreateRocketMqMigrationTaskInvoker(request *model.CreateRocketMqMigrationTaskRequest) *CreateRocketMqMigrationTaskInvoker {
	requestDef := GenReqDefForCreateRocketMqMigrationTask()
	return &CreateRocketMqMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) CreateUser(request *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	requestDef := GenReqDefForCreateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserResponse), nil
	}
}

func (c *RocketMQClient) CreateUserInvoker(request *model.CreateUserRequest) *CreateUserInvoker {
	requestDef := GenReqDefForCreateUser()
	return &CreateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) DeleteConsumerGroup(request *model.DeleteConsumerGroupRequest) (*model.DeleteConsumerGroupResponse, error) {
	requestDef := GenReqDefForDeleteConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConsumerGroupResponse), nil
	}
}

func (c *RocketMQClient) DeleteConsumerGroupInvoker(request *model.DeleteConsumerGroupRequest) *DeleteConsumerGroupInvoker {
	requestDef := GenReqDefForDeleteConsumerGroup()
	return &DeleteConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

func (c *RocketMQClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) DeleteRocketMqMigrationTask(request *model.DeleteRocketMqMigrationTaskRequest) (*model.DeleteRocketMqMigrationTaskResponse, error) {
	requestDef := GenReqDefForDeleteRocketMqMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRocketMqMigrationTaskResponse), nil
	}
}

func (c *RocketMQClient) DeleteRocketMqMigrationTaskInvoker(request *model.DeleteRocketMqMigrationTaskRequest) *DeleteRocketMqMigrationTaskInvoker {
	requestDef := GenReqDefForDeleteRocketMqMigrationTask()
	return &DeleteRocketMqMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) DeleteUser(request *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	requestDef := GenReqDefForDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserResponse), nil
	}
}

func (c *RocketMQClient) DeleteUserInvoker(request *model.DeleteUserRequest) *DeleteUserInvoker {
	requestDef := GenReqDefForDeleteUser()
	return &DeleteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ExportDlqMessage(request *model.ExportDlqMessageRequest) (*model.ExportDlqMessageResponse, error) {
	requestDef := GenReqDefForExportDlqMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportDlqMessageResponse), nil
	}
}

func (c *RocketMQClient) ExportDlqMessageInvoker(request *model.ExportDlqMessageRequest) *ExportDlqMessageInvoker {
	requestDef := GenReqDefForExportDlqMessage()
	return &ExportDlqMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

func (c *RocketMQClient) ListAvailableZonesInvoker(request *model.ListAvailableZonesRequest) *ListAvailableZonesInvoker {
	requestDef := GenReqDefForListAvailableZones()
	return &ListAvailableZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListBrokers(request *model.ListBrokersRequest) (*model.ListBrokersResponse, error) {
	requestDef := GenReqDefForListBrokers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBrokersResponse), nil
	}
}

func (c *RocketMQClient) ListBrokersInvoker(request *model.ListBrokersRequest) *ListBrokersInvoker {
	requestDef := GenReqDefForListBrokers()
	return &ListBrokersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListConsumeGroupAccessPolicy(request *model.ListConsumeGroupAccessPolicyRequest) (*model.ListConsumeGroupAccessPolicyResponse, error) {
	requestDef := GenReqDefForListConsumeGroupAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConsumeGroupAccessPolicyResponse), nil
	}
}

func (c *RocketMQClient) ListConsumeGroupAccessPolicyInvoker(request *model.ListConsumeGroupAccessPolicyRequest) *ListConsumeGroupAccessPolicyInvoker {
	requestDef := GenReqDefForListConsumeGroupAccessPolicy()
	return &ListConsumeGroupAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListInstanceConsumerGroups(request *model.ListInstanceConsumerGroupsRequest) (*model.ListInstanceConsumerGroupsResponse, error) {
	requestDef := GenReqDefForListInstanceConsumerGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceConsumerGroupsResponse), nil
	}
}

func (c *RocketMQClient) ListInstanceConsumerGroupsInvoker(request *model.ListInstanceConsumerGroupsRequest) *ListInstanceConsumerGroupsInvoker {
	requestDef := GenReqDefForListInstanceConsumerGroups()
	return &ListInstanceConsumerGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

func (c *RocketMQClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListMessageTrace(request *model.ListMessageTraceRequest) (*model.ListMessageTraceResponse, error) {
	requestDef := GenReqDefForListMessageTrace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessageTraceResponse), nil
	}
}

func (c *RocketMQClient) ListMessageTraceInvoker(request *model.ListMessageTraceRequest) *ListMessageTraceInvoker {
	requestDef := GenReqDefForListMessageTrace()
	return &ListMessageTraceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListMessages(request *model.ListMessagesRequest) (*model.ListMessagesResponse, error) {
	requestDef := GenReqDefForListMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessagesResponse), nil
	}
}

func (c *RocketMQClient) ListMessagesInvoker(request *model.ListMessagesRequest) *ListMessagesInvoker {
	requestDef := GenReqDefForListMessages()
	return &ListMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListRocketMqMigrationTask(request *model.ListRocketMqMigrationTaskRequest) (*model.ListRocketMqMigrationTaskResponse, error) {
	requestDef := GenReqDefForListRocketMqMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRocketMqMigrationTaskResponse), nil
	}
}

func (c *RocketMQClient) ListRocketMqMigrationTaskInvoker(request *model.ListRocketMqMigrationTaskRequest) *ListRocketMqMigrationTaskInvoker {
	requestDef := GenReqDefForListRocketMqMigrationTask()
	return &ListRocketMqMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListTopicAccessPolicy(request *model.ListTopicAccessPolicyRequest) (*model.ListTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForListTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicAccessPolicyResponse), nil
	}
}

func (c *RocketMQClient) ListTopicAccessPolicyInvoker(request *model.ListTopicAccessPolicyRequest) *ListTopicAccessPolicyInvoker {
	requestDef := GenReqDefForListTopicAccessPolicy()
	return &ListTopicAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListUser(request *model.ListUserRequest) (*model.ListUserResponse, error) {
	requestDef := GenReqDefForListUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserResponse), nil
	}
}

func (c *RocketMQClient) ListUserInvoker(request *model.ListUserRequest) *ListUserInvoker {
	requestDef := GenReqDefForListUser()
	return &ListUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ResetConsumeOffset(request *model.ResetConsumeOffsetRequest) (*model.ResetConsumeOffsetResponse, error) {
	requestDef := GenReqDefForResetConsumeOffset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetConsumeOffsetResponse), nil
	}
}

func (c *RocketMQClient) ResetConsumeOffsetInvoker(request *model.ResetConsumeOffsetRequest) *ResetConsumeOffsetInvoker {
	requestDef := GenReqDefForResetConsumeOffset()
	return &ResetConsumeOffsetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) SendDlqMessage(request *model.SendDlqMessageRequest) (*model.SendDlqMessageResponse, error) {
	requestDef := GenReqDefForSendDlqMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendDlqMessageResponse), nil
	}
}

func (c *RocketMQClient) SendDlqMessageInvoker(request *model.SendDlqMessageRequest) *SendDlqMessageInvoker {
	requestDef := GenReqDefForSendDlqMessage()
	return &SendDlqMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ShowConsumerConnections(request *model.ShowConsumerConnectionsRequest) (*model.ShowConsumerConnectionsResponse, error) {
	requestDef := GenReqDefForShowConsumerConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConsumerConnectionsResponse), nil
	}
}

func (c *RocketMQClient) ShowConsumerConnectionsInvoker(request *model.ShowConsumerConnectionsRequest) *ShowConsumerConnectionsInvoker {
	requestDef := GenReqDefForShowConsumerConnections()
	return &ShowConsumerConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ShowConsumerListOrDetails(request *model.ShowConsumerListOrDetailsRequest) (*model.ShowConsumerListOrDetailsResponse, error) {
	requestDef := GenReqDefForShowConsumerListOrDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConsumerListOrDetailsResponse), nil
	}
}

func (c *RocketMQClient) ShowConsumerListOrDetailsInvoker(request *model.ShowConsumerListOrDetailsRequest) *ShowConsumerListOrDetailsInvoker {
	requestDef := GenReqDefForShowConsumerListOrDetails()
	return &ShowConsumerListOrDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ShowGroup(request *model.ShowGroupRequest) (*model.ShowGroupResponse, error) {
	requestDef := GenReqDefForShowGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupResponse), nil
	}
}

func (c *RocketMQClient) ShowGroupInvoker(request *model.ShowGroupRequest) *ShowGroupInvoker {
	requestDef := GenReqDefForShowGroup()
	return &ShowGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

func (c *RocketMQClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ShowRocketmqProjectTags(request *model.ShowRocketmqProjectTagsRequest) (*model.ShowRocketmqProjectTagsResponse, error) {
	requestDef := GenReqDefForShowRocketmqProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRocketmqProjectTagsResponse), nil
	}
}

func (c *RocketMQClient) ShowRocketmqProjectTagsInvoker(request *model.ShowRocketmqProjectTagsRequest) *ShowRocketmqProjectTagsInvoker {
	requestDef := GenReqDefForShowRocketmqProjectTags()
	return &ShowRocketmqProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ShowRocketmqTags(request *model.ShowRocketmqTagsRequest) (*model.ShowRocketmqTagsResponse, error) {
	requestDef := GenReqDefForShowRocketmqTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRocketmqTagsResponse), nil
	}
}

func (c *RocketMQClient) ShowRocketmqTagsInvoker(request *model.ShowRocketmqTagsRequest) *ShowRocketmqTagsInvoker {
	requestDef := GenReqDefForShowRocketmqTags()
	return &ShowRocketmqTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ShowUser(request *model.ShowUserRequest) (*model.ShowUserResponse, error) {
	requestDef := GenReqDefForShowUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserResponse), nil
	}
}

func (c *RocketMQClient) ShowUserInvoker(request *model.ShowUserRequest) *ShowUserInvoker {
	requestDef := GenReqDefForShowUser()
	return &ShowUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) UpdateConsumerGroup(request *model.UpdateConsumerGroupRequest) (*model.UpdateConsumerGroupResponse, error) {
	requestDef := GenReqDefForUpdateConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConsumerGroupResponse), nil
	}
}

func (c *RocketMQClient) UpdateConsumerGroupInvoker(request *model.UpdateConsumerGroupRequest) *UpdateConsumerGroupInvoker {
	requestDef := GenReqDefForUpdateConsumerGroup()
	return &UpdateConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

func (c *RocketMQClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) UpdateUser(request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	requestDef := GenReqDefForUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserResponse), nil
	}
}

func (c *RocketMQClient) UpdateUserInvoker(request *model.UpdateUserRequest) *UpdateUserInvoker {
	requestDef := GenReqDefForUpdateUser()
	return &UpdateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ValidateConsumedMessage(request *model.ValidateConsumedMessageRequest) (*model.ValidateConsumedMessageResponse, error) {
	requestDef := GenReqDefForValidateConsumedMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateConsumedMessageResponse), nil
	}
}

func (c *RocketMQClient) ValidateConsumedMessageInvoker(request *model.ValidateConsumedMessageRequest) *ValidateConsumedMessageInvoker {
	requestDef := GenReqDefForValidateConsumedMessage()
	return &ValidateConsumedMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) CreateTopicOrBatchDeleteTopic(request *model.CreateTopicOrBatchDeleteTopicRequest) (*model.CreateTopicOrBatchDeleteTopicResponse, error) {
	requestDef := GenReqDefForCreateTopicOrBatchDeleteTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTopicOrBatchDeleteTopicResponse), nil
	}
}

func (c *RocketMQClient) CreateTopicOrBatchDeleteTopicInvoker(request *model.CreateTopicOrBatchDeleteTopicRequest) *CreateTopicOrBatchDeleteTopicInvoker {
	requestDef := GenReqDefForCreateTopicOrBatchDeleteTopic()
	return &CreateTopicOrBatchDeleteTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) DeleteTopic(request *model.DeleteTopicRequest) (*model.DeleteTopicResponse, error) {
	requestDef := GenReqDefForDeleteTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTopicResponse), nil
	}
}

func (c *RocketMQClient) DeleteTopicInvoker(request *model.DeleteTopicRequest) *DeleteTopicInvoker {
	requestDef := GenReqDefForDeleteTopic()
	return &DeleteTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListConsumerGroupOfTopic(request *model.ListConsumerGroupOfTopicRequest) (*model.ListConsumerGroupOfTopicResponse, error) {
	requestDef := GenReqDefForListConsumerGroupOfTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConsumerGroupOfTopicResponse), nil
	}
}

func (c *RocketMQClient) ListConsumerGroupOfTopicInvoker(request *model.ListConsumerGroupOfTopicRequest) *ListConsumerGroupOfTopicInvoker {
	requestDef := GenReqDefForListConsumerGroupOfTopic()
	return &ListConsumerGroupOfTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ListRocketInstanceTopics(request *model.ListRocketInstanceTopicsRequest) (*model.ListRocketInstanceTopicsResponse, error) {
	requestDef := GenReqDefForListRocketInstanceTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRocketInstanceTopicsResponse), nil
	}
}

func (c *RocketMQClient) ListRocketInstanceTopicsInvoker(request *model.ListRocketInstanceTopicsRequest) *ListRocketInstanceTopicsInvoker {
	requestDef := GenReqDefForListRocketInstanceTopics()
	return &ListRocketInstanceTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ShowOneTopic(request *model.ShowOneTopicRequest) (*model.ShowOneTopicResponse, error) {
	requestDef := GenReqDefForShowOneTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOneTopicResponse), nil
	}
}

func (c *RocketMQClient) ShowOneTopicInvoker(request *model.ShowOneTopicRequest) *ShowOneTopicInvoker {
	requestDef := GenReqDefForShowOneTopic()
	return &ShowOneTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) ShowTopicStatus(request *model.ShowTopicStatusRequest) (*model.ShowTopicStatusResponse, error) {
	requestDef := GenReqDefForShowTopicStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopicStatusResponse), nil
	}
}

func (c *RocketMQClient) ShowTopicStatusInvoker(request *model.ShowTopicStatusRequest) *ShowTopicStatusInvoker {
	requestDef := GenReqDefForShowTopicStatus()
	return &ShowTopicStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RocketMQClient) UpdateTopic(request *model.UpdateTopicRequest) (*model.UpdateTopicResponse, error) {
	requestDef := GenReqDefForUpdateTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicResponse), nil
	}
}

func (c *RocketMQClient) UpdateTopicInvoker(request *model.UpdateTopicRequest) *UpdateTopicInvoker {
	requestDef := GenReqDefForUpdateTopic()
	return &UpdateTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
