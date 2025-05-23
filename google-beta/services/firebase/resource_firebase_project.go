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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/firebase/Project.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package firebase

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func getExistingFirebaseProjectId(config *transport_tpg.Config, d *schema.ResourceData, billingProject string, userAgent string) (string, error) {
	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}")
	if err != nil {
		return "", err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err == nil {
		id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}")
		if err != nil {
			return "", fmt.Errorf("Error constructing id: %s", err)
		}
		return id, nil
	}

	if !transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		return "", err
	}

	return "", nil
}

func ResourceFirebaseProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseProjectCreate,
		Read:   resourceFirebaseProjectRead,
		Delete: resourceFirebaseProjectDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseProjectImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The GCP project display name`,
			},
			"project_number": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The number of the Google Project that Firebase is enabled on.`,
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

func resourceFirebaseProjectCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}:addFirebase")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Project: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Project: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	// Check if Firebase has already been enabled
	existingId, err := getExistingFirebaseProjectId(config, d, billingProject, userAgent)
	if err != nil {
		return fmt.Errorf("Error checking if Firebase is already enabled: %s", err)
	}

	if existingId != "" {
		log.Printf("[DEBUG] Firebase is already enabled for project %s", project)
		d.SetId(existingId)
		return resourceFirebaseProjectRead(d, meta)
	}
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
		return fmt.Errorf("Error creating Project: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = FirebaseOperationWaitTime(
		config, res, project, "Creating Project", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Project: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Project %q: %#v", d.Id(), res)

	return resourceFirebaseProjectRead(d, meta)
}

func resourceFirebaseProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Project: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirebaseProject %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Project: %s", err)
	}

	if err := d.Set("project_number", flattenFirebaseProjectProjectNumber(res["projectNumber"], d, config)); err != nil {
		return fmt.Errorf("Error reading Project: %s", err)
	}
	if err := d.Set("display_name", flattenFirebaseProjectDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Project: %s", err)
	}

	return nil
}

func resourceFirebaseProjectDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Firebase Project resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceFirebaseProjectImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)$",
		"^(?P<project>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseProjectProjectNumber(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseProjectDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
