# k8s-demo

## 部署一个 pure server
> pure server -> 只有 server, 不包含数据库之类的组件
### 编写 yaml
### 初始化 kind
```sh
kind create cluster
```
### 根据 yaml 部署
```sh
kubectl apply -f k8s-pureserver-deployment.yaml
```
### 查看集群状态
```sh
kubectl get nodes
kubectl get pods
```