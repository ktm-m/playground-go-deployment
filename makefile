build:
	docker build -t montheankulk/playground-go-fr:latest .

push:
	docker push montheankulk/playground-go-fr:latest

deploy:
	kubectl apply -k ./deployment/

delete-namespace:
	kubectl delete namespace app

pods:
	kubectl get pods -n app

forward-backend-port:
	kubectl port-forward svc/backend 8080:8080 -n app

forward-nginx-port:
	kubectl port-forward svc/nginx 8081:80 -n app