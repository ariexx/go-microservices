FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
ORDER_BINARY=orderApp
PRODUCT_BINARY=productApp
PAYMENT_BINARY=paymentApp

PRODUCT_CLIENT=productClient

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_proto build_order build_product build_broker build_payment #build_auth
	@echo "Stopping docker images (if running...)"
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${PRODUCT_CLIENT} ./cmd/client
	@echo "Done!"

build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

build_order:
	@echo "Building order binary..."
	cd ../order-service && env GOOS=linux CGO_ENABLED=0 go build -o ${ORDER_BINARY} ./cmd/api
	@echo "Done!"

build_product:
	@echo "Building product binary..."
	cd ../product-service && env GOOS=linux CGO_ENABLED=0 go build -o ${PRODUCT_BINARY} ./cmd/api
	@echo "Done!"

build_payment:
	@echo "Building product binary..."
	cd ../payment-service && env GOOS=linux CGO_ENABLED=0 go build -o ${PAYMENT_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd ../front-end && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"

order_proto:
	rm -rf ../order-service/pb/*.go
	protoc --proto_path=../order-service/proto --go_out=../order-service/pb --go_opt=paths=source_relative \
        --go-grpc_out=../order-service/pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=../order-service/pb --grpc-gateway_opt=paths=source_relative \
        ../order-service/proto/*.proto

product_proto:
	rm -rf ../product-service/pb/*.go
	protoc --proto_path=../product-service/proto --go_out=../product-service/pb --go_opt=paths=source_relative \
        --go-grpc_out=../product-service/pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=../product-service/pb --grpc-gateway_opt=paths=source_relative \
        ../product-service/proto/*.proto

payment_proto:
	rm -rf ../payment-service/pb/*.go
	protoc --proto_path=../payment-service/proto --go_out=../payment-service/pb --go_opt=paths=source_relative \
		--go-grpc_out=../payment-service/pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=../payment-service/pb --grpc-gateway_opt=paths=source_relative \
		../payment-service/proto/*.proto

email_proto:
	rm -rf ../email-service/pb/*.go
	protoc --proto_path=../email-service/proto --go_out=../order-service/pb --go_opt=paths=source_relative \
		--go-grpc_out=../order-service/pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=../order-service/pb --grpc-gateway_opt=paths=source_relative \
		../email-service/proto/*.proto

order_email:
	make order_proto && make email_proto

build_proto: product_proto payment_proto order_email