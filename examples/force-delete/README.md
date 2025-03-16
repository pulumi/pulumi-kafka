# Kafka Force Delete Example

This example demonstrates how to use the `force_delete_resources` option to allow deleting Kafka resources when the Kafka cluster is unresponsive.

## Running the Example

1. Start by setting up your Kafka cluster credentials:

```bash
export KAFKA_BOOTSTRAP_SERVERS=your-kafka-bootstrap-servers
```

2. Run `pulumi up` to create the resources:

```bash
pulumi up
```

3. To demonstrate force deletion, you can simulate a Kafka cluster being unavailable by:
   - Shutting down your Kafka cluster
   - Or changing the bootstrap servers to an invalid value:

```bash
export KAFKA_BOOTSTRAP_SERVERS=invalid-server:9092
export KAFKA_FORCE_DELETE_RESOURCES=true
```

4. Run `pulumi destroy` to delete the resources:

```bash
pulumi destroy
```

With `KAFKA_FORCE_DELETE_RESOURCES=true`, the resources will be removed from the Pulumi state even though the Kafka cluster is unreachable. 