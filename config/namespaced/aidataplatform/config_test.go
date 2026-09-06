package aidataplatform

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestConfigureAIDataPlatformMarksAdminCredentialSensitive(t *testing.T) {
	r := &config.Resource{TerraformResource: &schema.Resource{Schema: map[string]*schema.Schema{
		"vector_db_admin_cred": {},
	}}}

	configureAIDataPlatform(r)

	if !r.TerraformResource.Schema["vector_db_admin_cred"].Sensitive {
		t.Error("vector_db_admin_cred is not marked sensitive")
	}
}
