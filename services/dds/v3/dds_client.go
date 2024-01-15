package v3

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/dds/v3/model"
)

type DdsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDdsClient(hcClient *http_client.HcHttpClient) *DdsClient {
	return &DdsClient{HcClient: hcClient}
}

func DdsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *DdsClient) AddReadonlyNode(request *model.AddReadonlyNodeRequest) (*model.AddReadonlyNodeResponse, error) {
	requestDef := GenReqDefForAddReadonlyNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddReadonlyNodeResponse), nil
	}
}

func (c *DdsClient) AddReadonlyNodeInvoker(request *model.AddReadonlyNodeRequest) *AddReadonlyNodeInvoker {
	requestDef := GenReqDefForAddReadonlyNode()
	return &AddReadonlyNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) AddShardingNode(request *model.AddShardingNodeRequest) (*model.AddShardingNodeResponse, error) {
	requestDef := GenReqDefForAddShardingNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddShardingNodeResponse), nil
	}
}

func (c *DdsClient) AddShardingNodeInvoker(request *model.AddShardingNodeRequest) *AddShardingNodeInvoker {
	requestDef := GenReqDefForAddShardingNode()
	return &AddShardingNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) AttachEip(request *model.AttachEipRequest) (*model.AttachEipResponse, error) {
	requestDef := GenReqDefForAttachEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachEipResponse), nil
	}
}

func (c *DdsClient) AttachEipInvoker(request *model.AttachEipRequest) *AttachEipInvoker {
	requestDef := GenReqDefForAttachEip()
	return &AttachEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) AttachInternalIp(request *model.AttachInternalIpRequest) (*model.AttachInternalIpResponse, error) {
	requestDef := GenReqDefForAttachInternalIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachInternalIpResponse), nil
	}
}

func (c *DdsClient) AttachInternalIpInvoker(request *model.AttachInternalIpRequest) *AttachInternalIpInvoker {
	requestDef := GenReqDefForAttachInternalIp()
	return &AttachInternalIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) BatchTagAction(request *model.BatchTagActionRequest) (*model.BatchTagActionResponse, error) {
	requestDef := GenReqDefForBatchTagAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchTagActionResponse), nil
	}
}

func (c *DdsClient) BatchTagActionInvoker(request *model.BatchTagActionRequest) *BatchTagActionInvoker {
	requestDef := GenReqDefForBatchTagAction()
	return &BatchTagActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) CancelEip(request *model.CancelEipRequest) (*model.CancelEipResponse, error) {
	requestDef := GenReqDefForCancelEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelEipResponse), nil
	}
}

func (c *DdsClient) CancelEipInvoker(request *model.CancelEipRequest) *CancelEipInvoker {
	requestDef := GenReqDefForCancelEip()
	return &CancelEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ChangeOpsWindow(request *model.ChangeOpsWindowRequest) (*model.ChangeOpsWindowResponse, error) {
	requestDef := GenReqDefForChangeOpsWindow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeOpsWindowResponse), nil
	}
}

func (c *DdsClient) ChangeOpsWindowInvoker(request *model.ChangeOpsWindowRequest) *ChangeOpsWindowInvoker {
	requestDef := GenReqDefForChangeOpsWindow()
	return &ChangeOpsWindowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) CheckPassword(request *model.CheckPasswordRequest) (*model.CheckPasswordResponse, error) {
	requestDef := GenReqDefForCheckPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckPasswordResponse), nil
	}
}

func (c *DdsClient) CheckPasswordInvoker(request *model.CheckPasswordRequest) *CheckPasswordInvoker {
	requestDef := GenReqDefForCheckPassword()
	return &CheckPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) CheckWeakPassword(request *model.CheckWeakPasswordRequest) (*model.CheckWeakPasswordResponse, error) {
	requestDef := GenReqDefForCheckWeakPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckWeakPasswordResponse), nil
	}
}

func (c *DdsClient) CheckWeakPasswordInvoker(request *model.CheckWeakPasswordRequest) *CheckWeakPasswordInvoker {
	requestDef := GenReqDefForCheckWeakPassword()
	return &CheckWeakPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) CompareConfiguration(request *model.CompareConfigurationRequest) (*model.CompareConfigurationResponse, error) {
	requestDef := GenReqDefForCompareConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareConfigurationResponse), nil
	}
}

func (c *DdsClient) CompareConfigurationInvoker(request *model.CompareConfigurationRequest) *CompareConfigurationInvoker {
	requestDef := GenReqDefForCompareConfiguration()
	return &CompareConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) CopyConfiguration(request *model.CopyConfigurationRequest) (*model.CopyConfigurationResponse, error) {
	requestDef := GenReqDefForCopyConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyConfigurationResponse), nil
	}
}

func (c *DdsClient) CopyConfigurationInvoker(request *model.CopyConfigurationRequest) *CopyConfigurationInvoker {
	requestDef := GenReqDefForCopyConfiguration()
	return &CopyConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) CreateConfiguration(request *model.CreateConfigurationRequest) (*model.CreateConfigurationResponse, error) {
	requestDef := GenReqDefForCreateConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigurationResponse), nil
	}
}

func (c *DdsClient) CreateConfigurationInvoker(request *model.CreateConfigurationRequest) *CreateConfigurationInvoker {
	requestDef := GenReqDefForCreateConfiguration()
	return &CreateConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) CreateDatabaseRole(request *model.CreateDatabaseRoleRequest) (*model.CreateDatabaseRoleResponse, error) {
	requestDef := GenReqDefForCreateDatabaseRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseRoleResponse), nil
	}
}

func (c *DdsClient) CreateDatabaseRoleInvoker(request *model.CreateDatabaseRoleRequest) *CreateDatabaseRoleInvoker {
	requestDef := GenReqDefForCreateDatabaseRole()
	return &CreateDatabaseRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) CreateDatabaseUser(request *model.CreateDatabaseUserRequest) (*model.CreateDatabaseUserResponse, error) {
	requestDef := GenReqDefForCreateDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseUserResponse), nil
	}
}

func (c *DdsClient) CreateDatabaseUserInvoker(request *model.CreateDatabaseUserRequest) *CreateDatabaseUserInvoker {
	requestDef := GenReqDefForCreateDatabaseUser()
	return &CreateDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

func (c *DdsClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) CreateIp(request *model.CreateIpRequest) (*model.CreateIpResponse, error) {
	requestDef := GenReqDefForCreateIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIpResponse), nil
	}
}

func (c *DdsClient) CreateIpInvoker(request *model.CreateIpRequest) *CreateIpInvoker {
	requestDef := GenReqDefForCreateIp()
	return &CreateIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) CreateManualBackup(request *model.CreateManualBackupRequest) (*model.CreateManualBackupResponse, error) {
	requestDef := GenReqDefForCreateManualBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateManualBackupResponse), nil
	}
}

func (c *DdsClient) CreateManualBackupInvoker(request *model.CreateManualBackupRequest) *CreateManualBackupInvoker {
	requestDef := GenReqDefForCreateManualBackup()
	return &CreateManualBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) DeleteAuditLog(request *model.DeleteAuditLogRequest) (*model.DeleteAuditLogResponse, error) {
	requestDef := GenReqDefForDeleteAuditLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditLogResponse), nil
	}
}

func (c *DdsClient) DeleteAuditLogInvoker(request *model.DeleteAuditLogRequest) *DeleteAuditLogInvoker {
	requestDef := GenReqDefForDeleteAuditLog()
	return &DeleteAuditLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) DeleteConfiguration(request *model.DeleteConfigurationRequest) (*model.DeleteConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigurationResponse), nil
	}
}

func (c *DdsClient) DeleteConfigurationInvoker(request *model.DeleteConfigurationRequest) *DeleteConfigurationInvoker {
	requestDef := GenReqDefForDeleteConfiguration()
	return &DeleteConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) DeleteDatabaseRole(request *model.DeleteDatabaseRoleRequest) (*model.DeleteDatabaseRoleResponse, error) {
	requestDef := GenReqDefForDeleteDatabaseRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseRoleResponse), nil
	}
}

func (c *DdsClient) DeleteDatabaseRoleInvoker(request *model.DeleteDatabaseRoleRequest) *DeleteDatabaseRoleInvoker {
	requestDef := GenReqDefForDeleteDatabaseRole()
	return &DeleteDatabaseRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) DeleteDatabaseUser(request *model.DeleteDatabaseUserRequest) (*model.DeleteDatabaseUserResponse, error) {
	requestDef := GenReqDefForDeleteDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseUserResponse), nil
	}
}

func (c *DdsClient) DeleteDatabaseUserInvoker(request *model.DeleteDatabaseUserRequest) *DeleteDatabaseUserInvoker {
	requestDef := GenReqDefForDeleteDatabaseUser()
	return &DeleteDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

func (c *DdsClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) DeleteManualBackup(request *model.DeleteManualBackupRequest) (*model.DeleteManualBackupResponse, error) {
	requestDef := GenReqDefForDeleteManualBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteManualBackupResponse), nil
	}
}

func (c *DdsClient) DeleteManualBackupInvoker(request *model.DeleteManualBackupRequest) *DeleteManualBackupInvoker {
	requestDef := GenReqDefForDeleteManualBackup()
	return &DeleteManualBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) DeleteSession(request *model.DeleteSessionRequest) (*model.DeleteSessionResponse, error) {
	requestDef := GenReqDefForDeleteSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSessionResponse), nil
	}
}

func (c *DdsClient) DeleteSessionInvoker(request *model.DeleteSessionRequest) *DeleteSessionInvoker {
	requestDef := GenReqDefForDeleteSession()
	return &DeleteSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) DownloadErrorlog(request *model.DownloadErrorlogRequest) (*model.DownloadErrorlogResponse, error) {
	requestDef := GenReqDefForDownloadErrorlog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadErrorlogResponse), nil
	}
}

func (c *DdsClient) DownloadErrorlogInvoker(request *model.DownloadErrorlogRequest) *DownloadErrorlogInvoker {
	requestDef := GenReqDefForDownloadErrorlog()
	return &DownloadErrorlogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) DownloadSlowlog(request *model.DownloadSlowlogRequest) (*model.DownloadSlowlogResponse, error) {
	requestDef := GenReqDefForDownloadSlowlog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadSlowlogResponse), nil
	}
}

func (c *DdsClient) DownloadSlowlogInvoker(request *model.DownloadSlowlogRequest) *DownloadSlowlogInvoker {
	requestDef := GenReqDefForDownloadSlowlog()
	return &DownloadSlowlogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ExpandReplicasetNode(request *model.ExpandReplicasetNodeRequest) (*model.ExpandReplicasetNodeResponse, error) {
	requestDef := GenReqDefForExpandReplicasetNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandReplicasetNodeResponse), nil
	}
}

func (c *DdsClient) ExpandReplicasetNodeInvoker(request *model.ExpandReplicasetNodeRequest) *ExpandReplicasetNodeInvoker {
	requestDef := GenReqDefForExpandReplicasetNode()
	return &ExpandReplicasetNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListAppliedInstances(request *model.ListAppliedInstancesRequest) (*model.ListAppliedInstancesResponse, error) {
	requestDef := GenReqDefForListAppliedInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppliedInstancesResponse), nil
	}
}

func (c *DdsClient) ListAppliedInstancesInvoker(request *model.ListAppliedInstancesRequest) *ListAppliedInstancesInvoker {
	requestDef := GenReqDefForListAppliedInstances()
	return &ListAppliedInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListAuditlogLinks(request *model.ListAuditlogLinksRequest) (*model.ListAuditlogLinksResponse, error) {
	requestDef := GenReqDefForListAuditlogLinks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditlogLinksResponse), nil
	}
}

func (c *DdsClient) ListAuditlogLinksInvoker(request *model.ListAuditlogLinksRequest) *ListAuditlogLinksInvoker {
	requestDef := GenReqDefForListAuditlogLinks()
	return &ListAuditlogLinksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListAuditlogs(request *model.ListAuditlogsRequest) (*model.ListAuditlogsResponse, error) {
	requestDef := GenReqDefForListAuditlogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditlogsResponse), nil
	}
}

func (c *DdsClient) ListAuditlogsInvoker(request *model.ListAuditlogsRequest) *ListAuditlogsInvoker {
	requestDef := GenReqDefForListAuditlogs()
	return &ListAuditlogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListAz2Migrate(request *model.ListAz2MigrateRequest) (*model.ListAz2MigrateResponse, error) {
	requestDef := GenReqDefForListAz2Migrate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAz2MigrateResponse), nil
	}
}

func (c *DdsClient) ListAz2MigrateInvoker(request *model.ListAz2MigrateRequest) *ListAz2MigrateInvoker {
	requestDef := GenReqDefForListAz2Migrate()
	return &ListAz2MigrateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListBackups(request *model.ListBackupsRequest) (*model.ListBackupsResponse, error) {
	requestDef := GenReqDefForListBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupsResponse), nil
	}
}

func (c *DdsClient) ListBackupsInvoker(request *model.ListBackupsRequest) *ListBackupsInvoker {
	requestDef := GenReqDefForListBackups()
	return &ListBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListConfigurations(request *model.ListConfigurationsRequest) (*model.ListConfigurationsResponse, error) {
	requestDef := GenReqDefForListConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsResponse), nil
	}
}

func (c *DdsClient) ListConfigurationsInvoker(request *model.ListConfigurationsRequest) *ListConfigurationsInvoker {
	requestDef := GenReqDefForListConfigurations()
	return &ListConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListDatabaseRoles(request *model.ListDatabaseRolesRequest) (*model.ListDatabaseRolesResponse, error) {
	requestDef := GenReqDefForListDatabaseRoles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseRolesResponse), nil
	}
}

func (c *DdsClient) ListDatabaseRolesInvoker(request *model.ListDatabaseRolesRequest) *ListDatabaseRolesInvoker {
	requestDef := GenReqDefForListDatabaseRoles()
	return &ListDatabaseRolesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListDatabaseUsers(request *model.ListDatabaseUsersRequest) (*model.ListDatabaseUsersResponse, error) {
	requestDef := GenReqDefForListDatabaseUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseUsersResponse), nil
	}
}

func (c *DdsClient) ListDatabaseUsersInvoker(request *model.ListDatabaseUsersRequest) *ListDatabaseUsersInvoker {
	requestDef := GenReqDefForListDatabaseUsers()
	return &ListDatabaseUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListDatastoreVersions(request *model.ListDatastoreVersionsRequest) (*model.ListDatastoreVersionsResponse, error) {
	requestDef := GenReqDefForListDatastoreVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatastoreVersionsResponse), nil
	}
}

func (c *DdsClient) ListDatastoreVersionsInvoker(request *model.ListDatastoreVersionsRequest) *ListDatastoreVersionsInvoker {
	requestDef := GenReqDefForListDatastoreVersions()
	return &ListDatastoreVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListErrorLogs(request *model.ListErrorLogsRequest) (*model.ListErrorLogsResponse, error) {
	requestDef := GenReqDefForListErrorLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListErrorLogsResponse), nil
	}
}

func (c *DdsClient) ListErrorLogsInvoker(request *model.ListErrorLogsRequest) *ListErrorLogsInvoker {
	requestDef := GenReqDefForListErrorLogs()
	return &ListErrorLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListFlavorInfos(request *model.ListFlavorInfosRequest) (*model.ListFlavorInfosResponse, error) {
	requestDef := GenReqDefForListFlavorInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorInfosResponse), nil
	}
}

func (c *DdsClient) ListFlavorInfosInvoker(request *model.ListFlavorInfosRequest) *ListFlavorInfosInvoker {
	requestDef := GenReqDefForListFlavorInfos()
	return &ListFlavorInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

func (c *DdsClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListInstanceTags(request *model.ListInstanceTagsRequest) (*model.ListInstanceTagsResponse, error) {
	requestDef := GenReqDefForListInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTagsResponse), nil
	}
}

func (c *DdsClient) ListInstanceTagsInvoker(request *model.ListInstanceTagsRequest) *ListInstanceTagsInvoker {
	requestDef := GenReqDefForListInstanceTags()
	return &ListInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

func (c *DdsClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListInstancesByTags(request *model.ListInstancesByTagsRequest) (*model.ListInstancesByTagsResponse, error) {
	requestDef := GenReqDefForListInstancesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesByTagsResponse), nil
	}
}

func (c *DdsClient) ListInstancesByTagsInvoker(request *model.ListInstancesByTagsRequest) *ListInstancesByTagsInvoker {
	requestDef := GenReqDefForListInstancesByTags()
	return &ListInstancesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListLtsSlowLogs(request *model.ListLtsSlowLogsRequest) (*model.ListLtsSlowLogsResponse, error) {
	requestDef := GenReqDefForListLtsSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLtsSlowLogsResponse), nil
	}
}

func (c *DdsClient) ListLtsSlowLogsInvoker(request *model.ListLtsSlowLogsRequest) *ListLtsSlowLogsInvoker {
	requestDef := GenReqDefForListLtsSlowLogs()
	return &ListLtsSlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

func (c *DdsClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListRecycleInstances(request *model.ListRecycleInstancesRequest) (*model.ListRecycleInstancesResponse, error) {
	requestDef := GenReqDefForListRecycleInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecycleInstancesResponse), nil
	}
}

func (c *DdsClient) ListRecycleInstancesInvoker(request *model.ListRecycleInstancesRequest) *ListRecycleInstancesInvoker {
	requestDef := GenReqDefForListRecycleInstances()
	return &ListRecycleInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListRestoreCollections(request *model.ListRestoreCollectionsRequest) (*model.ListRestoreCollectionsResponse, error) {
	requestDef := GenReqDefForListRestoreCollections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreCollectionsResponse), nil
	}
}

func (c *DdsClient) ListRestoreCollectionsInvoker(request *model.ListRestoreCollectionsRequest) *ListRestoreCollectionsInvoker {
	requestDef := GenReqDefForListRestoreCollections()
	return &ListRestoreCollectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListRestoreDatabases(request *model.ListRestoreDatabasesRequest) (*model.ListRestoreDatabasesResponse, error) {
	requestDef := GenReqDefForListRestoreDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreDatabasesResponse), nil
	}
}

func (c *DdsClient) ListRestoreDatabasesInvoker(request *model.ListRestoreDatabasesRequest) *ListRestoreDatabasesInvoker {
	requestDef := GenReqDefForListRestoreDatabases()
	return &ListRestoreDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListRestoreTimes(request *model.ListRestoreTimesRequest) (*model.ListRestoreTimesResponse, error) {
	requestDef := GenReqDefForListRestoreTimes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreTimesResponse), nil
	}
}

func (c *DdsClient) ListRestoreTimesInvoker(request *model.ListRestoreTimesRequest) *ListRestoreTimesInvoker {
	requestDef := GenReqDefForListRestoreTimes()
	return &ListRestoreTimesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListSessions(request *model.ListSessionsRequest) (*model.ListSessionsResponse, error) {
	requestDef := GenReqDefForListSessions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSessionsResponse), nil
	}
}

func (c *DdsClient) ListSessionsInvoker(request *model.ListSessionsRequest) *ListSessionsInvoker {
	requestDef := GenReqDefForListSessions()
	return &ListSessionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListSlowLogs(request *model.ListSlowLogsRequest) (*model.ListSlowLogsResponse, error) {
	requestDef := GenReqDefForListSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowLogsResponse), nil
	}
}

func (c *DdsClient) ListSlowLogsInvoker(request *model.ListSlowLogsRequest) *ListSlowLogsInvoker {
	requestDef := GenReqDefForListSlowLogs()
	return &ListSlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListSslCertDownloadAddress(request *model.ListSslCertDownloadAddressRequest) (*model.ListSslCertDownloadAddressResponse, error) {
	requestDef := GenReqDefForListSslCertDownloadAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSslCertDownloadAddressResponse), nil
	}
}

func (c *DdsClient) ListSslCertDownloadAddressInvoker(request *model.ListSslCertDownloadAddressRequest) *ListSslCertDownloadAddressInvoker {
	requestDef := GenReqDefForListSslCertDownloadAddress()
	return &ListSslCertDownloadAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListStorageType(request *model.ListStorageTypeRequest) (*model.ListStorageTypeResponse, error) {
	requestDef := GenReqDefForListStorageType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStorageTypeResponse), nil
	}
}

func (c *DdsClient) ListStorageTypeInvoker(request *model.ListStorageTypeRequest) *ListStorageTypeInvoker {
	requestDef := GenReqDefForListStorageType()
	return &ListStorageTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

func (c *DdsClient) ListTasksInvoker(request *model.ListTasksRequest) *ListTasksInvoker {
	requestDef := GenReqDefForListTasks()
	return &ListTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) MigrateAz(request *model.MigrateAzRequest) (*model.MigrateAzResponse, error) {
	requestDef := GenReqDefForMigrateAz()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateAzResponse), nil
	}
}

func (c *DdsClient) MigrateAzInvoker(request *model.MigrateAzRequest) *MigrateAzInvoker {
	requestDef := GenReqDefForMigrateAz()
	return &MigrateAzInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ResetConfiguration(request *model.ResetConfigurationRequest) (*model.ResetConfigurationResponse, error) {
	requestDef := GenReqDefForResetConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetConfigurationResponse), nil
	}
}

func (c *DdsClient) ResetConfigurationInvoker(request *model.ResetConfigurationRequest) *ResetConfigurationInvoker {
	requestDef := GenReqDefForResetConfiguration()
	return &ResetConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

func (c *DdsClient) ResetPasswordInvoker(request *model.ResetPasswordRequest) *ResetPasswordInvoker {
	requestDef := GenReqDefForResetPassword()
	return &ResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

func (c *DdsClient) ResizeInstanceInvoker(request *model.ResizeInstanceRequest) *ResizeInstanceInvoker {
	requestDef := GenReqDefForResizeInstance()
	return &ResizeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ResizeInstanceVolume(request *model.ResizeInstanceVolumeRequest) (*model.ResizeInstanceVolumeResponse, error) {
	requestDef := GenReqDefForResizeInstanceVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceVolumeResponse), nil
	}
}

func (c *DdsClient) ResizeInstanceVolumeInvoker(request *model.ResizeInstanceVolumeRequest) *ResizeInstanceVolumeInvoker {
	requestDef := GenReqDefForResizeInstanceVolume()
	return &ResizeInstanceVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) RestartInstance(request *model.RestartInstanceRequest) (*model.RestartInstanceResponse, error) {
	requestDef := GenReqDefForRestartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartInstanceResponse), nil
	}
}

func (c *DdsClient) RestartInstanceInvoker(request *model.RestartInstanceRequest) *RestartInstanceInvoker {
	requestDef := GenReqDefForRestartInstance()
	return &RestartInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) RestoreInstance(request *model.RestoreInstanceRequest) (*model.RestoreInstanceResponse, error) {
	requestDef := GenReqDefForRestoreInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreInstanceResponse), nil
	}
}

func (c *DdsClient) RestoreInstanceInvoker(request *model.RestoreInstanceRequest) *RestoreInstanceInvoker {
	requestDef := GenReqDefForRestoreInstance()
	return &RestoreInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) RestoreInstanceFromCollection(request *model.RestoreInstanceFromCollectionRequest) (*model.RestoreInstanceFromCollectionResponse, error) {
	requestDef := GenReqDefForRestoreInstanceFromCollection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreInstanceFromCollectionResponse), nil
	}
}

func (c *DdsClient) RestoreInstanceFromCollectionInvoker(request *model.RestoreInstanceFromCollectionRequest) *RestoreInstanceFromCollectionInvoker {
	requestDef := GenReqDefForRestoreInstanceFromCollection()
	return &RestoreInstanceFromCollectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) RestoreNewInstance(request *model.RestoreNewInstanceRequest) (*model.RestoreNewInstanceResponse, error) {
	requestDef := GenReqDefForRestoreNewInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreNewInstanceResponse), nil
	}
}

func (c *DdsClient) RestoreNewInstanceInvoker(request *model.RestoreNewInstanceRequest) *RestoreNewInstanceInvoker {
	requestDef := GenReqDefForRestoreNewInstance()
	return &RestoreNewInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) SetAuditlogPolicy(request *model.SetAuditlogPolicyRequest) (*model.SetAuditlogPolicyResponse, error) {
	requestDef := GenReqDefForSetAuditlogPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAuditlogPolicyResponse), nil
	}
}

func (c *DdsClient) SetAuditlogPolicyInvoker(request *model.SetAuditlogPolicyRequest) *SetAuditlogPolicyInvoker {
	requestDef := GenReqDefForSetAuditlogPolicy()
	return &SetAuditlogPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) SetBackupPolicy(request *model.SetBackupPolicyRequest) (*model.SetBackupPolicyResponse, error) {
	requestDef := GenReqDefForSetBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBackupPolicyResponse), nil
	}
}

func (c *DdsClient) SetBackupPolicyInvoker(request *model.SetBackupPolicyRequest) *SetBackupPolicyInvoker {
	requestDef := GenReqDefForSetBackupPolicy()
	return &SetBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) SetBalancerSwitch(request *model.SetBalancerSwitchRequest) (*model.SetBalancerSwitchResponse, error) {
	requestDef := GenReqDefForSetBalancerSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBalancerSwitchResponse), nil
	}
}

func (c *DdsClient) SetBalancerSwitchInvoker(request *model.SetBalancerSwitchRequest) *SetBalancerSwitchInvoker {
	requestDef := GenReqDefForSetBalancerSwitch()
	return &SetBalancerSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) SetBalancerWindow(request *model.SetBalancerWindowRequest) (*model.SetBalancerWindowResponse, error) {
	requestDef := GenReqDefForSetBalancerWindow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBalancerWindowResponse), nil
	}
}

func (c *DdsClient) SetBalancerWindowInvoker(request *model.SetBalancerWindowRequest) *SetBalancerWindowInvoker {
	requestDef := GenReqDefForSetBalancerWindow()
	return &SetBalancerWindowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) SetRecyclePolicy(request *model.SetRecyclePolicyRequest) (*model.SetRecyclePolicyResponse, error) {
	requestDef := GenReqDefForSetRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRecyclePolicyResponse), nil
	}
}

func (c *DdsClient) SetRecyclePolicyInvoker(request *model.SetRecyclePolicyRequest) *SetRecyclePolicyInvoker {
	requestDef := GenReqDefForSetRecyclePolicy()
	return &SetRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowAuditlogPolicy(request *model.ShowAuditlogPolicyRequest) (*model.ShowAuditlogPolicyResponse, error) {
	requestDef := GenReqDefForShowAuditlogPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditlogPolicyResponse), nil
	}
}

func (c *DdsClient) ShowAuditlogPolicyInvoker(request *model.ShowAuditlogPolicyRequest) *ShowAuditlogPolicyInvoker {
	requestDef := GenReqDefForShowAuditlogPolicy()
	return &ShowAuditlogPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowBackupDownloadLink(request *model.ShowBackupDownloadLinkRequest) (*model.ShowBackupDownloadLinkResponse, error) {
	requestDef := GenReqDefForShowBackupDownloadLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupDownloadLinkResponse), nil
	}
}

func (c *DdsClient) ShowBackupDownloadLinkInvoker(request *model.ShowBackupDownloadLinkRequest) *ShowBackupDownloadLinkInvoker {
	requestDef := GenReqDefForShowBackupDownloadLink()
	return &ShowBackupDownloadLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowBackupPolicy(request *model.ShowBackupPolicyRequest) (*model.ShowBackupPolicyResponse, error) {
	requestDef := GenReqDefForShowBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupPolicyResponse), nil
	}
}

func (c *DdsClient) ShowBackupPolicyInvoker(request *model.ShowBackupPolicyRequest) *ShowBackupPolicyInvoker {
	requestDef := GenReqDefForShowBackupPolicy()
	return &ShowBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowConfigurationAppliedHistory(request *model.ShowConfigurationAppliedHistoryRequest) (*model.ShowConfigurationAppliedHistoryResponse, error) {
	requestDef := GenReqDefForShowConfigurationAppliedHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationAppliedHistoryResponse), nil
	}
}

func (c *DdsClient) ShowConfigurationAppliedHistoryInvoker(request *model.ShowConfigurationAppliedHistoryRequest) *ShowConfigurationAppliedHistoryInvoker {
	requestDef := GenReqDefForShowConfigurationAppliedHistory()
	return &ShowConfigurationAppliedHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowConfigurationModifyHistory(request *model.ShowConfigurationModifyHistoryRequest) (*model.ShowConfigurationModifyHistoryResponse, error) {
	requestDef := GenReqDefForShowConfigurationModifyHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationModifyHistoryResponse), nil
	}
}

func (c *DdsClient) ShowConfigurationModifyHistoryInvoker(request *model.ShowConfigurationModifyHistoryRequest) *ShowConfigurationModifyHistoryInvoker {
	requestDef := GenReqDefForShowConfigurationModifyHistory()
	return &ShowConfigurationModifyHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowConfigurationParameter(request *model.ShowConfigurationParameterRequest) (*model.ShowConfigurationParameterResponse, error) {
	requestDef := GenReqDefForShowConfigurationParameter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationParameterResponse), nil
	}
}

func (c *DdsClient) ShowConfigurationParameterInvoker(request *model.ShowConfigurationParameterRequest) *ShowConfigurationParameterInvoker {
	requestDef := GenReqDefForShowConfigurationParameter()
	return &ShowConfigurationParameterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowConnectionStatistics(request *model.ShowConnectionStatisticsRequest) (*model.ShowConnectionStatisticsResponse, error) {
	requestDef := GenReqDefForShowConnectionStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConnectionStatisticsResponse), nil
	}
}

func (c *DdsClient) ShowConnectionStatisticsInvoker(request *model.ShowConnectionStatisticsRequest) *ShowConnectionStatisticsInvoker {
	requestDef := GenReqDefForShowConnectionStatistics()
	return &ShowConnectionStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowDiskUsage(request *model.ShowDiskUsageRequest) (*model.ShowDiskUsageResponse, error) {
	requestDef := GenReqDefForShowDiskUsage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiskUsageResponse), nil
	}
}

func (c *DdsClient) ShowDiskUsageInvoker(request *model.ShowDiskUsageRequest) *ShowDiskUsageInvoker {
	requestDef := GenReqDefForShowDiskUsage()
	return &ShowDiskUsageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowEntityConfiguration(request *model.ShowEntityConfigurationRequest) (*model.ShowEntityConfigurationResponse, error) {
	requestDef := GenReqDefForShowEntityConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEntityConfigurationResponse), nil
	}
}

func (c *DdsClient) ShowEntityConfigurationInvoker(request *model.ShowEntityConfigurationRequest) *ShowEntityConfigurationInvoker {
	requestDef := GenReqDefForShowEntityConfiguration()
	return &ShowEntityConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobDetailResponse), nil
	}
}

func (c *DdsClient) ShowJobDetailInvoker(request *model.ShowJobDetailRequest) *ShowJobDetailInvoker {
	requestDef := GenReqDefForShowJobDetail()
	return &ShowJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

func (c *DdsClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowRecyclePolicy(request *model.ShowRecyclePolicyRequest) (*model.ShowRecyclePolicyResponse, error) {
	requestDef := GenReqDefForShowRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecyclePolicyResponse), nil
	}
}

func (c *DdsClient) ShowRecyclePolicyInvoker(request *model.ShowRecyclePolicyRequest) *ShowRecyclePolicyInvoker {
	requestDef := GenReqDefForShowRecyclePolicy()
	return &ShowRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowReplSetName(request *model.ShowReplSetNameRequest) (*model.ShowReplSetNameResponse, error) {
	requestDef := GenReqDefForShowReplSetName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReplSetNameResponse), nil
	}
}

func (c *DdsClient) ShowReplSetNameInvoker(request *model.ShowReplSetNameRequest) *ShowReplSetNameInvoker {
	requestDef := GenReqDefForShowReplSetName()
	return &ShowReplSetNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowSecondLevelMonitoringStatus(request *model.ShowSecondLevelMonitoringStatusRequest) (*model.ShowSecondLevelMonitoringStatusResponse, error) {
	requestDef := GenReqDefForShowSecondLevelMonitoringStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecondLevelMonitoringStatusResponse), nil
	}
}

func (c *DdsClient) ShowSecondLevelMonitoringStatusInvoker(request *model.ShowSecondLevelMonitoringStatusRequest) *ShowSecondLevelMonitoringStatusInvoker {
	requestDef := GenReqDefForShowSecondLevelMonitoringStatus()
	return &ShowSecondLevelMonitoringStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowShardingBalancer(request *model.ShowShardingBalancerRequest) (*model.ShowShardingBalancerResponse, error) {
	requestDef := GenReqDefForShowShardingBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShardingBalancerResponse), nil
	}
}

func (c *DdsClient) ShowShardingBalancerInvoker(request *model.ShowShardingBalancerRequest) *ShowShardingBalancerInvoker {
	requestDef := GenReqDefForShowShardingBalancer()
	return &ShowShardingBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowSlowlogDesensitizationSwitch(request *model.ShowSlowlogDesensitizationSwitchRequest) (*model.ShowSlowlogDesensitizationSwitchResponse, error) {
	requestDef := GenReqDefForShowSlowlogDesensitizationSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSlowlogDesensitizationSwitchResponse), nil
	}
}

func (c *DdsClient) ShowSlowlogDesensitizationSwitchInvoker(request *model.ShowSlowlogDesensitizationSwitchRequest) *ShowSlowlogDesensitizationSwitchInvoker {
	requestDef := GenReqDefForShowSlowlogDesensitizationSwitch()
	return &ShowSlowlogDesensitizationSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowUpgradeDuration(request *model.ShowUpgradeDurationRequest) (*model.ShowUpgradeDurationResponse, error) {
	requestDef := GenReqDefForShowUpgradeDuration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpgradeDurationResponse), nil
	}
}

func (c *DdsClient) ShowUpgradeDurationInvoker(request *model.ShowUpgradeDurationRequest) *ShowUpgradeDurationInvoker {
	requestDef := GenReqDefForShowUpgradeDuration()
	return &ShowUpgradeDurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShrinkInstanceNodes(request *model.ShrinkInstanceNodesRequest) (*model.ShrinkInstanceNodesResponse, error) {
	requestDef := GenReqDefForShrinkInstanceNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShrinkInstanceNodesResponse), nil
	}
}

func (c *DdsClient) ShrinkInstanceNodesInvoker(request *model.ShrinkInstanceNodesRequest) *ShrinkInstanceNodesInvoker {
	requestDef := GenReqDefForShrinkInstanceNodes()
	return &ShrinkInstanceNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) SwitchConfiguration(request *model.SwitchConfigurationRequest) (*model.SwitchConfigurationResponse, error) {
	requestDef := GenReqDefForSwitchConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchConfigurationResponse), nil
	}
}

func (c *DdsClient) SwitchConfigurationInvoker(request *model.SwitchConfigurationRequest) *SwitchConfigurationInvoker {
	requestDef := GenReqDefForSwitchConfiguration()
	return &SwitchConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) SwitchSecondLevelMonitoring(request *model.SwitchSecondLevelMonitoringRequest) (*model.SwitchSecondLevelMonitoringResponse, error) {
	requestDef := GenReqDefForSwitchSecondLevelMonitoring()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchSecondLevelMonitoringResponse), nil
	}
}

func (c *DdsClient) SwitchSecondLevelMonitoringInvoker(request *model.SwitchSecondLevelMonitoringRequest) *SwitchSecondLevelMonitoringInvoker {
	requestDef := GenReqDefForSwitchSecondLevelMonitoring()
	return &SwitchSecondLevelMonitoringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) SwitchSlowlogDesensitization(request *model.SwitchSlowlogDesensitizationRequest) (*model.SwitchSlowlogDesensitizationResponse, error) {
	requestDef := GenReqDefForSwitchSlowlogDesensitization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchSlowlogDesensitizationResponse), nil
	}
}

func (c *DdsClient) SwitchSlowlogDesensitizationInvoker(request *model.SwitchSlowlogDesensitizationRequest) *SwitchSlowlogDesensitizationInvoker {
	requestDef := GenReqDefForSwitchSlowlogDesensitization()
	return &SwitchSlowlogDesensitizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) SwitchSsl(request *model.SwitchSslRequest) (*model.SwitchSslResponse, error) {
	requestDef := GenReqDefForSwitchSsl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchSslResponse), nil
	}
}

func (c *DdsClient) SwitchSslInvoker(request *model.SwitchSslRequest) *SwitchSslInvoker {
	requestDef := GenReqDefForSwitchSsl()
	return &SwitchSslInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) SwitchoverReplicaSet(request *model.SwitchoverReplicaSetRequest) (*model.SwitchoverReplicaSetResponse, error) {
	requestDef := GenReqDefForSwitchoverReplicaSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchoverReplicaSetResponse), nil
	}
}

func (c *DdsClient) SwitchoverReplicaSetInvoker(request *model.SwitchoverReplicaSetRequest) *SwitchoverReplicaSetInvoker {
	requestDef := GenReqDefForSwitchoverReplicaSet()
	return &SwitchoverReplicaSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) UpdateClientNetwork(request *model.UpdateClientNetworkRequest) (*model.UpdateClientNetworkResponse, error) {
	requestDef := GenReqDefForUpdateClientNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClientNetworkResponse), nil
	}
}

func (c *DdsClient) UpdateClientNetworkInvoker(request *model.UpdateClientNetworkRequest) *UpdateClientNetworkInvoker {
	requestDef := GenReqDefForUpdateClientNetwork()
	return &UpdateClientNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) UpdateConfigurationParameter(request *model.UpdateConfigurationParameterRequest) (*model.UpdateConfigurationParameterResponse, error) {
	requestDef := GenReqDefForUpdateConfigurationParameter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigurationParameterResponse), nil
	}
}

func (c *DdsClient) UpdateConfigurationParameterInvoker(request *model.UpdateConfigurationParameterRequest) *UpdateConfigurationParameterInvoker {
	requestDef := GenReqDefForUpdateConfigurationParameter()
	return &UpdateConfigurationParameterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) UpdateEntityConfiguration(request *model.UpdateEntityConfigurationRequest) (*model.UpdateEntityConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateEntityConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEntityConfigurationResponse), nil
	}
}

func (c *DdsClient) UpdateEntityConfigurationInvoker(request *model.UpdateEntityConfigurationRequest) *UpdateEntityConfigurationInvoker {
	requestDef := GenReqDefForUpdateEntityConfiguration()
	return &UpdateEntityConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) UpdateInstanceName(request *model.UpdateInstanceNameRequest) (*model.UpdateInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceNameResponse), nil
	}
}

func (c *DdsClient) UpdateInstanceNameInvoker(request *model.UpdateInstanceNameRequest) *UpdateInstanceNameInvoker {
	requestDef := GenReqDefForUpdateInstanceName()
	return &UpdateInstanceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) UpdateInstancePort(request *model.UpdateInstancePortRequest) (*model.UpdateInstancePortResponse, error) {
	requestDef := GenReqDefForUpdateInstancePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstancePortResponse), nil
	}
}

func (c *DdsClient) UpdateInstancePortInvoker(request *model.UpdateInstancePortRequest) *UpdateInstancePortInvoker {
	requestDef := GenReqDefForUpdateInstancePort()
	return &UpdateInstancePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) UpdateInstanceRemark(request *model.UpdateInstanceRemarkRequest) (*model.UpdateInstanceRemarkResponse, error) {
	requestDef := GenReqDefForUpdateInstanceRemark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceRemarkResponse), nil
	}
}

func (c *DdsClient) UpdateInstanceRemarkInvoker(request *model.UpdateInstanceRemarkRequest) *UpdateInstanceRemarkInvoker {
	requestDef := GenReqDefForUpdateInstanceRemark()
	return &UpdateInstanceRemarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) UpdateReplSetName(request *model.UpdateReplSetNameRequest) (*model.UpdateReplSetNameResponse, error) {
	requestDef := GenReqDefForUpdateReplSetName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateReplSetNameResponse), nil
	}
}

func (c *DdsClient) UpdateReplSetNameInvoker(request *model.UpdateReplSetNameRequest) *UpdateReplSetNameInvoker {
	requestDef := GenReqDefForUpdateReplSetName()
	return &UpdateReplSetNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) UpdateSecurityGroup(request *model.UpdateSecurityGroupRequest) (*model.UpdateSecurityGroupResponse, error) {
	requestDef := GenReqDefForUpdateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityGroupResponse), nil
	}
}

func (c *DdsClient) UpdateSecurityGroupInvoker(request *model.UpdateSecurityGroupRequest) *UpdateSecurityGroupInvoker {
	requestDef := GenReqDefForUpdateSecurityGroup()
	return &UpdateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) UpgradeDatabaseVersion(request *model.UpgradeDatabaseVersionRequest) (*model.UpgradeDatabaseVersionResponse, error) {
	requestDef := GenReqDefForUpgradeDatabaseVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeDatabaseVersionResponse), nil
	}
}

func (c *DdsClient) UpgradeDatabaseVersionInvoker(request *model.UpgradeDatabaseVersionRequest) *UpgradeDatabaseVersionInvoker {
	requestDef := GenReqDefForUpgradeDatabaseVersion()
	return &UpgradeDatabaseVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ListApiVersion(request *model.ListApiVersionRequest) (*model.ListApiVersionResponse, error) {
	requestDef := GenReqDefForListApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionResponse), nil
	}
}

func (c *DdsClient) ListApiVersionInvoker(request *model.ListApiVersionRequest) *ListApiVersionInvoker {
	requestDef := GenReqDefForListApiVersion()
	return &ListApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DdsClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}

func (c *DdsClient) ShowApiVersionInvoker(request *model.ShowApiVersionRequest) *ShowApiVersionInvoker {
	requestDef := GenReqDefForShowApiVersion()
	return &ShowApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
