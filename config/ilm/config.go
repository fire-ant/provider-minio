package bucket

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_ilm_policy", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "bucket"
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
			base["bucket"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"bucket",
		}
	})
}
