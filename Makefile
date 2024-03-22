test-build-bake:
	docker build -t auth-service . -f build/Dockerfile

run-tests:
	go test ./... -coverprofile cover.out -tags=test && go tool cover -html=cover.out

run-db:
	docker-compose -f deployments/db-docker-compose.yml up -d