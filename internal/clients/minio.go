/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/fire-ant/provider-minio/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal minio credentials as JSON"
)

const (
	// minio credentials
	server       = "minio_server"
	user         = "minio_user"
	password     = "minio_password"
	sessionToken = "minio_session_token"
	region       = "minio_region"
	apiVersion   = "minio_api_version"
	ssl          = "minio_ssl"
	insecure     = "minio_insecure"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}
		if v, ok := creds[server]; ok {
			ps.Configuration[server] = v
		}
		if v, ok := creds[user]; ok {
			ps.Configuration[user] = v
		}
		if v, ok := creds[password]; ok {
			ps.Configuration[password] = v
		}
		if v, ok := creds[sessionToken]; ok {
			ps.Configuration[sessionToken] = v
		}
		if v, ok := creds[region]; ok {
			ps.Configuration[region] = v
		}
		if v, ok := creds[apiVersion]; ok {
			ps.Configuration[apiVersion] = v
		}
		if v, ok := creds[ssl]; ok {
			ps.Configuration[ssl] = v
		}
		if v, ok := creds[insecure]; ok {
			ps.Configuration[insecure] = v
		}
		return ps, nil
	}
}
