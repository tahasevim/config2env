### JSON Types
Let's have `config.json` with the contents of:
```json
{
    "service_a": {
        "addr": "localhost",
        "port": "8080"
    },

    "service_b": {
        "addr": "localhost",
        "port": "9090"
    }
}
```
Execute the tool as below:

`config2env --input config.json --type json --output config.env `

Generated environment variables are stored `config.env` as below:
```
SERVICE_A_ADDR=localhost
SERVICE_A_PORT=8080
SERVICE_B_ADDR=localhost
SERVICE_B_PORT=9090
```

Array types in JSON files are also supported:

`config.json`:

```json
{
    "service_a": {
        "addr": "localhost",
        "port": ["8080", "8081", "8082"]
    },

    "service_b": {
        "addr": "localhost",
        "port": ["9090", "9091", "9092"]
    }
}
```

`config.env`:

```
SERVICE_A_PORT_1=8080
SERVICE_A_PORT_2=8081
SERVICE_A_PORT_3=8082
SERVICE_A_ADDR=localhost
SERVICE_B_ADDR=localhost
SERVICE_B_PORT_1=9090
SERVICE_B_PORT_2=9091
SERVICE_B_PORT_3=9092
```

**Note:** You can use `--prefix` flag to specify a prefix for environment variables.