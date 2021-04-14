// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies an existing backup within the same AWS account to another Region
// (cross-Region copy) or within the same Region (in-Region copy). You can have up
// to five backup copy requests in progress to a single destination Region per
// account. You can use cross-Region backup copies for cross-region disaster
// recovery. You periodically take backups and copy them to another Region so that
// in the event of a disaster in the primary Region, you can restore from backup
// and recover availability quickly in the other Region. You can make cross-Region
// copies only within your AWS partition. You can also use backup copies to clone
// your file data set to another Region or within the same Region. You can use the
// SourceRegion parameter to specify the AWS Region from which the backup will be
// copied. For example, if you make the call from the us-west-1 Region and want to
// copy a backup from the us-east-2 Region, you specify us-east-2 in the
// SourceRegion parameter to make a cross-Region copy. If you don't specify a
// Region, the backup copy is created in the same Region where the request is sent
// from (in-Region copy). For more information on creating backup copies, see
// Copying backups
// (https://docs.aws.amazon.com/fsx/latest/WindowsGuide/copy-backups.html) in the
// Amazon FSx for Windows User Guide and Copying backups
// (https://docs.aws.amazon.com/fsx/latest/LustreGuide/copy-backups.html) in the
// Amazon FSx for Lustre User Guide.
func (c *Client) CopyBackup(ctx context.Context, params *CopyBackupInput, optFns ...func(*Options)) (*CopyBackupOutput, error) {
	if params == nil {
		params = &CopyBackupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyBackup", params, optFns, addOperationCopyBackupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyBackupInput struct {

	// The ID of the source backup. Specifies the ID of the backup that is being
	// copied.
	//
	// This member is required.
	SourceBackupId *string

	// (Optional) An idempotency token for resource creation, in a string of up to 64
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the AWS Command Line Interface (AWS CLI) or an AWS SDK.
	ClientRequestToken *string

	// A boolean flag indicating whether tags from the source backup should be copied
	// to the backup copy. This value defaults to false. If you set CopyTags to true
	// and the source backup has existing tags, you can use the Tags parameter to
	// create new tags, provided that the sum of the source backup tags and the new
	// tags doesn't exceed 50. Both sets of tags are merged. If there are tag conflicts
	// (for example, two tags with the same key but different values), the tags created
	// with the Tags parameter take precedence.
	CopyTags *bool

	// The ID of the AWS Key Management Service (AWS KMS) key used to encrypt the file
	// system's data for Amazon FSx for Windows File Server file systems and Amazon FSx
	// for Lustre PERSISTENT_1 file systems at rest. In either case, if not specified,
	// the Amazon FSx managed key is used. The Amazon FSx for Lustre SCRATCH_1 and
	// SCRATCH_2 file systems are always encrypted at rest using Amazon FSx managed
	// keys. For more information, see Encrypt
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_Encrypt.html) in the
	// AWS Key Management Service API Reference.
	KmsKeyId *string

	// The source AWS Region of the backup. Specifies the AWS Region from which the
	// backup is being copied. The source and destination Regions must be in the same
	// AWS partition. If you don't specify a Region, it defaults to the Region where
	// the request is sent from (in-Region copy).
	SourceRegion *string

	// A list of Tag values, with a maximum of 50 elements.
	Tags []types.Tag
}

type CopyBackupOutput struct {

	// A backup of an Amazon FSx file system.
	Backup *types.Backup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCopyBackupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCopyBackup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCopyBackup{}, middleware.After)
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
	if err = addIdempotencyToken_opCopyBackupMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCopyBackupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyBackup(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCopyBackup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCopyBackup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCopyBackup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CopyBackupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CopyBackupInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCopyBackupMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCopyBackup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCopyBackup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "CopyBackup",
	}
}
