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

// SetDeploymentEnvironment sets provided value as "deployment.environment" attribute.
func (rb *ResourceBuilder) SetDeploymentEnvironment(val string) {
	if rb.config.DeploymentEnvironment.Enabled {
		rb.res.Attributes().PutStr("deployment.environment", val)
	}
}

// SetServiceInstanceID sets provided value as "service.instance.id" attribute.
func (rb *ResourceBuilder) SetServiceInstanceID(val string) {
	if rb.config.ServiceInstanceID.Enabled {
		rb.res.Attributes().PutStr("service.instance.id", val)
	}
}

// SetServiceVersion sets provided value as "service.version" attribute.
func (rb *ResourceBuilder) SetServiceVersion(val string) {
	if rb.config.ServiceVersion.Enabled {
		rb.res.Attributes().PutStr("service.version", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
