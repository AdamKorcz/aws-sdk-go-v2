// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the count of compliant and noncompliant conformance packs across all AWS
// Accounts and AWS Regions in an aggregator. You can filter based on AWS Account
// ID or AWS Region. The results can return an empty result page, but if you have a
// nextToken, the results are displayed on the next page.
func (c *Client) GetAggregateConformancePackComplianceSummary(ctx context.Context, params *GetAggregateConformancePackComplianceSummaryInput, optFns ...func(*Options)) (*GetAggregateConformancePackComplianceSummaryOutput, error) {
	if params == nil {
		params = &GetAggregateConformancePackComplianceSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAggregateConformancePackComplianceSummary", params, optFns, addOperationGetAggregateConformancePackComplianceSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAggregateConformancePackComplianceSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAggregateConformancePackComplianceSummaryInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// Filters the results based on the
	// AggregateConformancePackComplianceSummaryFilters object.
	Filters *types.AggregateConformancePackComplianceSummaryFilters

	// Groups the result based on AWS Account ID or AWS Region.
	GroupByKey types.AggregateConformancePackComplianceSummaryGroupKey

	// The maximum number of results returned on each page. The default is maximum. If
	// you specify 0, AWS Config uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

type GetAggregateConformancePackComplianceSummaryOutput struct {

	// Returns a list of AggregateConformancePackComplianceSummary object.
	AggregateConformancePackComplianceSummaries []types.AggregateConformancePackComplianceSummary

	// Groups the result based on AWS Account ID or AWS Region.
	GroupByKey *string

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAggregateConformancePackComplianceSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAggregateConformancePackComplianceSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAggregateConformancePackComplianceSummary{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAggregateConformancePackComplianceSummaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAggregateConformancePackComplianceSummary(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetAggregateConformancePackComplianceSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetAggregateConformancePackComplianceSummary",
	}
}
