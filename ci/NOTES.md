## Reference

Run the following command:

### Vault Environment Setup

```bash
vault write concourse/main/mongo_connection_string value=$MONGO_CONNECTION_STRING
vault write concourse/main/mongo_database value=$MONGO_DATABASE

```

### Execute
```bash
	./run_ci.sh
```

}

