# go_user_rpc
go 的rpc单点登录项目

生成grpc文件：

protoc --proto_path=. --proto_path=%GOPATH%\src --govalidators_out=. --go_out=plugins=grpc:.  protoc\user.proto

增加其他tag：

protoc-go-inject-tag -input=./user/user_grpc/user.pb.go

安装应用：
kubectl apply -f go-server-rpc-deployment.yaml -n go-user-rpc
