package bucket

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_iam_group", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		config.MoveToStatus(r.TerraformResource, "id")
	})
	p.AddResourceConfigurator("minio_iam_group_membership", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		config.MoveToStatus(r.TerraformResource, "id")
		config.MarkAsRequired(r.TerraformResource, "group", "users")
	})
	p.AddResourceConfigurator("minio_iam_group_policy", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		config.MoveToStatus(r.TerraformResource, "id")
		config.MarkAsRequired(r.TerraformResource, "group", "policy")
	})
	p.AddResourceConfigurator("minio_iam_group_policy_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		config.MoveToStatus(r.TerraformResource, "id")
		config.MarkAsRequired(r.TerraformResource, "group_name", "policy_name")
	})
	p.AddResourceConfigurator("minio_iam_group_user_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		config.MoveToStatus(r.TerraformResource, "id")
		config.MarkAsRequired(r.TerraformResource, "group_name", "user_name")
	})
	p.AddResourceConfigurator("minio_iam_policy", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		config.MoveToStatus(r.TerraformResource, "id")
		config.MarkAsRequired(r.TerraformResource, "policy")
	})
	p.AddResourceConfigurator("minio_service_account", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
			base["target_user"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"target_user",
		}
		config.MoveToStatus(r.TerraformResource, "id")
	})
	p.AddResourceConfigurator("minio_iam_user", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		config.MoveToStatus(r.TerraformResource, "id")
	})
	p.AddResourceConfigurator("minio_iam_user_policy_atttachment", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
			base["user_name"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"user_name",
		}
		config.MoveToStatus(r.TerraformResource, "id")
	})
}
