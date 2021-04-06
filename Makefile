OS   := $(shell uname | awk '{print tolower($$0)}')
ARCH := $(shell case $$(uname -m) in (x86_64) echo -n amd64 ;; (aarch64) echo -n arm64 ;; (*) echo -n $$(uname -m) ;; esac)

BIN_DIR := $(shell pwd)/bin

KUBERNETES_VERSION := 1.20.2
ISTIO_VERSION      := 1.9.2
KIND_VERSION       := 0.10.0

KUBECTL  := $(abspath $(BIN_DIR)/kubectl)
ISTIOCTL := $(abspath $(BIN_DIR)/istioctl)
KIND     := $(abspath $(BIN_DIR)/kind)

KIND_CLUSTER_NAME := mercari-go-conference-2021-spring-office-hour

kubectl: $(KUBECTL)
$(KUBECTL):
	curl -Lso $(KUBECTL) https://storage.googleapis.com/kubernetes-release/release/v$(KUBERNETES_VERSION)/bin/$(OS)/$(ARCH)/kubectl
	chmod +x $(KUBECTL)

istioctl: $(ISTIOCTL)
$(ISTIOCTL):
	curl -sSL "https://storage.googleapis.com/istio-release/releases/${ISTIO_VERSION}/istioctl-${ISTIO_VERSION}-${OS}-${ARCH}.tar.gz" | tar -C $(BIN_DIR) -xzv istioctl

kind: $(KIND)
$(KIND):
	curl -Lso $(KIND) https://github.com/kubernetes-sigs/kind/releases/download/v$(KIND_VERSION)/kind-$(OS)-$(ARCH)
	chmod +x $(KIND)

.PHONY: cluster
cluster: $(KIND) $(KUBECTL) $(ISTIOCTL)
	$(KIND) delete cluster --name $(KIND_CLUSTER_NAME)
	$(KIND) create cluster --name $(KIND_CLUSTER_NAME) --image kindest/node:v${KUBERNETES_VERSION}
	$(ISTIOCTL) install -y
	make images
	$(KUBECTL) apply --filename ./services/payment/deployment.yaml
	$(KUBECTL) apply --filename ./services/balance/deployment.yaml

.PHONY: images
images:
	docker build -t mercari/go-conference-2021-spring-office-hour/payment:latest --file ./services/payment/Dockerfile .
	$(KIND) load docker-image mercari/go-conference-2021-spring-office-hour/payment:latest --name $(KIND_CLUSTER_NAME)
	docker build -t mercari/go-conference-2021-spring-office-hour/balance:latest --file ./services/balance/Dockerfile .
	$(KIND) load docker-image mercari/go-conference-2021-spring-office-hour/balance:latest --name $(KIND_CLUSTER_NAME)

.PHONY: clean
clean:
	$(KIND) delete cluster --name $(KIND_CLUSTER_NAME)
	rm -f $(BIN_DIR)/*
