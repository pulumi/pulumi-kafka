module github.com/pulumi/pulumi-kafka/provider

go 1.13

require (
	github.com/Mongey/terraform-provider-kafka v0.2.3
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge v1.8.4
	github.com/pulumi/pulumi/sdk v1.13.1
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
