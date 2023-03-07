```shell
goctl api go -api *.api -dir "./" --style=gozero --home="goctl"
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api *.api -dir "./"
rm Dockerfile
goctl docker -go *.go
rm k8s.yaml
goctl kube deploy -name "ooey-api" -namespace "ooey-api" -image "./" -o "k8s.yaml" -port "1001"
kubectl create namespace "ooey-api"
kubectl apply -f k8s.yaml
go run "ooey.go"
```