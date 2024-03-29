CHANGELOG
=========

## Notice (2022-01-06)

*As of this notice, using CHANGELOG.md is DEPRECATED. We will be using [GitHub Releases](https://github.com/pulumi/pulumi-kafka/releases) for this repository*

## HEAD (Unreleased)
_(none)_

---

## 3.1.2 (2022-01-05)
* Upgrade terraform-provider-kafka to v0.4.2

## 3.1.1 (2021-11-18)
* Upgrade to v0.4.1 of the Kafka Terraform Provider

## 3.1.0 (2021-11-11)
* Upgrade to terraform-bridge 3.11.0
* Upgrade to pulumi 3.17.0

## 3.0.1 (2021-05-12)
* Upgrade to v0.3.3 of the Kafka Terraform Provider

## 3.0.0 (2021-04-19)
* Depend on Pulumi 3.0, which includes improvements to Python resource arguments and key translation, Go SDK performance,
  Node SDK performance, general availability of Automation API, and more.

## 2.7.0 (2021-04-12)
* Upgrade to pulumi-terraform-bridge v2.23.0

## 2.6.1 (2021-03-23)
* Upgrade to v0.3.0 of the Kafka Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 2.6.0 (2021-03-16)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 2.5.2 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider
* Avoid storing config from the environment into the state

## 2.5.1 (2021-02-03)
* Upgrade to v0.2.12 of the Kafka Terraform Provider

## 2.5.0 (2021-01-29)
* Upgrade to pulumi-terraform-bridge v2.18.1

## 2.4.4 (2021-01-13)
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 2.4.3 (2020-12-21)
* Upgrade to v0.2.11 of the Kafka Terraform Provider

## 2.4.2 (2020-11-24)
* Upgrade to pulumi-terraform-bridge v2.13.2  
  * This adds support for import specific examples in documentation

## 2.4.1 (2020-11-09)
* Upgrade to pulumi-terraform-bridge v2.12.1

## 2.4.0 (2020-10-26)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 2.3.0 (2020-08-31)
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python

## 2.2.3 (2020-08-14)
* Upgrade to v0.2.10 of the Kafka Terraform Provider

## 2.2.2 (2020-07-21)
* Upgrade to v0.2.9 of the Kafka Terraform Provider

## 2.2.1 (2020-06-23)
* Upgrade to v0.2.7 of the Kafka Terraform Provider

## 2.2.0 (2020-06-11)
* Upgrade to v0.2.6 of the Kafka Terraform Provider
* Switch to GitHub actions for build

## 2.1.3 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 2.1.2 (2020-05-12)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.1.1 (2020-05-01)
* Upgrade to v0.2.5 of the Kafka Terraform Provider

## 2.1.0 (2020-04-28)
* Upgrade to pulumi-terraform-bridge v2.1.0

## 2.0.0 (2020-04-17)
* Upgrade to Pulumi v2.0.0
* Upgrade to pulumi-terraform-bridge v2.0.0

## 1.7.0 (2020-04-14)
* Upgrade to Pulumi v1.13.1
* Upgrade to pulumi-terraform-bridge v1.8.4
* Refactor layout to use Go modules

## 1.6.0 (2020-03-23)
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.5.0 (2020-01-29)
* Upgrade to pulumi-terraform-bridge v1.6.4

## 1.4.0 (2020-01-06)
* Upgrade to pulumi-terraform-bridge v1.5.2
* Upgrade to v0.2.3 of the Kafka Terraform Provider

## 1.3.0 (2019-12-04)
* Upgrade to support go 1.13.x
* Upgrade to pulumi-terraform-bridge v1.4.3

## 1.2.0 (2019-11-15)
* Add support for DotNet SDK Generation

## 1.1.1 (2019-10-21)
* Attribute upstream Terraform provider to Mongey Github Organisation

## 1.1.0 (2019-10-21)
* Upgrade to pulumi-terraform-bridge@2018d3659d

## 1.0.1 (2019-10-17)
* `kafkaAcl` renames `resourceName` to `aclResourceName`
* `kafkaAcl` renames `resourceType` to `aclResourceType`

## 1.0.0 (2019-10-17)
* Initial release of the provider
