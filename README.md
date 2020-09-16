# go_user_rpc
go 的rpc单点登录项目

安装应用：
kubectl create namespace go-user-rpc

kubectl apply -f cert.yaml -n go-user-rpc

kubectl apply -f go-server-rpc-deployment.yaml -n go-user-rpc
