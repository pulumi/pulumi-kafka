module github.com/pulumi/pulumi-kafka

go 1.12

require (
	github.com/Mongey/terraform-provider-kafka v0.2.3-0.20191020203511-f188c53e3581
	github.com/apparentlymart/go-dump v0.0.0-20190214190832-042adf3cf4a0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform-bridge v1.4.2
	github.com/vmihailenco/msgpack v4.0.1+incompatible // indirect
	github.com/xanzy/ssh-agent v0.2.1 // indirect
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
