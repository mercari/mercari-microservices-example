OS   := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

BIN_DIR := $(shell pwd)/bin

KUBERNETES_VERSION         := 1.20.2
ISTIO_VERSION              := 1.10.0-alpha.0
KIND_VERSION               := 0.10.0
BUF_VERSION                := 0.39.1
PROTOC_GEN_GO_VERSION      := 1.25.0
PROTOC_GEN_GO_GRPC_VERSION := 1.1.0

KUBECTL            := $(abspath $(BIN_DIR)/kubectl)
ISTIOCTL           := $(abspath $(BIN_DIR)/istioctl)
KIND               := $(abspath $(BIN_DIR)/kind)
BUF                := $(abspath $(BIN_DIR)/buf)
PROTOC_GEN_GO      := $(abspath $(BIN_DIR)/protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(abspath $(BIN_DIR)/protoc-gen-go-grpc)

KIND_CLUSTER_NAME := mercari-go-conference-2021-spring-office-hour

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

.PHONY: cluster
cluster: $(KIND) $(KUBECTL) $(ISTIOCTL)
	$(KIND_CMD) delete cluster
	$(KIND_CMD) create cluster --image kindest/node:v${KUBERNETES_VERSION}
	$(ISTIOCTL) install --context kind-$(KIND_CLUSTER_NAME) --kubeconfig ./.kubeconfig -y
	make images
	$(KUBECTL_CMD) apply --filename ./services/payment/deployment.yaml
	$(KUBECTL_CMD) apply --filename ./services/balance/deployment.yaml
	$(KUBECTL_CMD) apply --kustomize ./services/jaeger

.PHONY: images
images:
	docker build -t mercari/go-conference-2021-spring-office-hour/payment:latest --file ./services/payment/Dockerfile .
	$(KIND) load docker-image mercari/go-conference-2021-spring-office-hour/payment:latest --name $(KIND_CLUSTER_NAME)
	docker build -t mercari/go-conference-2021-spring-office-hour/balance:latest --file ./services/balance/Dockerfile .
	$(KIND) load docker-image mercari/go-conference-2021-spring-office-hour/balance:latest --name $(KIND_CLUSTER_NAME)

.PHONY: gen-proto
gen-proto: $(BUF) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC)
	$(BUF) generate --path ./services/

.PHONY: clean
clean:
	$(KIND_CMD) delete cluster
	rm -f $(BIN_DIR)/*
