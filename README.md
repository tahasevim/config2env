# config2env

config2env is a simple CLI tool that exports your configuration files in various formats such as JSON, YAML to Unix environment variables.

**Note:** For now, only JSON and YAML files are supported. 
TOML and HCL formats also will be supported.

## Install

`go get github.com/tahasevim/config2env/cmd/config2env`

## Usage

### [JSON Types](JSON.md)
### [YAML Types](YAML.md)

## Example with Docker

Docker client provides `--env-file` option to read a file of environment variables so that we can avoid specifiying each environment variable by using `-e` flag.

Let's assume that we got an `config.env` with the contents of: 
```
SERVICE_A_ADDR=localhost
SERVICE_A_PORT=8080
SERVICE_B_ADDR=localhost
SERVICE_B_PORT=9090
```

and we get a simply program `main.go` as below:
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}
```

Then you can build an image as below:

```docker
docker build -t testenv -f- . <<EOF
FROM golang:latest
COPY main.go .
CMD ["go", "run", "main.go"]
EOF
```
Finally run built docker image:

```docker
docker run --env-file config.env testenv
```

You will results printed to screen which will be similar to:
```
PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=febf40d13596
SERVICE_A_ADDR=localhost
SERVICE_A_PORT=8080
SERVICE_B_ADDR=localhost
SERVICE_B_PORT=9090
GOLANG_VERSION=1.13
GOPATH=/go
HOME=/root
```

Note that other environment variables come from other image layers that our image is built on top of those layers.

You can also pass these `.env` files using `docker-compose` with `env_file` option. For more detailed information [see the official docs](https://docs.docker.com/compose/environment-variables/)