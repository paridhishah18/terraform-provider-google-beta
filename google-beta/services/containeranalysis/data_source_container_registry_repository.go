// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package containeranalysis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func DataSourceGoogleContainerRepo() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Container Registry is deprecated. Effective March 18, 2025, Container Registry is shut down and writing images to Container Registry is unavailable. Resource will be removed in future major release.",
		Read:               containerRegistryRepoRead,
		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"repository_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func containerRegistryRepoRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error setting project: %s", err)
	}
	region, ok := d.GetOk("region")
	escapedProject := strings.Replace(project, ":", "/", -1)
	if ok && region != nil && region != "" {
		if err := d.Set("repository_url", fmt.Sprintf("%s.gcr.io/%s", region, escapedProject)); err != nil {
			return fmt.Errorf("Error setting repository_url: %s", err)
		}
	} else {
		if err := d.Set("repository_url", fmt.Sprintf("gcr.io/%s", escapedProject)); err != nil {
			return fmt.Errorf("Error setting repository_url: %s", err)
		}
	}
	d.SetId(d.Get("repository_url").(string))
	return nil
}
