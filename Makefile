boiler-plate: test-build-bake docker-push

test-build-bake:
	docker build -t docker.io/mauricio1998/authentication-service . -f build/Dockerfile

run-tests:
	go test ./... -coverprofile cover.out -tags=test && go tool cover -html=cover.out

run-db:
	docker-compose -f deployments/db-docker-compose.yml up -d

docker-push:
	docker push docker.io/mauricio1998/authentication-service