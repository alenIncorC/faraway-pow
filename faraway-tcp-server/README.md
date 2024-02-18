

# POW TCP SERVER

## run test

```bash
go test -v $(go list ./... | grep -v /vendor/)
```

### how to run, steps:

1) build images for server and client

```bash
cd faraway-tcp-server

docker build -t tcpserver .

cd faraway-tcp-client

docker build -t tcpclient .
```

2) create user defined network

```bash
docker network create faraway-pow-network
```

3) run built containers and connect it to the bridge

```bash
docker run -it -d --name fapowsrv --network faraway-pow-network tcpserver
docker run -it --name fapowcli --network faraway-pow-network faraway-tcp-client
```
 