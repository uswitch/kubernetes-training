tag = "latest"

docker:
	docker image build -t myapp:$(tag) .

minikube:
	minikube start
	eval $$(minikube docker-env)
	minikube addons enable ingress

hosts:
	echo "$$(minikube ip) myapp.learning.com" >> /etc/hosts
