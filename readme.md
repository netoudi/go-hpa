# Desenvolvimento de Aplicações Modernas e Escaláveis com Microserviços - [CodeEducation](https://code.education)

Projeto prático / Iniciando com Docker / Kubernetes e hpa

## Instruções

```bash
# clonar repositório
$ git clone https://github.com/netoudi/go-hpa.git go-hpa

# acessar projeto
$ cd go-hpa

# rodar o minikube
$ minikube start --vm-driver=virtualbox

# configurar hpa
$ kubectl apply -f deployment.yaml
$ kubectl apply -f service.yaml
$ kubectl apply -f hpa.yaml

# checar autoscaler
$ kubectl run -it loader --image=busybox /bin/sh
$ while true; do wget -q -O- http://go-hpa.default.svc.cluster.local:8000; done;
```

## Docker Hub

```bash
$ docker run -d --name go-hpa -p 8080:8000 tineto/go-hpa
```

## Links
[Installing Kubernetes with Minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/)
