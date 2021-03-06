// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DTS API
//
// A description of the DTS API
//

package dts

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//TransferPackageClient a client for TransferPackage
type TransferPackageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTransferPackageClientWithConfigurationProvider Creates a new default TransferPackage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTransferPackageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TransferPackageClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = TransferPackageClient{BaseClient: baseClient}
	client.BasePath = "20171001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TransferPackageClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dts", "https://datatransfer.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TransferPackageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *TransferPackageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AttachDevicesToTransferPackage Attaches Devices to a Transfer Package
func (client TransferPackageClient) AttachDevicesToTransferPackage(ctx context.Context, request AttachDevicesToTransferPackageRequest) (response AttachDevicesToTransferPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.attachDevicesToTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			response = AttachDevicesToTransferPackageResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachDevicesToTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachDevicesToTransferPackageResponse")
	}
	return
}

// attachDevicesToTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) attachDevicesToTransferPackage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs/{id}/transferPackages/{transferPackageLabel}/actions/attachDevices")
	if err != nil {
		return nil, err
	}

	var response AttachDevicesToTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTransferPackage Create a new Transfer Package
func (client TransferPackageClient) CreateTransferPackage(ctx context.Context, request CreateTransferPackageRequest) (response CreateTransferPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateTransferPackageResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTransferPackageResponse")
	}
	return
}

// createTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) createTransferPackage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs/{id}/transferPackages")
	if err != nil {
		return nil, err
	}

	var response CreateTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTransferPackage deletes a transfer Package
func (client TransferPackageClient) DeleteTransferPackage(ctx context.Context, request DeleteTransferPackageRequest) (response DeleteTransferPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteTransferPackageResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTransferPackageResponse")
	}
	return
}

// deleteTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) deleteTransferPackage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/transferJobs/{id}/transferPackages/{transferPackageLabel}")
	if err != nil {
		return nil, err
	}

	var response DeleteTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachDevicesFromTransferPackage Detaches Devices from a Transfer Package
func (client TransferPackageClient) DetachDevicesFromTransferPackage(ctx context.Context, request DetachDevicesFromTransferPackageRequest) (response DetachDevicesFromTransferPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.detachDevicesFromTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			response = DetachDevicesFromTransferPackageResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachDevicesFromTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachDevicesFromTransferPackageResponse")
	}
	return
}

// detachDevicesFromTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) detachDevicesFromTransferPackage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs/{id}/transferPackages/{transferPackageLabel}/actions/detachDevices")
	if err != nil {
		return nil, err
	}

	var response DetachDevicesFromTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTransferPackage Describes a transfer package in detail
func (client TransferPackageClient) GetTransferPackage(ctx context.Context, request GetTransferPackageRequest) (response GetTransferPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetTransferPackageResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTransferPackageResponse")
	}
	return
}

// getTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) getTransferPackage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferPackages/{transferPackageLabel}")
	if err != nil {
		return nil, err
	}

	var response GetTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTransferPackages Lists Transfer Packages associated with a transferJob
func (client TransferPackageClient) ListTransferPackages(ctx context.Context, request ListTransferPackagesRequest) (response ListTransferPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTransferPackages, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListTransferPackagesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTransferPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTransferPackagesResponse")
	}
	return
}

// listTransferPackages implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) listTransferPackages(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferPackages")
	if err != nil {
		return nil, err
	}

	var response ListTransferPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTransferPackage Updates a Transfer Package
func (client TransferPackageClient) UpdateTransferPackage(ctx context.Context, request UpdateTransferPackageRequest) (response UpdateTransferPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateTransferPackageResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTransferPackageResponse")
	}
	return
}

// updateTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) updateTransferPackage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/transferJobs/{id}/transferPackages/{transferPackageLabel}")
	if err != nil {
		return nil, err
	}

	var response UpdateTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
