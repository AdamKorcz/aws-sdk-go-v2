// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarconnections

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsAwsjson10_serializeOpCreateConnection struct {
}

func (*awsAwsjson10_serializeOpCreateConnection) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpCreateConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateConnectionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.CreateConnection")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentCreateConnectionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpCreateHost struct {
}

func (*awsAwsjson10_serializeOpCreateHost) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpCreateHost) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateHostInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.CreateHost")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentCreateHostInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDeleteConnection struct {
}

func (*awsAwsjson10_serializeOpDeleteConnection) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDeleteConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteConnectionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.DeleteConnection")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDeleteConnectionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDeleteHost struct {
}

func (*awsAwsjson10_serializeOpDeleteHost) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDeleteHost) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteHostInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.DeleteHost")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDeleteHostInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpGetConnection struct {
}

func (*awsAwsjson10_serializeOpGetConnection) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpGetConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetConnectionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.GetConnection")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentGetConnectionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpGetHost struct {
}

func (*awsAwsjson10_serializeOpGetHost) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpGetHost) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetHostInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.GetHost")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentGetHostInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListConnections struct {
}

func (*awsAwsjson10_serializeOpListConnections) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListConnections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListConnectionsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.ListConnections")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListConnectionsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListHosts struct {
}

func (*awsAwsjson10_serializeOpListHosts) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListHosts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListHostsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.ListHosts")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListHostsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListTagsForResource struct {
}

func (*awsAwsjson10_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTagsForResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.ListTagsForResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListTagsForResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpTagResource struct {
}

func (*awsAwsjson10_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.TagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpUntagResource struct {
}

func (*awsAwsjson10_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.UntagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentUntagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpUpdateHost struct {
}

func (*awsAwsjson10_serializeOpUpdateHost) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpUpdateHost) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateHostInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("CodeStar_connections_20191201.UpdateHost")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentUpdateHostInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson10_serializeDocumentSecurityGroupIds(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson10_serializeDocumentSubnetIds(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson10_serializeDocumentTag(v *types.Tag, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson10_serializeDocumentTagKeyList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson10_serializeDocumentTagList(v []types.Tag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson10_serializeDocumentTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson10_serializeDocumentVpcConfiguration(v *types.VpcConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.SecurityGroupIds != nil {
		ok := object.Key("SecurityGroupIds")
		if err := awsAwsjson10_serializeDocumentSecurityGroupIds(v.SecurityGroupIds, ok); err != nil {
			return err
		}
	}

	if v.SubnetIds != nil {
		ok := object.Key("SubnetIds")
		if err := awsAwsjson10_serializeDocumentSubnetIds(v.SubnetIds, ok); err != nil {
			return err
		}
	}

	if v.TlsCertificate != nil {
		ok := object.Key("TlsCertificate")
		ok.String(*v.TlsCertificate)
	}

	if v.VpcId != nil {
		ok := object.Key("VpcId")
		ok.String(*v.VpcId)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentCreateConnectionInput(v *CreateConnectionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ConnectionName != nil {
		ok := object.Key("ConnectionName")
		ok.String(*v.ConnectionName)
	}

	if v.HostArn != nil {
		ok := object.Key("HostArn")
		ok.String(*v.HostArn)
	}

	if len(v.ProviderType) > 0 {
		ok := object.Key("ProviderType")
		ok.String(string(v.ProviderType))
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsAwsjson10_serializeDocumentTagList(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentCreateHostInput(v *CreateHostInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Name != nil {
		ok := object.Key("Name")
		ok.String(*v.Name)
	}

	if v.ProviderEndpoint != nil {
		ok := object.Key("ProviderEndpoint")
		ok.String(*v.ProviderEndpoint)
	}

	if len(v.ProviderType) > 0 {
		ok := object.Key("ProviderType")
		ok.String(string(v.ProviderType))
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsAwsjson10_serializeDocumentTagList(v.Tags, ok); err != nil {
			return err
		}
	}

	if v.VpcConfiguration != nil {
		ok := object.Key("VpcConfiguration")
		if err := awsAwsjson10_serializeDocumentVpcConfiguration(v.VpcConfiguration, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentDeleteConnectionInput(v *DeleteConnectionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ConnectionArn != nil {
		ok := object.Key("ConnectionArn")
		ok.String(*v.ConnectionArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentDeleteHostInput(v *DeleteHostInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.HostArn != nil {
		ok := object.Key("HostArn")
		ok.String(*v.HostArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentGetConnectionInput(v *GetConnectionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ConnectionArn != nil {
		ok := object.Key("ConnectionArn")
		ok.String(*v.ConnectionArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentGetHostInput(v *GetHostInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.HostArn != nil {
		ok := object.Key("HostArn")
		ok.String(*v.HostArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListConnectionsInput(v *ListConnectionsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.HostArnFilter != nil {
		ok := object.Key("HostArnFilter")
		ok.String(*v.HostArnFilter)
	}

	if v.MaxResults != 0 {
		ok := object.Key("MaxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if len(v.ProviderTypeFilter) > 0 {
		ok := object.Key("ProviderTypeFilter")
		ok.String(string(v.ProviderTypeFilter))
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListHostsInput(v *ListHostsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != 0 {
		ok := object.Key("MaxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListTagsForResourceInput(v *ListTagsForResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("ResourceArn")
		ok.String(*v.ResourceArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("ResourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsAwsjson10_serializeDocumentTagList(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentUntagResourceInput(v *UntagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("ResourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.TagKeys != nil {
		ok := object.Key("TagKeys")
		if err := awsAwsjson10_serializeDocumentTagKeyList(v.TagKeys, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentUpdateHostInput(v *UpdateHostInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.HostArn != nil {
		ok := object.Key("HostArn")
		ok.String(*v.HostArn)
	}

	if v.ProviderEndpoint != nil {
		ok := object.Key("ProviderEndpoint")
		ok.String(*v.ProviderEndpoint)
	}

	if v.VpcConfiguration != nil {
		ok := object.Key("VpcConfiguration")
		if err := awsAwsjson10_serializeDocumentVpcConfiguration(v.VpcConfiguration, ok); err != nil {
			return err
		}
	}

	return nil
}
