// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetAwsEcsClusterArn sets provided value as "aws.ecs.cluster.arn" attribute.
func (rb *ResourceBuilder) SetAwsEcsClusterArn(val string) {
	if rb.config.AwsEcsClusterArn.Enabled {
		rb.res.Attributes().PutStr("aws.ecs.cluster.arn", val)
	}
}

// SetAwsEcsLaunchtype sets provided value as "aws.ecs.launchtype" attribute.
func (rb *ResourceBuilder) SetAwsEcsLaunchtype(val string) {
	if rb.config.AwsEcsLaunchtype.Enabled {
		rb.res.Attributes().PutStr("aws.ecs.launchtype", val)
	}
}

// SetAwsEcsTaskArn sets provided value as "aws.ecs.task.arn" attribute.
func (rb *ResourceBuilder) SetAwsEcsTaskArn(val string) {
	if rb.config.AwsEcsTaskArn.Enabled {
		rb.res.Attributes().PutStr("aws.ecs.task.arn", val)
	}
}

// SetAwsEcsTaskFamily sets provided value as "aws.ecs.task.family" attribute.
func (rb *ResourceBuilder) SetAwsEcsTaskFamily(val string) {
	if rb.config.AwsEcsTaskFamily.Enabled {
		rb.res.Attributes().PutStr("aws.ecs.task.family", val)
	}
}

// SetAwsEcsTaskRevision sets provided value as "aws.ecs.task.revision" attribute.
func (rb *ResourceBuilder) SetAwsEcsTaskRevision(val string) {
	if rb.config.AwsEcsTaskRevision.Enabled {
		rb.res.Attributes().PutStr("aws.ecs.task.revision", val)
	}
}

// SetAwsLogGroupArns sets provided value as "aws.log.group.arns" attribute.
func (rb *ResourceBuilder) SetAwsLogGroupArns(val []any) {
	if rb.config.AwsLogGroupArns.Enabled {
		rb.res.Attributes().PutEmptySlice("aws.log.group.arns").FromRaw(val)
	}
}

// SetAwsLogGroupNames sets provided value as "aws.log.group.names" attribute.
func (rb *ResourceBuilder) SetAwsLogGroupNames(val []any) {
	if rb.config.AwsLogGroupNames.Enabled {
		rb.res.Attributes().PutEmptySlice("aws.log.group.names").FromRaw(val)
	}
}

// SetAwsLogStreamArns sets provided value as "aws.log.stream.arns" attribute.
func (rb *ResourceBuilder) SetAwsLogStreamArns(val []any) {
	if rb.config.AwsLogStreamArns.Enabled {
		rb.res.Attributes().PutEmptySlice("aws.log.stream.arns").FromRaw(val)
	}
}

// SetAwsLogStreamNames sets provided value as "aws.log.stream.names" attribute.
func (rb *ResourceBuilder) SetAwsLogStreamNames(val []any) {
	if rb.config.AwsLogStreamNames.Enabled {
		rb.res.Attributes().PutEmptySlice("aws.log.stream.names").FromRaw(val)
	}
}

// SetCloudAccountID sets provided value as "cloud.account.id" attribute.
func (rb *ResourceBuilder) SetCloudAccountID(val string) {
	if rb.config.CloudAccountID.Enabled {
		rb.res.Attributes().PutStr("cloud.account.id", val)
	}
}

// SetCloudAvailabilityZone sets provided value as "cloud.availability_zone" attribute.
func (rb *ResourceBuilder) SetCloudAvailabilityZone(val string) {
	if rb.config.CloudAvailabilityZone.Enabled {
		rb.res.Attributes().PutStr("cloud.availability_zone", val)
	}
}

// SetCloudPlatform sets provided value as "cloud.platform" attribute.
func (rb *ResourceBuilder) SetCloudPlatform(val string) {
	if rb.config.CloudPlatform.Enabled {
		rb.res.Attributes().PutStr("cloud.platform", val)
	}
}

// SetCloudProvider sets provided value as "cloud.provider" attribute.
func (rb *ResourceBuilder) SetCloudProvider(val string) {
	if rb.config.CloudProvider.Enabled {
		rb.res.Attributes().PutStr("cloud.provider", val)
	}
}

// SetCloudRegion sets provided value as "cloud.region" attribute.
func (rb *ResourceBuilder) SetCloudRegion(val string) {
	if rb.config.CloudRegion.Enabled {
		rb.res.Attributes().PutStr("cloud.region", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
