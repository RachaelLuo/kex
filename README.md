# kex

## Vision
Become a kubernetes web extension and provide a low-threshold entry project for k8s newcomers.

## Quick Start
1. kind creates a cluster
```shell
kind create cluster --name=multi-node
```

2. get kubeconfig and paste it into the project `.config/kube.config`
```shell
kubectl config view --raw
```

3. run
> basic
```shell
go run ./main.go --kubeconfig .config/kube.config
```

> deploy
```shell
kubectl apply -f sample/kex.yaml
```
