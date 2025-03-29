// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates auto scaling settings on your global tables at once. This operation
// only applies to Version 2019.11.21 (Current) (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html)
// of global tables.
func (c *Client) UpdateTableReplicaAutoScaling(ctx context.Context, params *UpdateTableReplicaAutoScalingInput, optFns ...func(*Options)) (*UpdateTableReplicaAutoScalingOutput, error) {
	if params == nil {
		params = &UpdateTableReplicaAutoScalingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTableReplicaAutoScaling", params, optFns, c.addOperationUpdateTableReplicaAutoScalingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTableReplicaAutoScalingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTableReplicaAutoScalingInput struct {

	// The name of the global table to be updated.
	//
	// This member is required.
	TableName *string

	// Represents the auto scaling settings of the global secondary indexes of the
	// replica to be updated.
	GlobalSecondaryIndexUpdates []types.GlobalSecondaryIndexAutoScalingUpdate

	// Represents the auto scaling settings to be modified for a global table or
	// global secondary index.
	ProvisionedWriteCapacityAutoScalingUpdate *types.AutoScalingSettingsUpdate

	// Represents the auto scaling settings of replicas of the table that will be
	// modified.
	ReplicaUpdates []types.ReplicaAutoScalingUpdate

	noSmithyDocumentSerde
}

type UpdateTableReplicaAutoScalingOutput struct {

	// Returns information about the auto scaling settings of a table with replicas.
	TableAutoScalingDescription *types.TableAutoScalingDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTableReplicaAutoScalingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateTableReplicaAutoScaling{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateTableReplicaAutoScaling{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTableReplicaAutoScaling"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateTableReplicaAutoScalingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTableReplicaAutoScaling(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateTableReplicaAutoScaling(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTableReplicaAutoScaling",
	}
}
