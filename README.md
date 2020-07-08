# grpc-sample-app
 Simple grpc server and client apps
### Compile proto file to go code
```
protoc proto/reverse.proto --go_out=plugins=grpc:.
```

### Build docker image
```
make build
```

### Push docker image
```
make push
```

### Deploy app to K8S cluster
Adjust `k8s_resources.yaml` and then
```
kubectl apply -f k8s_resources.yaml
```
### Test grpc server by running client locally
```
docker run -e GRPC_SERVER=<your_fqdn> -e GRPC_PORT=80 --net host --add-host <your_fqdn>:<host_IP> aaneci/grpc-sample-app /grpc/client "The app is reversing text"
```