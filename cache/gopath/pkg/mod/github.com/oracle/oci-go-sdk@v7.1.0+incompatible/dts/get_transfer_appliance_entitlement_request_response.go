// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetTransferApplianceEntitlementRequest wrapper for the GetTransferApplianceEntitlement operation
type GetTransferApplianceEntitlementRequest struct {

	// Tenant Id
	TenantId *string `mandatory:"true" contributesTo:"path" name:"tenantId"`

	// opc request id
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetTransferApplianceEntitlementRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetTransferApplianceEntitlementRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetTransferApplianceEntitlementRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetTransferApplianceEntitlementResponse wrapper for the GetTransferApplianceEntitlement operation
type GetTransferApplianceEntitlementResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The TransferApplianceEntitlement instance
	TransferApplianceEntitlement `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetTransferApplianceEntitlementResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetTransferApplianceEntitlementResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
