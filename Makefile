.PHONY: k8s
k8s:
	kubectl apply -f k8s/redis.yaml
	kubectl apply -f k8s/ws-server.yaml
	#kubectl apply -f k8s/ws-client.yaml
	kubectl apply -f k8s/kong.yaml

clean:
	kubectl delete deployment ws-client