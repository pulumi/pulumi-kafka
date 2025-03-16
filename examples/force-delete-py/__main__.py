import pulumi
import pulumi_kafka as kafka

# Create a Kafka provider instance with force_delete_resources enabled
kafka_provider = kafka.Provider("kafka-provider",
    bootstrap_servers=["localhost:9092"],  # Replace with your Kafka bootstrap servers
    # Enable force deletion of resources when Kafka is unresponsive
    # This can also be set via the KAFKA_FORCE_DELETE_RESOURCES environment variable
    force_delete_resources=True
)

# Create a Kafka topic using our provider with force_delete_resources enabled
topic = kafka.Topic("example-topic",
    name="example-topic",
    partitions=1,
    replication_factor=1,
    config={
        "cleanup.policy": "delete",
        "retention.ms": "86400000",  # 1 day
    },
    opts=pulumi.ResourceOptions(provider=kafka_provider)
)

# Export the topic name
pulumi.export("topic_name", topic.name)

# Create an ACL using our provider with force_delete_resources enabled
acl = kafka.Acl("example-acl",
    acl_resource_name=topic.name,
    acl_resource_type="Topic",
    acl_principal="User:example",
    acl_host="*",
    acl_operation="Read",
    acl_permission_type="Allow",
    opts=pulumi.ResourceOptions(provider=kafka_provider)
)

# Export the ACL ID
pulumi.export("acl_id", acl.id) 