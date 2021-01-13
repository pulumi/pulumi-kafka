module github.com/pulumi/pulumi-kafka/provider/v2

go 1.14

require (
	github.com/Mongey/terraform-provider-kafka v0.2.11
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.17.0
	github.com/pulumi/pulumi/sdk/v2 v2.17.0
	github.com/satori/go.uuid v1.2.0 // indirect
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
