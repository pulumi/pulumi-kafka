module github.com/pulumi/pulumi-kafka/provider/v3

go 1.16

require (
	github.com/Mongey/terraform-provider-kafka v0.3.3
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.11.0
	github.com/pulumi/pulumi/sdk/v3 v3.17.0
)

replace github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
