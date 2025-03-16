# CHANGELOG

## HEAD (Unreleased)

* Add `force_delete_resources` option to allow deleting Kafka resources when the Kafka cluster is unresponsive
  * This can be set via the `KAFKA_FORCE_DELETE_RESOURCES` environment variable or in the provider configuration
  * When enabled, resources will be removed from state even if the Kafka cluster cannot be reached

## 3.0.0 (2023-01-17)

* Upgrade to v3.0.0 of the pulumi-terraform-bridge

## 2.0.0 (2022-06-29)

* Upgrade to v2.0.0 of the pulumi-terraform-bridge

## 1.5.0 (2022-05-25)

* Upgrade to v0.9.0 of the Kafka Terraform Provider

## 1.4.0 (2022-03-30)

* Upgrade to v4.5.0 of the pulumi-terraform-bridge

## 1.3.0 (2021-11-11)

* Upgrade to v4.3.0 of the pulumi-terraform-bridge

## 1.2.0 (2021-11-04)

* Upgrade to terraform-bridge 3.11.0
* Upgrade to pulumi 3.17.0

## 1.1.0 (2021-10-19)

* Upgrade to v0.5.0 of the Kafka Terraform Provider

## 1.0.0 (2021-04-19)

* Depend on Pulumi 3.0, which includes improvements to Python resource arguments and key translation, Go SDK performance,
  Node SDK performance, general availability of Automation API, and more.

## 0.4.0 (2021-04-12)

* Upgrade to pulumi-terraform-bridge v2.23.0

## 0.3.1 (2021-03-23)

* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 0.3.0 (2021-03-15)

* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 0.2.0 (2021-02-19)

* Upgrade to v0.3.0 of the Kafka Terraform Provider

## 0.1.0 (2021-01-19)

* Initial creation of the provider against v0.2.11 of the Kafka Terraform Provider 