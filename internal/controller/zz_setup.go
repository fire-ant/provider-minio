/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	bucket "github.com/fire-ant/provider-minio/internal/controller/bucket/bucket"
	bucketnotification "github.com/fire-ant/provider-minio/internal/controller/bucket/bucketnotification"
	bucketpolicy "github.com/fire-ant/provider-minio/internal/controller/bucket/bucketpolicy"
	bucketversioning "github.com/fire-ant/provider-minio/internal/controller/bucket/bucketversioning"
	object "github.com/fire-ant/provider-minio/internal/controller/bucket/object"
	policy "github.com/fire-ant/provider-minio/internal/controller/bucket/policy"
	group "github.com/fire-ant/provider-minio/internal/controller/iam/group"
	groupmembership "github.com/fire-ant/provider-minio/internal/controller/iam/groupmembership"
	grouppolicy "github.com/fire-ant/provider-minio/internal/controller/iam/grouppolicy"
	grouppolicyattachment "github.com/fire-ant/provider-minio/internal/controller/iam/grouppolicyattachment"
	groupuserattachment "github.com/fire-ant/provider-minio/internal/controller/iam/groupuserattachment"
	policyiam "github.com/fire-ant/provider-minio/internal/controller/iam/policy"
	serviceaccount "github.com/fire-ant/provider-minio/internal/controller/iam/serviceaccount"
	user "github.com/fire-ant/provider-minio/internal/controller/iam/user"
	userpolicyattachment "github.com/fire-ant/provider-minio/internal/controller/iam/userpolicyattachment"
	providerconfig "github.com/fire-ant/provider-minio/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		bucketnotification.Setup,
		bucketpolicy.Setup,
		bucketversioning.Setup,
		object.Setup,
		policy.Setup,
		group.Setup,
		groupmembership.Setup,
		grouppolicy.Setup,
		grouppolicyattachment.Setup,
		groupuserattachment.Setup,
		policyiam.Setup,
		serviceaccount.Setup,
		user.Setup,
		userpolicyattachment.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
