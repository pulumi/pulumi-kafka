module github.com/pulumi/pulumi-kafka

go 1.12

require (
	cloud.google.com/go/logging v1.0.0 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/Mongey/terraform-provider-kafka v0.2.3-0.20191017121936-5e7910ec5835
	github.com/hashicorp/terraform v0.12.10
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.3.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190905205929-ed0b5c29edd1
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/ulikunitz/xz v0.5.6 // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
