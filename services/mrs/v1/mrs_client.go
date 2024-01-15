package v1

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/mrs/v1/model"
)

type MrsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewMrsClient(hcClient *http_client.HcHttpClient) *MrsClient {
	return &MrsClient{HcClient: hcClient}
}

func MrsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *MrsClient) BatchCreateClusterTags(request *model.BatchCreateClusterTagsRequest) (*model.BatchCreateClusterTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateClusterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateClusterTagsResponse), nil
	}
}

func (c *MrsClient) BatchCreateClusterTagsInvoker(request *model.BatchCreateClusterTagsRequest) *BatchCreateClusterTagsInvoker {
	requestDef := GenReqDefForBatchCreateClusterTags()
	return &BatchCreateClusterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) BatchDeleteClusterTags(request *model.BatchDeleteClusterTagsRequest) (*model.BatchDeleteClusterTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteClusterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteClusterTagsResponse), nil
	}
}

func (c *MrsClient) BatchDeleteClusterTagsInvoker(request *model.BatchDeleteClusterTagsRequest) *BatchDeleteClusterTagsInvoker {
	requestDef := GenReqDefForBatchDeleteClusterTags()
	return &BatchDeleteClusterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

func (c *MrsClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) CreateClusterTag(request *model.CreateClusterTagRequest) (*model.CreateClusterTagResponse, error) {
	requestDef := GenReqDefForCreateClusterTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterTagResponse), nil
	}
}

func (c *MrsClient) CreateClusterTagInvoker(request *model.CreateClusterTagRequest) *CreateClusterTagInvoker {
	requestDef := GenReqDefForCreateClusterTag()
	return &CreateClusterTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) CreateScalingPolicy(request *model.CreateScalingPolicyRequest) (*model.CreateScalingPolicyResponse, error) {
	requestDef := GenReqDefForCreateScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingPolicyResponse), nil
	}
}

func (c *MrsClient) CreateScalingPolicyInvoker(request *model.CreateScalingPolicyRequest) *CreateScalingPolicyInvoker {
	requestDef := GenReqDefForCreateScalingPolicy()
	return &CreateScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

func (c *MrsClient) DeleteClusterInvoker(request *model.DeleteClusterRequest) *DeleteClusterInvoker {
	requestDef := GenReqDefForDeleteCluster()
	return &DeleteClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) DeleteClusterTag(request *model.DeleteClusterTagRequest) (*model.DeleteClusterTagResponse, error) {
	requestDef := GenReqDefForDeleteClusterTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterTagResponse), nil
	}
}

func (c *MrsClient) DeleteClusterTagInvoker(request *model.DeleteClusterTagRequest) *DeleteClusterTagInvoker {
	requestDef := GenReqDefForDeleteClusterTag()
	return &DeleteClusterTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ListAllTags(request *model.ListAllTagsRequest) (*model.ListAllTagsResponse, error) {
	requestDef := GenReqDefForListAllTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllTagsResponse), nil
	}
}

func (c *MrsClient) ListAllTagsInvoker(request *model.ListAllTagsRequest) *ListAllTagsInvoker {
	requestDef := GenReqDefForListAllTags()
	return &ListAllTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ListClusterTags(request *model.ListClusterTagsRequest) (*model.ListClusterTagsResponse, error) {
	requestDef := GenReqDefForListClusterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterTagsResponse), nil
	}
}

func (c *MrsClient) ListClusterTagsInvoker(request *model.ListClusterTagsRequest) *ListClusterTagsInvoker {
	requestDef := GenReqDefForListClusterTags()
	return &ListClusterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ListClusters(request *model.ListClustersRequest) (*model.ListClustersResponse, error) {
	requestDef := GenReqDefForListClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersResponse), nil
	}
}

func (c *MrsClient) ListClustersInvoker(request *model.ListClustersRequest) *ListClustersInvoker {
	requestDef := GenReqDefForListClusters()
	return &ListClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ListClustersByTags(request *model.ListClustersByTagsRequest) (*model.ListClustersByTagsResponse, error) {
	requestDef := GenReqDefForListClustersByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersByTagsResponse), nil
	}
}

func (c *MrsClient) ListClustersByTagsInvoker(request *model.ListClustersByTagsRequest) *ListClustersByTagsInvoker {
	requestDef := GenReqDefForListClustersByTags()
	return &ListClustersByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ListHosts(request *model.ListHostsRequest) (*model.ListHostsResponse, error) {
	requestDef := GenReqDefForListHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostsResponse), nil
	}
}

func (c *MrsClient) ListHostsInvoker(request *model.ListHostsRequest) *ListHostsInvoker {
	requestDef := GenReqDefForListHosts()
	return &ListHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ShowClusterDetails(request *model.ShowClusterDetailsRequest) (*model.ShowClusterDetailsResponse, error) {
	requestDef := GenReqDefForShowClusterDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterDetailsResponse), nil
	}
}

func (c *MrsClient) ShowClusterDetailsInvoker(request *model.ShowClusterDetailsRequest) *ShowClusterDetailsInvoker {
	requestDef := GenReqDefForShowClusterDetails()
	return &ShowClusterDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) UpdateClusterScaling(request *model.UpdateClusterScalingRequest) (*model.UpdateClusterScalingResponse, error) {
	requestDef := GenReqDefForUpdateClusterScaling()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterScalingResponse), nil
	}
}

func (c *MrsClient) UpdateClusterScalingInvoker(request *model.UpdateClusterScalingRequest) *UpdateClusterScalingInvoker {
	requestDef := GenReqDefForUpdateClusterScaling()
	return &UpdateClusterScalingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
