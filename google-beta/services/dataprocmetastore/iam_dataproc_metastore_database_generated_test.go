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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/metastore/Database.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dataprocmetastore_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataprocMetastoreDatabaseIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreDatabaseIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataproc_metastore_database_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s/databases/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-metastore-srv-%s", context["random_suffix"]), "testdb"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDataprocMetastoreDatabaseIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_dataproc_metastore_database_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s/databases/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-metastore-srv-%s", context["random_suffix"]), "testdb"),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataprocMetastoreDatabaseIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataprocMetastoreDatabaseIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataproc_metastore_database_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s/databases/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-metastore-srv-%s", context["random_suffix"]), "testdb"),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataprocMetastoreDatabaseIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreDatabaseIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_dataproc_metastore_database_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_dataproc_metastore_database_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s/databases/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-metastore-srv-%s", context["random_suffix"]), "testdb"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataprocMetastoreDatabaseIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_dataproc_metastore_database_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s/databases/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-metastore-srv-%s", context["random_suffix"]), "testdb"),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataprocMetastoreDatabaseIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataproc_metastore_service" "dpms_service" {
  service_id = "tf-test-metastore-srv-%{random_suffix}"
  location   = "us-central1"

  tier       = "DEVELOPER"

  hive_metastore_config {
    version = "3.1.2"
  }
}

resource "google_dataproc_cluster" "dp_cluster" {
  name   = "tf-test-dpms-tbl-creator-%{random_suffix}"
  region = google_dataproc_metastore_service.dpms_service.location

  cluster_config {
    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    endpoint_config {
      enable_http_port_access = true
    }

    master_config {
      num_instances = 1
      machine_type  = "e2-standard-2"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    metastore_config {
      dataproc_metastore_service = google_dataproc_metastore_service.dpms_service.name
    }
  }
}

resource "google_dataproc_job" "hive" {
  region = google_dataproc_cluster.dp_cluster.region

  force_delete = true
  placement {
    cluster_name = google_dataproc_cluster.dp_cluster.name
  }

  hive_config {
    properties = {
      "database" = "testdb"
    }
    query_list = [
      "DROP DATABASE IF EXISTS testdb CASCADE",
      "CREATE DATABASE testdb",
    ]
  }
}

resource "google_dataproc_metastore_database_iam_member" "foo" {
  project = google_dataproc_metastore_service.dpms_service.project
  location = google_dataproc_metastore_service.dpms_service.location
  service_id = google_dataproc_metastore_service.dpms_service.service_id
  database = google_dataproc_job.hive.hive_config[0].properties["database"]
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataprocMetastoreDatabaseIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataproc_metastore_service" "dpms_service" {
  service_id = "tf-test-metastore-srv-%{random_suffix}"
  location   = "us-central1"

  tier       = "DEVELOPER"

  hive_metastore_config {
    version = "3.1.2"
  }
}

resource "google_dataproc_cluster" "dp_cluster" {
  name   = "tf-test-dpms-tbl-creator-%{random_suffix}"
  region = google_dataproc_metastore_service.dpms_service.location

  cluster_config {
    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    endpoint_config {
      enable_http_port_access = true
    }

    master_config {
      num_instances = 1
      machine_type  = "e2-standard-2"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    metastore_config {
      dataproc_metastore_service = google_dataproc_metastore_service.dpms_service.name
    }
  }
}

resource "google_dataproc_job" "hive" {
  region = google_dataproc_cluster.dp_cluster.region

  force_delete = true
  placement {
    cluster_name = google_dataproc_cluster.dp_cluster.name
  }

  hive_config {
    properties = {
      "database" = "testdb"
    }
    query_list = [
      "DROP DATABASE IF EXISTS testdb CASCADE",
      "CREATE DATABASE testdb",
    ]
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dataproc_metastore_database_iam_policy" "foo" {
  project = google_dataproc_metastore_service.dpms_service.project
  location = google_dataproc_metastore_service.dpms_service.location
  service_id = google_dataproc_metastore_service.dpms_service.service_id
  database = google_dataproc_job.hive.hive_config[0].properties["database"]
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_dataproc_metastore_database_iam_policy" "foo" {
  project = google_dataproc_metastore_service.dpms_service.project
  location = google_dataproc_metastore_service.dpms_service.location
  service_id = google_dataproc_metastore_service.dpms_service.service_id
  database = google_dataproc_job.hive.hive_config[0].properties["database"]
  depends_on = [
    google_dataproc_metastore_database_iam_policy.foo
  ]
}
`, context)
}

func testAccDataprocMetastoreDatabaseIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataproc_metastore_service" "dpms_service" {
  service_id = "tf-test-metastore-srv-%{random_suffix}"
  location   = "us-central1"

  tier       = "DEVELOPER"

  hive_metastore_config {
    version = "3.1.2"
  }
}

resource "google_dataproc_cluster" "dp_cluster" {
  name   = "tf-test-dpms-tbl-creator-%{random_suffix}"
  region = google_dataproc_metastore_service.dpms_service.location

  cluster_config {
    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    endpoint_config {
      enable_http_port_access = true
    }

    master_config {
      num_instances = 1
      machine_type  = "e2-standard-2"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    metastore_config {
      dataproc_metastore_service = google_dataproc_metastore_service.dpms_service.name
    }
  }
}

resource "google_dataproc_job" "hive" {
  region = google_dataproc_cluster.dp_cluster.region

  force_delete = true
  placement {
    cluster_name = google_dataproc_cluster.dp_cluster.name
  }

  hive_config {
    properties = {
      "database" = "testdb"
    }
    query_list = [
      "DROP DATABASE IF EXISTS testdb CASCADE",
      "CREATE DATABASE testdb",
    ]
  }
}

data "google_iam_policy" "foo" {
}

resource "google_dataproc_metastore_database_iam_policy" "foo" {
  project = google_dataproc_metastore_service.dpms_service.project
  location = google_dataproc_metastore_service.dpms_service.location
  service_id = google_dataproc_metastore_service.dpms_service.service_id
  database = google_dataproc_job.hive.hive_config[0].properties["database"]
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataprocMetastoreDatabaseIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataproc_metastore_service" "dpms_service" {
  service_id = "tf-test-metastore-srv-%{random_suffix}"
  location   = "us-central1"

  tier       = "DEVELOPER"

  hive_metastore_config {
    version = "3.1.2"
  }
}

resource "google_dataproc_cluster" "dp_cluster" {
  name   = "tf-test-dpms-tbl-creator-%{random_suffix}"
  region = google_dataproc_metastore_service.dpms_service.location

  cluster_config {
    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    endpoint_config {
      enable_http_port_access = true
    }

    master_config {
      num_instances = 1
      machine_type  = "e2-standard-2"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    metastore_config {
      dataproc_metastore_service = google_dataproc_metastore_service.dpms_service.name
    }
  }
}

resource "google_dataproc_job" "hive" {
  region = google_dataproc_cluster.dp_cluster.region

  force_delete = true
  placement {
    cluster_name = google_dataproc_cluster.dp_cluster.name
  }

  hive_config {
    properties = {
      "database" = "testdb"
    }
    query_list = [
      "DROP DATABASE IF EXISTS testdb CASCADE",
      "CREATE DATABASE testdb",
    ]
  }
}

resource "google_dataproc_metastore_database_iam_binding" "foo" {
  project = google_dataproc_metastore_service.dpms_service.project
  location = google_dataproc_metastore_service.dpms_service.location
  service_id = google_dataproc_metastore_service.dpms_service.service_id
  database = google_dataproc_job.hive.hive_config[0].properties["database"]
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataprocMetastoreDatabaseIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataproc_metastore_service" "dpms_service" {
  service_id = "tf-test-metastore-srv-%{random_suffix}"
  location   = "us-central1"

  tier       = "DEVELOPER"

  hive_metastore_config {
    version = "3.1.2"
  }
}

resource "google_dataproc_cluster" "dp_cluster" {
  name   = "tf-test-dpms-tbl-creator-%{random_suffix}"
  region = google_dataproc_metastore_service.dpms_service.location

  cluster_config {
    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    endpoint_config {
      enable_http_port_access = true
    }

    master_config {
      num_instances = 1
      machine_type  = "e2-standard-2"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    metastore_config {
      dataproc_metastore_service = google_dataproc_metastore_service.dpms_service.name
    }
  }
}

resource "google_dataproc_job" "hive" {
  region = google_dataproc_cluster.dp_cluster.region

  force_delete = true
  placement {
    cluster_name = google_dataproc_cluster.dp_cluster.name
  }

  hive_config {
    properties = {
      "database" = "testdb"
    }
    query_list = [
      "DROP DATABASE IF EXISTS testdb CASCADE",
      "CREATE DATABASE testdb",
    ]
  }
}

resource "google_dataproc_metastore_database_iam_binding" "foo" {
  project = google_dataproc_metastore_service.dpms_service.project
  location = google_dataproc_metastore_service.dpms_service.location
  service_id = google_dataproc_metastore_service.dpms_service.service_id
  database = google_dataproc_job.hive.hive_config[0].properties["database"]
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
