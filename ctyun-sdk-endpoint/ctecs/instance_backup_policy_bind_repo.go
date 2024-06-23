package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// InstanceBackupPolicyBindRepoApi
type InstanceBackupPolicyBindRepoApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewInstanceBackupPolicyBindRepoApi(client *ctyunsdk.CtyunClient) *InstanceBackupPolicyBindRepoApi {
	return &InstanceBackupPolicyBindRepoApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-policy/bind-repo",
		},
	}
}

func (this *InstanceBackupPolicyBindRepoApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *InstanceBackupPolicyBindRepoRequest) (*InstanceBackupPolicyBindRepoResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&InstanceBackupPolicyBindRepoRealRequest{
		RegionID:     req.RegionID,
		RepositoryID: req.RepositoryID,
		PolicyID:     req.PolicyID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse InstanceBackupPolicyBindRepoRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &InstanceBackupPolicyBindRepoResponse{
		RepositoryID: realResponse.RepositoryID,
		PolicyID:     realResponse.PolicyID,
	}, nil
}

type InstanceBackupPolicyBindRepoRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	RepositoryID string `json:"repositoryID,omitempty"`
	PolicyID     string `json:"policyID,omitempty"`
}

type InstanceBackupPolicyBindRepoRequest struct {
	RegionID     string
	RepositoryID string
	PolicyID     string
}

type InstanceBackupPolicyBindRepoRealResponse struct {
	RepositoryID string `json:"repositoryID,omitempty"`
	PolicyID     string `json:"policyID,omitempty"`
}

type InstanceBackupPolicyBindRepoResponse struct {
	RepositoryID string
	PolicyID     string
}
