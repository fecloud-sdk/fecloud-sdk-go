package v1

import "github.com/fecloud-sdk/fecloud-sdk-go/services/mrs/v1/model"

func (c *MrsClient) GetBatchCreateClusterTagsRequest() *model.BatchCreateClusterTagsRequest {
	return new(model.BatchCreateClusterTagsRequest)
}

func (c *MrsClient) GetBatchDeleteClusterTagsRequest() *model.BatchDeleteClusterTagsRequest {
	return new(model.BatchDeleteClusterTagsRequest)
}

func (c *MrsClient) GetCreateClusterRequest() *model.CreateClusterRequest {
	return new(model.CreateClusterRequest)
}

func (c *MrsClient) GetCreateClusterTagRequest() *model.CreateClusterTagRequest {
	return new(model.CreateClusterTagRequest)
}

func (c *MrsClient) GetCreateScalingPolicyRequest() *model.CreateScalingPolicyRequest {
	return new(model.CreateScalingPolicyRequest)
}

func (c *MrsClient) GetDeleteClusterRequest() *model.DeleteClusterRequest {
	return new(model.DeleteClusterRequest)
}

func (c *MrsClient) GetDeleteClusterTagRequest() *model.DeleteClusterTagRequest {
	return new(model.DeleteClusterTagRequest)
}

func (c *MrsClient) GetListAllTagsRequest() *model.ListAllTagsRequest {
	return new(model.ListAllTagsRequest)
}

func (c *MrsClient) GetListClusterTagsRequest() *model.ListClusterTagsRequest {
	return new(model.ListClusterTagsRequest)
}

func (c *MrsClient) GetListClustersRequest() *model.ListClustersRequest {
	return new(model.ListClustersRequest)
}

func (c *MrsClient) GetListClustersByTagsRequest() *model.ListClustersByTagsRequest {
	return new(model.ListClustersByTagsRequest)
}

func (c *MrsClient) GetListHostsRequest() *model.ListHostsRequest {
	return new(model.ListHostsRequest)
}

func (c *MrsClient) GetShowClusterDetailsRequest() *model.ShowClusterDetailsRequest {
	return new(model.ShowClusterDetailsRequest)
}

func (c *MrsClient) GetUpdateClusterScalingRequest() *model.UpdateClusterScalingRequest {
	return new(model.UpdateClusterScalingRequest)
}
