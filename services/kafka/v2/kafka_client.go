package v2

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/kafka/v2/model"
)

type KafkaClient struct {
	HcClient *http_client.HcHttpClient
}

func NewKafkaClient(hcClient *http_client.HcHttpClient) *KafkaClient {
	return &KafkaClient{HcClient: hcClient}
}

func KafkaClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *KafkaClient) BatchCreateOrDeleteKafkaTag(request *model.BatchCreateOrDeleteKafkaTagRequest) (*model.BatchCreateOrDeleteKafkaTagResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteKafkaTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteKafkaTagResponse), nil
	}
}

func (c *KafkaClient) BatchCreateOrDeleteKafkaTagInvoker(request *model.BatchCreateOrDeleteKafkaTagRequest) *BatchCreateOrDeleteKafkaTagInvoker {
	requestDef := GenReqDefForBatchCreateOrDeleteKafkaTag()
	return &BatchCreateOrDeleteKafkaTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) BatchDeleteGroup(request *model.BatchDeleteGroupRequest) (*model.BatchDeleteGroupResponse, error) {
	requestDef := GenReqDefForBatchDeleteGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteGroupResponse), nil
	}
}

func (c *KafkaClient) BatchDeleteGroupInvoker(request *model.BatchDeleteGroupRequest) *BatchDeleteGroupInvoker {
	requestDef := GenReqDefForBatchDeleteGroup()
	return &BatchDeleteGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) BatchDeleteInstanceTopic(request *model.BatchDeleteInstanceTopicRequest) (*model.BatchDeleteInstanceTopicResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstanceTopicResponse), nil
	}
}

func (c *KafkaClient) BatchDeleteInstanceTopicInvoker(request *model.BatchDeleteInstanceTopicRequest) *BatchDeleteInstanceTopicInvoker {
	requestDef := GenReqDefForBatchDeleteInstanceTopic()
	return &BatchDeleteInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) BatchDeleteInstanceUsers(request *model.BatchDeleteInstanceUsersRequest) (*model.BatchDeleteInstanceUsersResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstanceUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstanceUsersResponse), nil
	}
}

func (c *KafkaClient) BatchDeleteInstanceUsersInvoker(request *model.BatchDeleteInstanceUsersRequest) *BatchDeleteInstanceUsersInvoker {
	requestDef := GenReqDefForBatchDeleteInstanceUsers()
	return &BatchDeleteInstanceUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) BatchRestartOrDeleteInstances(request *model.BatchRestartOrDeleteInstancesRequest) (*model.BatchRestartOrDeleteInstancesResponse, error) {
	requestDef := GenReqDefForBatchRestartOrDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRestartOrDeleteInstancesResponse), nil
	}
}

func (c *KafkaClient) BatchRestartOrDeleteInstancesInvoker(request *model.BatchRestartOrDeleteInstancesRequest) *BatchRestartOrDeleteInstancesInvoker {
	requestDef := GenReqDefForBatchRestartOrDeleteInstances()
	return &BatchRestartOrDeleteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) CloseKafkaManager(request *model.CloseKafkaManagerRequest) (*model.CloseKafkaManagerResponse, error) {
	requestDef := GenReqDefForCloseKafkaManager()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CloseKafkaManagerResponse), nil
	}
}

func (c *KafkaClient) CloseKafkaManagerInvoker(request *model.CloseKafkaManagerRequest) *CloseKafkaManagerInvoker {
	requestDef := GenReqDefForCloseKafkaManager()
	return &CloseKafkaManagerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) CreateConnector(request *model.CreateConnectorRequest) (*model.CreateConnectorResponse, error) {
	requestDef := GenReqDefForCreateConnector()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectorResponse), nil
	}
}

func (c *KafkaClient) CreateConnectorInvoker(request *model.CreateConnectorRequest) *CreateConnectorInvoker {
	requestDef := GenReqDefForCreateConnector()
	return &CreateConnectorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) CreateDeleteConnectorOrder(request *model.CreateDeleteConnectorOrderRequest) (*model.CreateDeleteConnectorOrderResponse, error) {
	requestDef := GenReqDefForCreateDeleteConnectorOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeleteConnectorOrderResponse), nil
	}
}

func (c *KafkaClient) CreateDeleteConnectorOrderInvoker(request *model.CreateDeleteConnectorOrderRequest) *CreateDeleteConnectorOrderInvoker {
	requestDef := GenReqDefForCreateDeleteConnectorOrder()
	return &CreateDeleteConnectorOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) CreateInstanceByEngine(request *model.CreateInstanceByEngineRequest) (*model.CreateInstanceByEngineResponse, error) {
	requestDef := GenReqDefForCreateInstanceByEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceByEngineResponse), nil
	}
}

func (c *KafkaClient) CreateInstanceByEngineInvoker(request *model.CreateInstanceByEngineRequest) *CreateInstanceByEngineInvoker {
	requestDef := GenReqDefForCreateInstanceByEngine()
	return &CreateInstanceByEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) CreateInstanceTopic(request *model.CreateInstanceTopicRequest) (*model.CreateInstanceTopicResponse, error) {
	requestDef := GenReqDefForCreateInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceTopicResponse), nil
	}
}

func (c *KafkaClient) CreateInstanceTopicInvoker(request *model.CreateInstanceTopicRequest) *CreateInstanceTopicInvoker {
	requestDef := GenReqDefForCreateInstanceTopic()
	return &CreateInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) CreateInstanceUser(request *model.CreateInstanceUserRequest) (*model.CreateInstanceUserResponse, error) {
	requestDef := GenReqDefForCreateInstanceUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceUserResponse), nil
	}
}

func (c *KafkaClient) CreateInstanceUserInvoker(request *model.CreateInstanceUserRequest) *CreateInstanceUserInvoker {
	requestDef := GenReqDefForCreateInstanceUser()
	return &CreateInstanceUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) CreateKafkaConsumerGroup(request *model.CreateKafkaConsumerGroupRequest) (*model.CreateKafkaConsumerGroupResponse, error) {
	requestDef := GenReqDefForCreateKafkaConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKafkaConsumerGroupResponse), nil
	}
}

func (c *KafkaClient) CreateKafkaConsumerGroupInvoker(request *model.CreateKafkaConsumerGroupRequest) *CreateKafkaConsumerGroupInvoker {
	requestDef := GenReqDefForCreateKafkaConsumerGroup()
	return &CreateKafkaConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) CreatePartition(request *model.CreatePartitionRequest) (*model.CreatePartitionResponse, error) {
	requestDef := GenReqDefForCreatePartition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePartitionResponse), nil
	}
}

func (c *KafkaClient) CreatePartitionInvoker(request *model.CreatePartitionRequest) *CreatePartitionInvoker {
	requestDef := GenReqDefForCreatePartition()
	return &CreatePartitionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) CreatePostPaidInstance(request *model.CreatePostPaidInstanceRequest) (*model.CreatePostPaidInstanceResponse, error) {
	requestDef := GenReqDefForCreatePostPaidInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidInstanceResponse), nil
	}
}

func (c *KafkaClient) CreatePostPaidInstanceInvoker(request *model.CreatePostPaidInstanceRequest) *CreatePostPaidInstanceInvoker {
	requestDef := GenReqDefForCreatePostPaidInstance()
	return &CreatePostPaidInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) CreateReassignmentTask(request *model.CreateReassignmentTaskRequest) (*model.CreateReassignmentTaskResponse, error) {
	requestDef := GenReqDefForCreateReassignmentTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReassignmentTaskResponse), nil
	}
}

func (c *KafkaClient) CreateReassignmentTaskInvoker(request *model.CreateReassignmentTaskRequest) *CreateReassignmentTaskInvoker {
	requestDef := GenReqDefForCreateReassignmentTask()
	return &CreateReassignmentTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) CreateSinkTask(request *model.CreateSinkTaskRequest) (*model.CreateSinkTaskResponse, error) {
	requestDef := GenReqDefForCreateSinkTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSinkTaskResponse), nil
	}
}

func (c *KafkaClient) CreateSinkTaskInvoker(request *model.CreateSinkTaskRequest) *CreateSinkTaskInvoker {
	requestDef := GenReqDefForCreateSinkTask()
	return &CreateSinkTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) DeleteBackgroundTask(request *model.DeleteBackgroundTaskRequest) (*model.DeleteBackgroundTaskResponse, error) {
	requestDef := GenReqDefForDeleteBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackgroundTaskResponse), nil
	}
}

func (c *KafkaClient) DeleteBackgroundTaskInvoker(request *model.DeleteBackgroundTaskRequest) *DeleteBackgroundTaskInvoker {
	requestDef := GenReqDefForDeleteBackgroundTask()
	return &DeleteBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) DeleteConnector(request *model.DeleteConnectorRequest) (*model.DeleteConnectorResponse, error) {
	requestDef := GenReqDefForDeleteConnector()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConnectorResponse), nil
	}
}

func (c *KafkaClient) DeleteConnectorInvoker(request *model.DeleteConnectorRequest) *DeleteConnectorInvoker {
	requestDef := GenReqDefForDeleteConnector()
	return &DeleteConnectorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

func (c *KafkaClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) DeleteSinkTask(request *model.DeleteSinkTaskRequest) (*model.DeleteSinkTaskResponse, error) {
	requestDef := GenReqDefForDeleteSinkTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSinkTaskResponse), nil
	}
}

func (c *KafkaClient) DeleteSinkTaskInvoker(request *model.DeleteSinkTaskRequest) *DeleteSinkTaskInvoker {
	requestDef := GenReqDefForDeleteSinkTask()
	return &DeleteSinkTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

func (c *KafkaClient) ListAvailableZonesInvoker(request *model.ListAvailableZonesRequest) *ListAvailableZonesInvoker {
	requestDef := GenReqDefForListAvailableZones()
	return &ListAvailableZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ListBackgroundTasks(request *model.ListBackgroundTasksRequest) (*model.ListBackgroundTasksResponse, error) {
	requestDef := GenReqDefForListBackgroundTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackgroundTasksResponse), nil
	}
}

func (c *KafkaClient) ListBackgroundTasksInvoker(request *model.ListBackgroundTasksRequest) *ListBackgroundTasksInvoker {
	requestDef := GenReqDefForListBackgroundTasks()
	return &ListBackgroundTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ListEngineProducts(request *model.ListEngineProductsRequest) (*model.ListEngineProductsResponse, error) {
	requestDef := GenReqDefForListEngineProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEngineProductsResponse), nil
	}
}

func (c *KafkaClient) ListEngineProductsInvoker(request *model.ListEngineProductsRequest) *ListEngineProductsInvoker {
	requestDef := GenReqDefForListEngineProducts()
	return &ListEngineProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ListInstanceConsumerGroups(request *model.ListInstanceConsumerGroupsRequest) (*model.ListInstanceConsumerGroupsResponse, error) {
	requestDef := GenReqDefForListInstanceConsumerGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceConsumerGroupsResponse), nil
	}
}

func (c *KafkaClient) ListInstanceConsumerGroupsInvoker(request *model.ListInstanceConsumerGroupsRequest) *ListInstanceConsumerGroupsInvoker {
	requestDef := GenReqDefForListInstanceConsumerGroups()
	return &ListInstanceConsumerGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ListInstanceTopics(request *model.ListInstanceTopicsRequest) (*model.ListInstanceTopicsResponse, error) {
	requestDef := GenReqDefForListInstanceTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTopicsResponse), nil
	}
}

func (c *KafkaClient) ListInstanceTopicsInvoker(request *model.ListInstanceTopicsRequest) *ListInstanceTopicsInvoker {
	requestDef := GenReqDefForListInstanceTopics()
	return &ListInstanceTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

func (c *KafkaClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ListProducts(request *model.ListProductsRequest) (*model.ListProductsResponse, error) {
	requestDef := GenReqDefForListProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductsResponse), nil
	}
}

func (c *KafkaClient) ListProductsInvoker(request *model.ListProductsRequest) *ListProductsInvoker {
	requestDef := GenReqDefForListProducts()
	return &ListProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ListSinkTasks(request *model.ListSinkTasksRequest) (*model.ListSinkTasksResponse, error) {
	requestDef := GenReqDefForListSinkTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSinkTasksResponse), nil
	}
}

func (c *KafkaClient) ListSinkTasksInvoker(request *model.ListSinkTasksRequest) *ListSinkTasksInvoker {
	requestDef := GenReqDefForListSinkTasks()
	return &ListSinkTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ListTopicPartitions(request *model.ListTopicPartitionsRequest) (*model.ListTopicPartitionsResponse, error) {
	requestDef := GenReqDefForListTopicPartitions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicPartitionsResponse), nil
	}
}

func (c *KafkaClient) ListTopicPartitionsInvoker(request *model.ListTopicPartitionsRequest) *ListTopicPartitionsInvoker {
	requestDef := GenReqDefForListTopicPartitions()
	return &ListTopicPartitionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ListTopicProducers(request *model.ListTopicProducersRequest) (*model.ListTopicProducersResponse, error) {
	requestDef := GenReqDefForListTopicProducers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicProducersResponse), nil
	}
}

func (c *KafkaClient) ListTopicProducersInvoker(request *model.ListTopicProducersRequest) *ListTopicProducersInvoker {
	requestDef := GenReqDefForListTopicProducers()
	return &ListTopicProducersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ResetManagerPassword(request *model.ResetManagerPasswordRequest) (*model.ResetManagerPasswordResponse, error) {
	requestDef := GenReqDefForResetManagerPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetManagerPasswordResponse), nil
	}
}

func (c *KafkaClient) ResetManagerPasswordInvoker(request *model.ResetManagerPasswordRequest) *ResetManagerPasswordInvoker {
	requestDef := GenReqDefForResetManagerPassword()
	return &ResetManagerPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ResetMessageOffset(request *model.ResetMessageOffsetRequest) (*model.ResetMessageOffsetResponse, error) {
	requestDef := GenReqDefForResetMessageOffset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetMessageOffsetResponse), nil
	}
}

func (c *KafkaClient) ResetMessageOffsetInvoker(request *model.ResetMessageOffsetRequest) *ResetMessageOffsetInvoker {
	requestDef := GenReqDefForResetMessageOffset()
	return &ResetMessageOffsetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

func (c *KafkaClient) ResetPasswordInvoker(request *model.ResetPasswordRequest) *ResetPasswordInvoker {
	requestDef := GenReqDefForResetPassword()
	return &ResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ResetUserPasswrod(request *model.ResetUserPasswrodRequest) (*model.ResetUserPasswrodResponse, error) {
	requestDef := GenReqDefForResetUserPasswrod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetUserPasswrodResponse), nil
	}
}

func (c *KafkaClient) ResetUserPasswrodInvoker(request *model.ResetUserPasswrodRequest) *ResetUserPasswrodInvoker {
	requestDef := GenReqDefForResetUserPasswrod()
	return &ResetUserPasswrodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ResizeEngineInstance(request *model.ResizeEngineInstanceRequest) (*model.ResizeEngineInstanceResponse, error) {
	requestDef := GenReqDefForResizeEngineInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeEngineInstanceResponse), nil
	}
}

func (c *KafkaClient) ResizeEngineInstanceInvoker(request *model.ResizeEngineInstanceRequest) *ResizeEngineInstanceInvoker {
	requestDef := GenReqDefForResizeEngineInstance()
	return &ResizeEngineInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

func (c *KafkaClient) ResizeInstanceInvoker(request *model.ResizeInstanceRequest) *ResizeInstanceInvoker {
	requestDef := GenReqDefForResizeInstance()
	return &ResizeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) RestartManager(request *model.RestartManagerRequest) (*model.RestartManagerResponse, error) {
	requestDef := GenReqDefForRestartManager()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartManagerResponse), nil
	}
}

func (c *KafkaClient) RestartManagerInvoker(request *model.RestartManagerRequest) *RestartManagerInvoker {
	requestDef := GenReqDefForRestartManager()
	return &RestartManagerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowBackgroundTask(request *model.ShowBackgroundTaskRequest) (*model.ShowBackgroundTaskResponse, error) {
	requestDef := GenReqDefForShowBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackgroundTaskResponse), nil
	}
}

func (c *KafkaClient) ShowBackgroundTaskInvoker(request *model.ShowBackgroundTaskRequest) *ShowBackgroundTaskInvoker {
	requestDef := GenReqDefForShowBackgroundTask()
	return &ShowBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowCesHierarchy(request *model.ShowCesHierarchyRequest) (*model.ShowCesHierarchyResponse, error) {
	requestDef := GenReqDefForShowCesHierarchy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCesHierarchyResponse), nil
	}
}

func (c *KafkaClient) ShowCesHierarchyInvoker(request *model.ShowCesHierarchyRequest) *ShowCesHierarchyInvoker {
	requestDef := GenReqDefForShowCesHierarchy()
	return &ShowCesHierarchyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowCluster(request *model.ShowClusterRequest) (*model.ShowClusterResponse, error) {
	requestDef := GenReqDefForShowCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterResponse), nil
	}
}

func (c *KafkaClient) ShowClusterInvoker(request *model.ShowClusterRequest) *ShowClusterInvoker {
	requestDef := GenReqDefForShowCluster()
	return &ShowClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowCoordinators(request *model.ShowCoordinatorsRequest) (*model.ShowCoordinatorsResponse, error) {
	requestDef := GenReqDefForShowCoordinators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCoordinatorsResponse), nil
	}
}

func (c *KafkaClient) ShowCoordinatorsInvoker(request *model.ShowCoordinatorsRequest) *ShowCoordinatorsInvoker {
	requestDef := GenReqDefForShowCoordinators()
	return &ShowCoordinatorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowEngineInstanceExtendProductInfo(request *model.ShowEngineInstanceExtendProductInfoRequest) (*model.ShowEngineInstanceExtendProductInfoResponse, error) {
	requestDef := GenReqDefForShowEngineInstanceExtendProductInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineInstanceExtendProductInfoResponse), nil
	}
}

func (c *KafkaClient) ShowEngineInstanceExtendProductInfoInvoker(request *model.ShowEngineInstanceExtendProductInfoRequest) *ShowEngineInstanceExtendProductInfoInvoker {
	requestDef := GenReqDefForShowEngineInstanceExtendProductInfo()
	return &ShowEngineInstanceExtendProductInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowGroups(request *model.ShowGroupsRequest) (*model.ShowGroupsResponse, error) {
	requestDef := GenReqDefForShowGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupsResponse), nil
	}
}

func (c *KafkaClient) ShowGroupsInvoker(request *model.ShowGroupsRequest) *ShowGroupsInvoker {
	requestDef := GenReqDefForShowGroups()
	return &ShowGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

func (c *KafkaClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowInstanceExtendProductInfo(request *model.ShowInstanceExtendProductInfoRequest) (*model.ShowInstanceExtendProductInfoResponse, error) {
	requestDef := GenReqDefForShowInstanceExtendProductInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceExtendProductInfoResponse), nil
	}
}

func (c *KafkaClient) ShowInstanceExtendProductInfoInvoker(request *model.ShowInstanceExtendProductInfoRequest) *ShowInstanceExtendProductInfoInvoker {
	requestDef := GenReqDefForShowInstanceExtendProductInfo()
	return &ShowInstanceExtendProductInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowInstanceMessages(request *model.ShowInstanceMessagesRequest) (*model.ShowInstanceMessagesResponse, error) {
	requestDef := GenReqDefForShowInstanceMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceMessagesResponse), nil
	}
}

func (c *KafkaClient) ShowInstanceMessagesInvoker(request *model.ShowInstanceMessagesRequest) *ShowInstanceMessagesInvoker {
	requestDef := GenReqDefForShowInstanceMessages()
	return &ShowInstanceMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowInstanceTopicDetail(request *model.ShowInstanceTopicDetailRequest) (*model.ShowInstanceTopicDetailResponse, error) {
	requestDef := GenReqDefForShowInstanceTopicDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceTopicDetailResponse), nil
	}
}

func (c *KafkaClient) ShowInstanceTopicDetailInvoker(request *model.ShowInstanceTopicDetailRequest) *ShowInstanceTopicDetailInvoker {
	requestDef := GenReqDefForShowInstanceTopicDetail()
	return &ShowInstanceTopicDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowInstanceUsers(request *model.ShowInstanceUsersRequest) (*model.ShowInstanceUsersResponse, error) {
	requestDef := GenReqDefForShowInstanceUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceUsersResponse), nil
	}
}

func (c *KafkaClient) ShowInstanceUsersInvoker(request *model.ShowInstanceUsersRequest) *ShowInstanceUsersInvoker {
	requestDef := GenReqDefForShowInstanceUsers()
	return &ShowInstanceUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowKafkaProjectTags(request *model.ShowKafkaProjectTagsRequest) (*model.ShowKafkaProjectTagsResponse, error) {
	requestDef := GenReqDefForShowKafkaProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKafkaProjectTagsResponse), nil
	}
}

func (c *KafkaClient) ShowKafkaProjectTagsInvoker(request *model.ShowKafkaProjectTagsRequest) *ShowKafkaProjectTagsInvoker {
	requestDef := GenReqDefForShowKafkaProjectTags()
	return &ShowKafkaProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowKafkaTags(request *model.ShowKafkaTagsRequest) (*model.ShowKafkaTagsResponse, error) {
	requestDef := GenReqDefForShowKafkaTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKafkaTagsResponse), nil
	}
}

func (c *KafkaClient) ShowKafkaTagsInvoker(request *model.ShowKafkaTagsRequest) *ShowKafkaTagsInvoker {
	requestDef := GenReqDefForShowKafkaTags()
	return &ShowKafkaTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowKafkaTopicPartitionDiskusage(request *model.ShowKafkaTopicPartitionDiskusageRequest) (*model.ShowKafkaTopicPartitionDiskusageResponse, error) {
	requestDef := GenReqDefForShowKafkaTopicPartitionDiskusage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKafkaTopicPartitionDiskusageResponse), nil
	}
}

func (c *KafkaClient) ShowKafkaTopicPartitionDiskusageInvoker(request *model.ShowKafkaTopicPartitionDiskusageRequest) *ShowKafkaTopicPartitionDiskusageInvoker {
	requestDef := GenReqDefForShowKafkaTopicPartitionDiskusage()
	return &ShowKafkaTopicPartitionDiskusageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowMaintainWindows(request *model.ShowMaintainWindowsRequest) (*model.ShowMaintainWindowsResponse, error) {
	requestDef := GenReqDefForShowMaintainWindows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMaintainWindowsResponse), nil
	}
}

func (c *KafkaClient) ShowMaintainWindowsInvoker(request *model.ShowMaintainWindowsRequest) *ShowMaintainWindowsInvoker {
	requestDef := GenReqDefForShowMaintainWindows()
	return &ShowMaintainWindowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowMessages(request *model.ShowMessagesRequest) (*model.ShowMessagesResponse, error) {
	requestDef := GenReqDefForShowMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMessagesResponse), nil
	}
}

func (c *KafkaClient) ShowMessagesInvoker(request *model.ShowMessagesRequest) *ShowMessagesInvoker {
	requestDef := GenReqDefForShowMessages()
	return &ShowMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowPartitionBeginningMessage(request *model.ShowPartitionBeginningMessageRequest) (*model.ShowPartitionBeginningMessageResponse, error) {
	requestDef := GenReqDefForShowPartitionBeginningMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartitionBeginningMessageResponse), nil
	}
}

func (c *KafkaClient) ShowPartitionBeginningMessageInvoker(request *model.ShowPartitionBeginningMessageRequest) *ShowPartitionBeginningMessageInvoker {
	requestDef := GenReqDefForShowPartitionBeginningMessage()
	return &ShowPartitionBeginningMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowPartitionEndMessage(request *model.ShowPartitionEndMessageRequest) (*model.ShowPartitionEndMessageResponse, error) {
	requestDef := GenReqDefForShowPartitionEndMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartitionEndMessageResponse), nil
	}
}

func (c *KafkaClient) ShowPartitionEndMessageInvoker(request *model.ShowPartitionEndMessageRequest) *ShowPartitionEndMessageInvoker {
	requestDef := GenReqDefForShowPartitionEndMessage()
	return &ShowPartitionEndMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowPartitionMessage(request *model.ShowPartitionMessageRequest) (*model.ShowPartitionMessageResponse, error) {
	requestDef := GenReqDefForShowPartitionMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartitionMessageResponse), nil
	}
}

func (c *KafkaClient) ShowPartitionMessageInvoker(request *model.ShowPartitionMessageRequest) *ShowPartitionMessageInvoker {
	requestDef := GenReqDefForShowPartitionMessage()
	return &ShowPartitionMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowSinkTaskDetail(request *model.ShowSinkTaskDetailRequest) (*model.ShowSinkTaskDetailResponse, error) {
	requestDef := GenReqDefForShowSinkTaskDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSinkTaskDetailResponse), nil
	}
}

func (c *KafkaClient) ShowSinkTaskDetailInvoker(request *model.ShowSinkTaskDetailRequest) *ShowSinkTaskDetailInvoker {
	requestDef := GenReqDefForShowSinkTaskDetail()
	return &ShowSinkTaskDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) ShowTopicAccessPolicy(request *model.ShowTopicAccessPolicyRequest) (*model.ShowTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForShowTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopicAccessPolicyResponse), nil
	}
}

func (c *KafkaClient) ShowTopicAccessPolicyInvoker(request *model.ShowTopicAccessPolicyRequest) *ShowTopicAccessPolicyInvoker {
	requestDef := GenReqDefForShowTopicAccessPolicy()
	return &ShowTopicAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

func (c *KafkaClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) UpdateInstanceAutoCreateTopic(request *model.UpdateInstanceAutoCreateTopicRequest) (*model.UpdateInstanceAutoCreateTopicResponse, error) {
	requestDef := GenReqDefForUpdateInstanceAutoCreateTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceAutoCreateTopicResponse), nil
	}
}

func (c *KafkaClient) UpdateInstanceAutoCreateTopicInvoker(request *model.UpdateInstanceAutoCreateTopicRequest) *UpdateInstanceAutoCreateTopicInvoker {
	requestDef := GenReqDefForUpdateInstanceAutoCreateTopic()
	return &UpdateInstanceAutoCreateTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) UpdateInstanceConsumerGroup(request *model.UpdateInstanceConsumerGroupRequest) (*model.UpdateInstanceConsumerGroupResponse, error) {
	requestDef := GenReqDefForUpdateInstanceConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceConsumerGroupResponse), nil
	}
}

func (c *KafkaClient) UpdateInstanceConsumerGroupInvoker(request *model.UpdateInstanceConsumerGroupRequest) *UpdateInstanceConsumerGroupInvoker {
	requestDef := GenReqDefForUpdateInstanceConsumerGroup()
	return &UpdateInstanceConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) UpdateInstanceCrossVpcIp(request *model.UpdateInstanceCrossVpcIpRequest) (*model.UpdateInstanceCrossVpcIpResponse, error) {
	requestDef := GenReqDefForUpdateInstanceCrossVpcIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceCrossVpcIpResponse), nil
	}
}

func (c *KafkaClient) UpdateInstanceCrossVpcIpInvoker(request *model.UpdateInstanceCrossVpcIpRequest) *UpdateInstanceCrossVpcIpInvoker {
	requestDef := GenReqDefForUpdateInstanceCrossVpcIp()
	return &UpdateInstanceCrossVpcIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) UpdateInstanceTopic(request *model.UpdateInstanceTopicRequest) (*model.UpdateInstanceTopicResponse, error) {
	requestDef := GenReqDefForUpdateInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceTopicResponse), nil
	}
}

func (c *KafkaClient) UpdateInstanceTopicInvoker(request *model.UpdateInstanceTopicRequest) *UpdateInstanceTopicInvoker {
	requestDef := GenReqDefForUpdateInstanceTopic()
	return &UpdateInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) UpdateInstanceUser(request *model.UpdateInstanceUserRequest) (*model.UpdateInstanceUserResponse, error) {
	requestDef := GenReqDefForUpdateInstanceUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceUserResponse), nil
	}
}

func (c *KafkaClient) UpdateInstanceUserInvoker(request *model.UpdateInstanceUserRequest) *UpdateInstanceUserInvoker {
	requestDef := GenReqDefForUpdateInstanceUser()
	return &UpdateInstanceUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) UpdateSinkTaskQuota(request *model.UpdateSinkTaskQuotaRequest) (*model.UpdateSinkTaskQuotaResponse, error) {
	requestDef := GenReqDefForUpdateSinkTaskQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSinkTaskQuotaResponse), nil
	}
}

func (c *KafkaClient) UpdateSinkTaskQuotaInvoker(request *model.UpdateSinkTaskQuotaRequest) *UpdateSinkTaskQuotaInvoker {
	requestDef := GenReqDefForUpdateSinkTaskQuota()
	return &UpdateSinkTaskQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) UpdateTopicAccessPolicy(request *model.UpdateTopicAccessPolicyRequest) (*model.UpdateTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForUpdateTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicAccessPolicyResponse), nil
	}
}

func (c *KafkaClient) UpdateTopicAccessPolicyInvoker(request *model.UpdateTopicAccessPolicyRequest) *UpdateTopicAccessPolicyInvoker {
	requestDef := GenReqDefForUpdateTopicAccessPolicy()
	return &UpdateTopicAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KafkaClient) UpdateTopicReplica(request *model.UpdateTopicReplicaRequest) (*model.UpdateTopicReplicaResponse, error) {
	requestDef := GenReqDefForUpdateTopicReplica()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicReplicaResponse), nil
	}
}

func (c *KafkaClient) UpdateTopicReplicaInvoker(request *model.UpdateTopicReplicaRequest) *UpdateTopicReplicaInvoker {
	requestDef := GenReqDefForUpdateTopicReplica()
	return &UpdateTopicReplicaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
