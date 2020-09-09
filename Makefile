tag = "latest"

docker:
	docker image build -f examples/Dockerfile -t myapp:$(tag) .

minikube:
	minikube start
	minikube addons enable ingress
	echo "$$(minikube docker-env)"

hosts:
	echo "$$(minikube ip) myapp.learning.com" >> /etc/hosts
