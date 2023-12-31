/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	bucket "github.com/fire-ant/provider-minio/config/bucket"
	iam "github.com/fire-ant/provider-minio/config/iam"
	ilm "github.com/fire-ant/provider-minio/config/ilm"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "minio"
	modulePath     = "github.com/fire-ant/provider-minio"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		bucket.Configure,
		ilm.Configure,
		iam.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
