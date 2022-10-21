OS   := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

KUBERNETES_VERSION         := 1.25.0
ISTIO_VERSION              := 1.15.0
KIND_VERSION               := 0.15.0
INGRESS_NGINX_VERSION      := 1.3.0
BUF_VERSION                := 1.9.0

BIN_DIR := $(shell pwd)/bin

KUBECTL                 := $(abspath $(BIN_DIR)/kubectl)
ISTIOCTL                := $(abspath $(BIN_DIR)/istioctl)
KIND                    := $(abspath $(BIN_DIR)/kind)
BUF                     := $(abspath $(BIN_DIR)/buf)

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
	$(KUBECTL_CMD) apply --filename ./platform/kiali/prometheus.yaml
	sleep 5
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
		golang:1.19-bullseye \
		make gen-proto

.PHONY: gen-proto
gen-proto: $(BUF)
	$(BUF) generate \
		--path ./services/ \
		--path ./platform/db/

.PHONY: ingress-nginx-manifest
ingress-nginx-manifest:
	@curl -s https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v$(INGRESS_NGINX_VERSION)/deploy/static/provider/kind/deploy.yaml > ./platform/ingress-nginx/ingress-nginx.yaml

.PHONY: istio-addons-manifest
istio-addons-manifest:
	@curl -s https://raw.githubusercontent.com/istio/istio/$(ISTIO_VERSION)/samples/addons/kiali.yaml      > ./platform/kiali/kiali.yaml
	@curl -s https://raw.githubusercontent.com/istio/istio/$(ISTIO_VERSION)/samples/addons/prometheus.yaml > ./platform/kiali/prometheus.yaml

.PHONY: clean
clean:
	$(KIND_CMD) delete cluster
	rm -f $(BIN_DIR)/*
