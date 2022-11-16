package datadog

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/validators"
)

func resourceDatadogCloudConfigurationRule() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides a Datadog Cloud Configuration Rule API resource. This can be used to create and manage Datadog cloud configuration rules.",
		CreateContext: nil,
		ReadContext:   nil,
		UpdateContext: nil,
		DeleteContext: nil,
		Importer:      nil,

		Schema: datadogCloudConfigurationRuleSchema(),
	}
}

func datadogCloudConfigurationRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "The name of the rule. Must not be blank.",
			ValidateDiagFunc: validators.ValidateNonBlankStringField,
		},
		"message": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Message for generated findings.", // TODO: check if no message/empty string works
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether the rule is enabled.",
		},
		"severity": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "Severity of generated findings.",
			ValidateDiagFunc: validators.ValidateStringEnumValue("low", "info", "medium", "high", "critical"),
		},
		"filter": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     nil,
			Description: "Filter that will be checked before generating a finding.", // TODO: check what happends when no filter
		},
		"rego_policy": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "Rego policy that will be executed on resources to check for compliance.",
			ValidateDiagFunc: validators.ValidateNonBlankStringField,
		},
		"resource_type": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "Name of the main type of resources to be checked.",
			ValidateDiagFunc: validators.ValidateNonBlankStringField,
		},
		"related_resource_types": {
			Type:             schema.TypeList,
			Optional:         true,
			Default:          []string{},
			Description:      "Name of the related resource types to be checked, for complex rules spanning over several resource types.",
			ValidateDiagFunc: validators.ValidateNonBlankStringListField,
		},
		"notifications": {
			Type:        schema.TypeMap,
			Optional:    true,
			Default:     nil,
			Description: "Notifications of non-compliant findings.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"group_by_fields": {
						Type:             schema.TypeList,
						Required:         true,
						Description:      "Name of the fields by which to group findings.",
						Elem:             &schema.Schema{Type: schema.TypeString},
						MinItems:         1,
						ValidateDiagFunc: validators.ValidateNonBlankStringListField,
					},
					"notified_signals": {
						Type:             schema.TypeList,
						Required:         true,
						Description:      "Name of the notified channels.",
						Elem:             &schema.Schema{Type: schema.TypeString},
						MinItems:         1,
						ValidateDiagFunc: validators.ValidateNonBlankStringListField,
					},
				},
			},
		},
		"tags": {
			Type:             schema.TypeList,
			Optional:         true,
			Default:          []string{},
			Description:      "Tags for generated signals.",
			Elem:             &schema.Schema{Type: schema.TypeString},
			ValidateDiagFunc: validators.ValidateNonBlankStringListField,
		},
	}
}
