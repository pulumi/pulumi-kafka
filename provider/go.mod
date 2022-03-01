module github.com/pulumi/pulumi-kafka/provider/v3

go 1.16

require (
	github.com/Mongey/terraform-provider-kafka v0.4.2
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.19.1
	github.com/pulumi/pulumi/sdk/v3 v3.25.0
)

replace (
	github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211019194827-62530c6537a4
)
