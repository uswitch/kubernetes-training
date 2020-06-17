# Kubernetes training

## Building
```
make docker
```
Will build the go app and create a docker image called myapp:latest.

## Minikube
To start minikube and install the ingress addon run:
```
make minikube
```

Afterwards, to set your docker client to use the minikube docker daemon run:
```
eval $(minikube -p minikube docker-env)
```
Finally, this will add an entry to `/etc/hosts` containing your minkube ip so that we can spoof DNS.
```
sudo make hosts
```
## Kubernetes Manifests
Empty templates for the Kubernetes files are in `/templates`, finished versions can be found in `/examples`.
