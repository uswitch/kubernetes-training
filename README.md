# Kubernetes training

## Building
```
make docker
```
Will build the go app and create a docker image called myapp:latest.

## Minikube
```
make minikube
```
This will start minikube, install the ingress addon and set your docker client to use the minikube docker daemon
```
sudo make hosts
```
This well add an entry to `/etc/hosts` containing your minkube ip so that we can spoof DNS.
## Kubernetes Manifests
Empty templates for the Kubernetes files are in `/templates`, finished versions can be found in `/examples`.
