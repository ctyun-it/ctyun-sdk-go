package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupPolicyBindRepoApi
type EcsBackupPolicyBindRepoApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupPolicyBindRepoApi(client *ctyunsdk.CtyunClient) *EcsBackupPolicyBindRepoApi {
	return &EcsBackupPolicyBindRepoApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-policy/bind-repo",
		},
	}
}

func (this *EcsBackupPolicyBindRepoApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupPolicyBindRepoRequest) (*EcsBackupPolicyBindRepoResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupPolicyBindRepoRealRequest{
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

	var realResponse EcsBackupPolicyBindRepoRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupPolicyBindRepoResponse{
		RepositoryID: realResponse.RepositoryID,
		PolicyID:     realResponse.PolicyID,
	}, nil
}

type EcsBackupPolicyBindRepoRealRequest struct {
	RegionID     *string `json:"regionID,omitempty"`
	RepositoryID *string `json:"repositoryID,omitempty"`
	PolicyID     *string `json:"policyID,omitempty"`
}

type EcsBackupPolicyBindRepoRequest struct {
	RegionID     *string
	RepositoryID *string
	PolicyID     *string
}

type EcsBackupPolicyBindRepoRealResponse struct {
	RepositoryID string `json:"repositoryID,omitempty"`
	PolicyID     string `json:"policyID,omitempty"`
}

type EcsBackupPolicyBindRepoResponse struct {
	RepositoryID string
	PolicyID     string
}
