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
### 删除部署的内容
```sh
kubectl delete -f k8s-pureserver-deployment.yaml
```
### 查看集群状态
- 常用命令
  ```sh
  kubectl get nodes
  kubectl get pods
  kubectl get service
  kubectl get ingress
  kubectl get namespace
  kubectl describe [resource type] [resource name]
  ```
- 可以使用 `watch` 命令实时查看
  ```sh
  watch kubectl get pods
  ```
- 也可以使用某些插件, 右键点击查看
### 本地使用 minikube 需要进行端口转发
- & 表示在后台运行
- 每次重启 minikube 都需要重新执行
- 每次重建 deployment 也需要重新执行
    ```sh
    kubectl port-forward service/deploymentName 3318:3318 &
    ```

### PV 和 PVC
- deployment 中的 volumeMounts 和 volumes
  - volumeMounts: 挂载到容器中的路径(使用容器中的哪个存储)
    - mysql 需要挂载到 /var/lib/mysql (必须)
  - volumes: 容器中的路径对应的存储(容器中有几个存储)
    - persistentVolumeClaim.claimName: 指定 PVC 的名称(需要与 pvc 中的 name 一致)
- PVC 需要与 PV 绑定
  - PVC 中的 storageClassName 需要与 PV 中的 storageClassName 一致
  - PVC 中的 storage 需要小于等于 PV 中的 capacity.storage
  - PVC 中的 accessModes 需要与 PV 中的 accessModes 一致
  - PVC 中的 selector.matchLabels 需要与 PV 中的 metadata.labels 一致
  
### Ingress and Ingress Controller
- Ingress: 用于将外部请求转发到集群内部的服务
- Ingress Controller: 用于处理 Ingress 的请求
  - Nginx
  - Traefik
  - HAProxy
- 安装 `Ingress-nginx`
  - 安装 `helm`
    ```sh
    brew install helm
    ```
  - [使用 helm 安装 ingress-nginx](https://kubernetes.github.io/ingress-nginx/deploy/#quick-start)

### 遇到的坑
1. 注意代理问题
   1. 如果没有添加本地 `ip` 到 `no_proxy` 中, 会导致无法访问集群内部的服务
   2. 如果使用了代理, 要将 `ingress-nginx` 的服务地址添加到 `no_proxy` 中
2. `curl`的时候后面最后加上 `-v` 可以查看详细信息

### 小技巧
1. 使用 go:build 标签
   - 可以在编译的时候根据不同的条件进行编译
   1. 编译时指定标签
      ```sh
      go build -tags=[tagName] -o main main.go
      ```
### 记录目前的工作流
1. 局域网内 2 台主机, windows 主机中创建了一个 ubuntu 虚拟机, 通过 Mac vscode remote ssh 进行连接
2. 在 ubuntu 虚拟机中安装了 `kubectl`, `minikube`, `helm`
3. 在 mac 中通过 datagrip 连接到了 ubuntu 虚拟机中的 `mysql`, `redis`
   1. 通过 `kubectl port-forward` 进行端口转发 [port]
      1. 要转发的端口可以根据 `kubectl get service` 查看
   2. 通过 vscode 进行端口转发 [port]
   3. 使用 datagrip 连接 localhost:port