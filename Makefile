OS   := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

KUBERNETES_VERSION         := 1.21.1
ISTIO_VERSION              := 1.11.0
KIND_VERSION               := 0.11.1
BUF_VERSION                := 0.44.0
PROTOC_GEN_GO_VERSION      := 1.27.1
PROTOC_GEN_GO_GRPC_VERSION := 1.1.0

BIN_DIR := $(shell pwd)/bin

KUBECTL                 := $(abspath $(BIN_DIR)/kubectl)
ISTIOCTL                := $(abspath $(BIN_DIR)/istioctl)
KIND                    := $(abspath $(BIN_DIR)/kind)
BUF                     := $(abspath $(BIN_DIR)/buf)
PROTOC_GEN_GO           := $(abspath $(BIN_DIR)/protoc-gen-go)
PROTOC_GEN_GO_GRPC      := $(abspath $(BIN_DIR)/protoc-gen-go-grpc)
PROTOC_GEN_GRPC_GATEWAY := $(abspath $(BIN_DIR)/protoc-gen-grpc-gateway)

KIND_CLUSTER_NAME := mercari-microservices-example

KUBECTL_CMD := KUBECONFIG=./.kubeconfig $(KUBECTL)
KIND_CMD    := $(KIND) --name $(KIND_CLUSTER_NAME) --kubeconfig ./.kubeconfig

kubectl: $(KUBECTL)
$(KUBECTL):
	curl -Lso $(KUBECTL) https://storage.googleapis.com/kubernetes-release/release/v$(KUBERNETES_VERSION)/bin/$(OS)/$(ARCH)/kubectl
	chmod +x $(KUBECTL)

istioctl: $(ISTIOCTL)
$(ISTIOCTL):
ifeq ($(OS)-$(ARCH), darwin-amd64)
	curl -sSL "https://storage.googleapis.com/istio-release/releases/$(ISTIO_VERSION)/istioctl-$(ISTIO_VERSION)-osx.tar.gz" | tar -C $(BIN_DIR) -xzv istioctl
else ifeq ($(OS)-$(ARCH), darwin-arm64)
	curl -sSL "https://storage.googleapis.com/istio-release/releases/$(ISTIO_VERSION)/istioctl-$(ISTIO_VERSION)-osx-arm64.tar.gz" | tar -C $(BIN_DIR) -xzv istioctl
else
	curl -sSL "https://storage.googleapis.com/istio-release/releases/$(ISTIO_VERSION)/istioctl-$(ISTIO_VERSION)-$(OS)-$(ARCH).tar.gz" | tar -C $(BIN_DIR) -xzv istioctl
endif

kind: $(KIND)
$(KIND):
	curl -Lso $(KIND) https://github.com/kubernetes-sigs/kind/releases/download/v$(KIND_VERSION)/kind-$(OS)-$(ARCH)
	chmod +x $(KIND)

buf: $(BUF)
$(BUF):
	curl -sSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" -o $(BUF) && chmod +x $(BUF)

protoc-gen-go: $(PROTOC_GEN_GO)
$(PROTOC_GEN_GO):
	curl -sSL https://github.com/protocolbuffers/protobuf-go/releases/download/v$(PROTOC_GEN_GO_VERSION)/protoc-gen-go.v$(PROTOC_GEN_GO_VERSION).$(OS).$(ARCH).tar.gz | tar -C $(BIN_DIR) -xzv protoc-gen-go

protoc-gen-go-grpc: $(PROTOC_GEN_GO_GRPC)
$(PROTOC_GEN_GO_GRPC):
	curl -sSL https://github.com/grpc/grpc-go/releases/download/cmd%2Fprotoc-gen-go-grpc%2Fv$(PROTOC_GEN_GO_GRPC_VERSION)/protoc-gen-go-grpc.v$(PROTOC_GEN_GO_GRPC_VERSION).$(OS).$(ARCH).tar.gz | tar -C $(BIN_DIR) -xzv ./protoc-gen-go-grpc

protoc-gen-grpc-gateway: $(PROTOC_GEN_GRPC_GATEWAY)
$(PROTOC_GEN_GRPC_GATEWAY):
	cd ./tools && go build -o ../bin/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

.PHONY: cluster
cluster: $(KIND) $(KUBECTL) $(ISTIOCTL)
	$(KIND_CMD) delete cluster
	$(KIND_CMD) create cluster --image kindest/node:v${KUBERNETES_VERSION} --config ./kind.yaml
	./script/istioctl install --set meshConfig.defaultConfig.tracing.zipkin.address=jaeger.jaeger.svc.cluster.local:9411 -y
	$(KUBECTL_CMD) apply --filename ./platform/ingress-nginx/ingress-nginx.yaml
	$(KUBECTL_CMD) wait \
		--namespace ingress-nginx \
		--for=condition=ready pod \
		--selector=app.kubernetes.io/component=controller \
		--timeout=90s
	$(KUBECTL_CMD) apply --filename ./platform/kiali/kiali.yaml
	sleep 5
	$(KUBECTL_CMD) apply --filename ./platform/kiali/dashboard.yaml
	$(KUBECTL_CMD) apply --kustomize ./platform/jaeger
	make db
	make gateway
	make authority
	make catalog
	make customer
	make item

.PHONY: db
db:
	$(KUBECTL_CMD) delete deploy -n db --ignore-not-found app
	docker build -t mercari/mercari-microservices-example/db:latest --file ./platform/db/Dockerfile .
	$(KIND) load docker-image mercari/mercari-microservices-example/db:latest --name $(KIND_CLUSTER_NAME)
	$(KUBECTL_CMD) apply --filename ./platform/db/deployment.yaml

.PHONY: gateway
gateway:
	$(KUBECTL_CMD) delete deploy -n gateway --ignore-not-found app
	docker build -t mercari/mercari-microservices-example/gateway:latest --file ./services/gateway/Dockerfile .
	$(KIND) load docker-image mercari/mercari-microservices-example/gateway:latest --name $(KIND_CLUSTER_NAME)
	$(KUBECTL_CMD) apply --filename ./services/gateway/deployment.yaml

.PHONY: authority
authority:
	$(KUBECTL_CMD) delete deploy -n authority --ignore-not-found app
	docker build -t mercari/mercari-microservices-example/authority:latest --file ./services/authority/Dockerfile .
	$(KIND) load docker-image mercari/mercari-microservices-example/authority:latest --name $(KIND_CLUSTER_NAME)
	$(KUBECTL_CMD) apply --filename ./services/authority/deployment.yaml

.PHONY: catalog
catalog:
	$(KUBECTL_CMD) delete deploy -n catalog --ignore-not-found app
	docker build -t mercari/mercari-microservices-example/catalog:latest --file ./services/catalog/Dockerfile .
	$(KIND) load docker-image mercari/mercari-microservices-example/catalog:latest --name $(KIND_CLUSTER_NAME)
	$(KUBECTL_CMD) apply --filename ./services/catalog/deployment.yaml

.PHONY: customer
customer:
	$(KUBECTL_CMD) delete deploy -n customer --ignore-not-found app
	docker build -t mercari/mercari-microservices-example/customer:latest --file ./services/customer/Dockerfile .
	$(KIND) load docker-image mercari/mercari-microservices-example/customer:latest --name $(KIND_CLUSTER_NAME)
	$(KUBECTL_CMD) apply --filename ./services/customer/deployment.yaml

.PHONY: item
item:
	$(KUBECTL_CMD) delete deploy -n item --ignore-not-found app
	docker build -t mercari/mercari-microservices-example/item:latest --file ./services/item/Dockerfile .
	$(KIND) load docker-image mercari/mercari-microservices-example/item:latest --name $(KIND_CLUSTER_NAME)
	$(KUBECTL_CMD) apply --filename ./services/item/deployment.yaml

.PHONY: pb
pb:
	docker run \
		--volume "$(shell pwd):/go/src/github.com/mercari/mercari-microservices-example" \
		--workdir /go/src/github.com/mercari/mercari-microservices-example \
		--rm \
		golang:1.17.1-bullseye \
		make gen-proto

.PHONY: gen-proto
gen-proto: $(BUF) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) $(PROTOC_GEN_GRPC_GATEWAY)
	$(BUF) generate \
		--path ./services/ \
		--path ./platform/db/

.PHONY: clean
clean:
	$(KIND_CMD) delete cluster
	rm -f $(BIN_DIR)/*
