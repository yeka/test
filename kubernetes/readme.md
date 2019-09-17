# Kubernetes

There are several ways to test kubernetes in local machine.

## Minikube

Minikube creates a virtual machine, it kinda slow to provision. But it will get you started on using kubernetes.

1. Get `kubctl` (https://kubernetes.io/docs/tasks/tools/install-kubectl/)
```
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x kubectl
mv kubectl /usr/local/bin/
```
2. Get `minikube` (https://kubernetes.io/docs/tasks/tools/install-minikube/)
```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64
chmod +x minikube
mv minikube /usr/local/bin/
```
3. Start `minikube` by calling ```minikube start```. If it fails, you might need to install [VirtualBox](https://www.virtualbox.org/wiki/Downloads).
4. Run ```kubectl create deployment webserver --image=nginx:alpine```
5. Use ```kubectl get deployment``` to see if it's deployed.
6. Use ```kubectl get pods``` to see the pods under deployment.
7. Try deleting the pod ```kubectl delete pods [pod_name]```, new pod will be created.
8. Exposing the nginx ```kubectl expose deployment webserver --type=LoadBalancer --port=80```
9. Check available service using ```kubectl get services```
10. Get the url using ```minikube service webserver --url```. Try openning the url in a browser.
11. Minikube has dashboard, run ```minikube dashboard```
12. If you done playing, you can clean it up by removing the VM using ```minikube delete```

Reference: https://www.youtube.com/watch?v=BDrcUjOczsE

## KinD

KinD is a short for `Kubernetes in Docker`. With this, you don't need VM. Docker is all you need.

1. Download ```kind``` from https://github.com/kubernetes-sigs/kind/releases. Chose the binary that suits your operating system.
```
curl -Lo ./kind https://github.com/kubernetes-sigs/kind/releases/download/v0.5.1/kind-$(uname)-amd64
chmod +x ./kind
mv kind /usr/local/bin/kind
```
2. Run ```kind create cluster```. It will download docker image `kindest/node`, which is about 500MB. So depend on the speed of your internet connection, it might take a while. Once it done, it will tell you to run these:
```
export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"
kubectl cluster-info
```
3. You can try to run ```kubectl port-forward svc/webserver 8080:80``` and access `127.0.0.1:8080` to get to nginx. But when the pods get deleted, you'll lost access to it, even if the pods get replaced kubernetes.
4. Clean up:
```
unset KUBECONFIG
kind delete cluster
```
Well, that's how far I could go with KinD for the day.

References:
- https://www.bogotobogo.com/DevOps/Docker/Docker-Kubernetes-Multi-Node-Local-Clusters-kind.php
- https://banzaicloud.com/blog/kind-ingress/
