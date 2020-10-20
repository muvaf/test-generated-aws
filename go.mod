module github.com/muvaf/test-generated-aws

go 1.13

require (
	github.com/aws/aws-controllers-k8s v0.0.0-20201019131256-39f1498f3d8b
	github.com/aws/aws-sdk-go v1.34.32
	github.com/crossplane/crossplane-runtime v0.10.0
	github.com/crossplane/crossplane-tools v0.0.0-20201007233256-88b291e145bb
	github.com/crossplane/provider-aws v0.12.0
	github.com/google/go-cmp v0.5.0
	github.com/pkg/errors v0.9.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.18.8
	sigs.k8s.io/controller-runtime v0.6.2
	sigs.k8s.io/controller-tools v0.2.4
)
