package aidataplatform

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures AI Data Platform resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_ai_data_platform_ai_data_platform", configureAIDataPlatform)
}

func configureAIDataPlatform(r *config.Resource) {
	r.TerraformResource.Schema["vector_db_admin_cred"].Sensitive = true
}
