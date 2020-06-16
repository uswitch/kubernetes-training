tag = "latest"

docker:
	docker image build -t myapp:$(tag) .

minikube:
	minikube start
	eval $$(minikube docker-env)
	minikube addons enable ingress
