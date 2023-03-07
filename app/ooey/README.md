```shell
goctl api go -api *.api -dir "./" --style=gozero --home="./goctl"
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api *.api -dir "./"
rm Dockerfile
goctl docker -go *.go
rm k8s.yaml
goctl kube deploy -name "greet1" -namespace "greet-api" -image ./ -o "k8s.yaml" -port "1001"
kubectl create namespace "greet-api"
kubectl apply -f k8s.yaml
go run greet.go
```