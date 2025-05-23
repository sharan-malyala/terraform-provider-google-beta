// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/edgecontainer/NodePool.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package edgecontainer

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceEdgecontainerNodePool() *schema.Resource {
	return &schema.Resource{
		Create: resourceEdgecontainerNodePoolCreate,
		Read:   resourceEdgecontainerNodePoolRead,
		Update: resourceEdgecontainerNodePoolUpdate,
		Delete: resourceEdgecontainerNodePoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceEdgecontainerNodePoolImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(480 * time.Minute),
			Update: schema.DefaultTimeout(480 * time.Minute),
			Delete: schema.DefaultTimeout(480 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the target Distributed Cloud Edge Cluster.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the resource.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource name of the node pool.`,
			},
			"node_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `The number of nodes in the pool.`,
			},
			"node_location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the Google Distributed Cloud Edge zone where this node pool will be created. For example: 'us-central1-edge-customer-a'.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels associated with this resource.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"local_disk_encryption": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Local disk encryption options. This field is only used when enabling CMEK support.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The Cloud KMS CryptoKey e.g. projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey} to use for protecting node local disks.
If not specified, a Google-managed key will be used instead.`,
						},
						"kms_key_active_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The Cloud KMS CryptoKeyVersion currently in use for protecting node local disks. Only applicable if kmsKey is set.`,
						},
						"kms_key_state": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `Availability of the Cloud KMS CryptoKey. If not KEY_AVAILABLE, then nodes may go offline as they cannot access their local data.
This can be caused by a lack of permissions to use the key, or if the key is disabled or deleted.`,
						},
					},
				},
			},
			"machine_filter": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `Only machines matching this filter will be allowed to join the node pool.
The filtering language accepts strings like "name=<name>", and is
documented in more detail in [AIP-160](https://google.aip.dev/160).`,
			},
			"node_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `Configuration for each node in the NodePool`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"labels": {
							Type:        schema.TypeMap,
							Computed:    true,
							Optional:    true,
							Description: `"The Kubernetes node labels"`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the node pool was created.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"node_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The lowest release version among all worker nodes.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the node pool was last updated.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceEdgecontainerNodePoolCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nodeLocationProp, err := expandEdgecontainerNodePoolNodeLocation(d.Get("node_location"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_location"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeLocationProp)) && (ok || !reflect.DeepEqual(v, nodeLocationProp)) {
		obj["nodeLocation"] = nodeLocationProp
	}
	nodeCountProp, err := expandEdgecontainerNodePoolNodeCount(d.Get("node_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeCountProp)) && (ok || !reflect.DeepEqual(v, nodeCountProp)) {
		obj["nodeCount"] = nodeCountProp
	}
	machineFilterProp, err := expandEdgecontainerNodePoolMachineFilter(d.Get("machine_filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("machine_filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(machineFilterProp)) && (ok || !reflect.DeepEqual(v, machineFilterProp)) {
		obj["machineFilter"] = machineFilterProp
	}
	localDiskEncryptionProp, err := expandEdgecontainerNodePoolLocalDiskEncryption(d.Get("local_disk_encryption"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("local_disk_encryption"); !tpgresource.IsEmptyValue(reflect.ValueOf(localDiskEncryptionProp)) && (ok || !reflect.DeepEqual(v, localDiskEncryptionProp)) {
		obj["localDiskEncryption"] = localDiskEncryptionProp
	}
	nodeConfigProp, err := expandEdgecontainerNodePoolNodeConfig(d.Get("node_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeConfigProp)) && (ok || !reflect.DeepEqual(v, nodeConfigProp)) {
		obj["nodeConfig"] = nodeConfigProp
	}
	labelsProp, err := expandEdgecontainerNodePoolEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{EdgecontainerBasePath}}projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools?nodePoolId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NodePool: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NodePool: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating NodePool: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = EdgecontainerOperationWaitTime(
		config, res, project, "Creating NodePool", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NodePool: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NodePool %q: %#v", d.Id(), res)

	return resourceEdgecontainerNodePoolRead(d, meta)
}

func resourceEdgecontainerNodePoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{EdgecontainerBasePath}}projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NodePool: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("EdgecontainerNodePool %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}

	if err := d.Set("create_time", flattenEdgecontainerNodePoolCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}
	if err := d.Set("update_time", flattenEdgecontainerNodePoolUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}
	if err := d.Set("labels", flattenEdgecontainerNodePoolLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}
	if err := d.Set("node_location", flattenEdgecontainerNodePoolNodeLocation(res["nodeLocation"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}
	if err := d.Set("node_count", flattenEdgecontainerNodePoolNodeCount(res["nodeCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}
	if err := d.Set("machine_filter", flattenEdgecontainerNodePoolMachineFilter(res["machineFilter"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}
	if err := d.Set("local_disk_encryption", flattenEdgecontainerNodePoolLocalDiskEncryption(res["localDiskEncryption"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}
	if err := d.Set("node_version", flattenEdgecontainerNodePoolNodeVersion(res["nodeVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}
	if err := d.Set("node_config", flattenEdgecontainerNodePoolNodeConfig(res["nodeConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}
	if err := d.Set("terraform_labels", flattenEdgecontainerNodePoolTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}
	if err := d.Set("effective_labels", flattenEdgecontainerNodePoolEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodePool: %s", err)
	}

	return nil
}

func resourceEdgecontainerNodePoolUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NodePool: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	nodeCountProp, err := expandEdgecontainerNodePoolNodeCount(d.Get("node_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nodeCountProp)) {
		obj["nodeCount"] = nodeCountProp
	}
	machineFilterProp, err := expandEdgecontainerNodePoolMachineFilter(d.Get("machine_filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("machine_filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, machineFilterProp)) {
		obj["machineFilter"] = machineFilterProp
	}
	localDiskEncryptionProp, err := expandEdgecontainerNodePoolLocalDiskEncryption(d.Get("local_disk_encryption"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("local_disk_encryption"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, localDiskEncryptionProp)) {
		obj["localDiskEncryption"] = localDiskEncryptionProp
	}
	nodeConfigProp, err := expandEdgecontainerNodePoolNodeConfig(d.Get("node_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nodeConfigProp)) {
		obj["nodeConfig"] = nodeConfigProp
	}
	labelsProp, err := expandEdgecontainerNodePoolEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{EdgecontainerBasePath}}projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating NodePool %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("node_count") {
		updateMask = append(updateMask, "nodeCount")
	}

	if d.HasChange("machine_filter") {
		updateMask = append(updateMask, "machineFilter")
	}

	if d.HasChange("local_disk_encryption") {
		updateMask = append(updateMask, "localDiskEncryption")
	}

	if d.HasChange("node_config") {
		updateMask = append(updateMask, "nodeConfig")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating NodePool %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating NodePool %q: %#v", d.Id(), res)
		}

		err = EdgecontainerOperationWaitTime(
			config, res, project, "Updating NodePool", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceEdgecontainerNodePoolRead(d, meta)
}

func resourceEdgecontainerNodePoolDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NodePool: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{EdgecontainerBasePath}}projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting NodePool %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "NodePool")
	}

	err = EdgecontainerOperationWaitTime(
		config, res, project, "Deleting NodePool", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NodePool %q: %#v", d.Id(), res)
	return nil
}

func resourceEdgecontainerNodePoolImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/clusters/(?P<cluster>[^/]+)/nodePools/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<cluster>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<cluster>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenEdgecontainerNodePoolCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerNodePoolUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerNodePoolLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenEdgecontainerNodePoolNodeLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerNodePoolNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenEdgecontainerNodePoolMachineFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerNodePoolLocalDiskEncryption(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key"] =
		flattenEdgecontainerNodePoolLocalDiskEncryptionKmsKey(original["kmsKey"], d, config)
	transformed["kms_key_active_version"] =
		flattenEdgecontainerNodePoolLocalDiskEncryptionKmsKeyActiveVersion(original["kmsKeyActiveVersion"], d, config)
	transformed["kms_key_state"] =
		flattenEdgecontainerNodePoolLocalDiskEncryptionKmsKeyState(original["kmsKeyState"], d, config)
	return []interface{}{transformed}
}
func flattenEdgecontainerNodePoolLocalDiskEncryptionKmsKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerNodePoolLocalDiskEncryptionKmsKeyActiveVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerNodePoolLocalDiskEncryptionKmsKeyState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerNodePoolNodeVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerNodePoolNodeConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["labels"] =
		flattenEdgecontainerNodePoolNodeConfigLabels(original["labels"], d, config)
	return []interface{}{transformed}
}
func flattenEdgecontainerNodePoolNodeConfigLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerNodePoolTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenEdgecontainerNodePoolEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandEdgecontainerNodePoolNodeLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolMachineFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolLocalDiskEncryption(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKey, err := expandEdgecontainerNodePoolLocalDiskEncryptionKmsKey(original["kms_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKey"] = transformedKmsKey
	}

	transformedKmsKeyActiveVersion, err := expandEdgecontainerNodePoolLocalDiskEncryptionKmsKeyActiveVersion(original["kms_key_active_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyActiveVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyActiveVersion"] = transformedKmsKeyActiveVersion
	}

	transformedKmsKeyState, err := expandEdgecontainerNodePoolLocalDiskEncryptionKmsKeyState(original["kms_key_state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyState"] = transformedKmsKeyState
	}

	return transformed, nil
}

func expandEdgecontainerNodePoolLocalDiskEncryptionKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolLocalDiskEncryptionKmsKeyActiveVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolLocalDiskEncryptionKmsKeyState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolNodeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLabels, err := expandEdgecontainerNodePoolNodeConfigLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	return transformed, nil
}

func expandEdgecontainerNodePoolNodeConfigLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandEdgecontainerNodePoolEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
