module github.com/pulumi/pulumi-kafka/provider/v2

go 1.14

require (
	github.com/Mongey/terraform-provider-kafka v0.2.7
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.5.3
	github.com/pulumi/pulumi/sdk/v2 v2.5.1-0.20200701223250-45d2fa95d60b
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
