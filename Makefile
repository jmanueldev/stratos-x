run-api:
	go run api/main.go

run-agent:
	cd node-agent && cargo run

submit:
	python sdk/client.py missions/example.yaml

run-dashboard:
	cd web-dashboard && npm start
