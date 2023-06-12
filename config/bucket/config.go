package bucket

import "github.com/upbound/upjet/pkg/config"

const (
	bucket = "bucket"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_s3_bucket", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = bucket
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
			base["bucket"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"bucket",
			"bucket_prefix",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_notification", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = bucket
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
			base["bucket"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_policy", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = bucket
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
			base["bucket"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_versioning", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = bucket
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
			base["bucket"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_object", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = bucket
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
			base["bucket_name"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"bucket_name",
		}
	})
}
