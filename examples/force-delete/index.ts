import * as pulumi from "@pulumi/pulumi";
import * as kafka from "@pulumi/kafka";

// Create a Kafka provider instance with force_delete_resources enabled
const kafkaProvider = new kafka.Provider("kafka-provider", {
    bootstrapServers: ["localhost:9092"], // Replace with your Kafka bootstrap servers
    // Enable force deletion of resources when Kafka is unresponsive
    // This can also be set via the KAFKA_FORCE_DELETE_RESOURCES environment variable
    forceDeleteResources: true,
});

// Create a Kafka topic using our provider with force_delete_resources enabled
const topic = new kafka.Topic("example-topic", {
    name: "example-topic",
    partitions: 1,
    replicationFactor: 1,
    config: {
        "cleanup.policy": "delete",
        "retention.ms": "86400000", // 1 day
    },
}, { provider: kafkaProvider });

// Export the topic name
export const topicName = topic.name;

// Create an ACL using our provider with force_delete_resources enabled
const acl = new kafka.Acl("example-acl", {
    aclResourceName: topic.name,
    aclResourceType: "Topic",
    aclPrincipal: "User:example",
    aclHost: "*",
    aclOperation: "Read",
    aclPermissionType: "Allow",
}, { provider: kafkaProvider });

// Export the ACL ID
export const aclId = acl.id; 