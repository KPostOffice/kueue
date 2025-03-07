PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
ARTIFACTS ?= $(PROJECT_DIR)/bin

ifeq ($(shell uname),Darwin)
    GOFLAGS ?= -ldflags=-linkmode=internal
endif

ifeq (,$(shell go env GOBIN))
	GOBIN=$(shell go env GOPATH)/bin
else
	GOBIN=$(shell go env GOBIN)
endif

GO_CMD ?= go
GO_TEST_FLAGS ?= -race
version_pkg = sigs.k8s.io/kueue/pkg/version
LD_FLAGS += -X '$(version_pkg).GitVersion=$(GIT_TAG)'
LD_FLAGS += -X '$(version_pkg).GitCommit=$(shell git rev-parse HEAD)'

# Use go.mod go version as source.
KUSTOMIZE_OCP_VERSION ?= $(shell $(GO_CMD) list -m -mod=mod -f '{{.Version}}' sigs.k8s.io/kustomize/kustomize/v5)
GINKGO_OCP_VERSION ?= $(shell $(GO_CMD) list -m -mod=mod -f '{{.Version}}' github.com/onsi/ginkgo/v2)
YQ_OCP_VERSION ?= $(shell $(GO_CMD) list -m -mod=mod -f '{{.Version}}' github.com/mikefarah/yq/v4)
ENVTEST_OCP_VERSION ?= $(shell $(GO_CMD) list -m -mod=mod -f '{{.Version}}' sigs.k8s.io/controller-runtime/tools/setup-envtest)
GOTESTSUM_OCP_VERSION ?= $(shell $(GO_CMD) list -m -mod=mod -f '{{.Version}}' gotest.tools/gotestsum)

KUSTOMIZE = $(PROJECT_DIR)/bin/kustomize
.PHONY: kustomize-ocp
kustomize-ocp: ## Download kustomize locally if necessary.
	@GOBIN=$(PROJECT_DIR)/bin GO111MODULE=on $(GO_CMD) install -mod=mod sigs.k8s.io/kustomize/kustomize/v5@$(KUSTOMIZE_OCP_VERSION)

GINKGO = $(PROJECT_DIR)/bin/ginkgo
.PHONY: ginkgo-ocp
ginkgo-ocp: ## Download ginkgo locally if necessary.
	@GOBIN=$(PROJECT_DIR)/bin GO111MODULE=on $(GO_CMD) install -mod=mod github.com/onsi/ginkgo/v2/ginkgo@$(GINKGO_OCP_VERSION)

ENVTEST = $(PROJECT_DIR)/bin/setup-envtest
.PHONY: envtest
envtest-ocp: ## Download envtest-setup locally if necessary.
	@GOBIN=$(PROJECT_DIR)/bin GO111MODULE=on $(GO_CMD) install sigs.k8s.io/controller-runtime/tools/setup-envtest@$(ENVTEST_OCP_VERSION)

GOTESTSUM = $(PROJECT_DIR)/bin/gotestsum
.PHONY: gotestsum
gotestsum-ocp: ## Download gotestsum locally if necessary.
	@GOBIN=$(PROJECT_DIR)/bin GO111MODULE=on $(GO_CMD) install gotest.tools/gotestsum@$(GOTESTSUM_OCP_VERSION)

YQ = $(PROJECT_DIR)/bin/yq
.PHONY: yq-ocp
yq-ocp: ## Download yq locally if necessary.
	@GOBIN=$(PROJECT_DIR)/bin GO111MODULE=on $(GO_CMD) install -mod=mod github.com/mikefarah/yq/v4@$(YQ_OCP_VERSION)

.PHONY: test-ocp
test-ocp: gotestsum-ocp ## Run tests.
	$(GOTESTSUM) --junitfile $(ARTIFACTS)/junit.xml -- $(GOFLAGS) $(GO_TEST_FLAGS) $(shell $(GO_CMD) list ./... | grep -v '/test/') -coverpkg=./... -coverprofile $(ARTIFACTS)/cover.out

.PHONY: test-integration-ocp
test-integration: envtest-cop ginkgo-ocp dep-crds kueuect-ocp ginkgo-top-ocp ## Run integration tests for all singlecluster suites.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) -p path)" \
	PROJECT_DIR=$(PROJECT_DIR)/ \
	KUEUE_BIN=$(PROJECT_DIR)/bin \
	ENVTEST_K8S_VERSION=$(ENVTEST_K8S_VERSION) \
	API_LOG_LEVEL=$(INTEGRATION_API_LOG_LEVEL) \
	$(GINKGO) $(INTEGRATION_FILTERS) $(GINKGO_ARGS) -procs=$(INTEGRATION_NPROCS) --race --junit-report=junit.xml --json-report=integration.json --output-dir=$(ARTIFACTS) -v $(INTEGRATION_TARGET)
	$(PROJECT_DIR)/bin/ginkgo-top -i $(ARTIFACTS)/integration.json > $(ARTIFACTS)/integration-top.yaml

.PHONY: test-e2e-ocp
test-e2e-ocp: kustomize-ocp ginkgo-ocp yq-ocp kueuectl-ocp ginkgo-top-ocp run-test-e2e-ocp-singlecluster
run-test-e2e-ocp-singlecluster:
	@echo "Running e2e tests on OpenShift cluster ($(shell oc whoami --show-server))"
	ARTIFACTS="$(ARTIFACTS)/$@" IMAGE_TAG=$(IMAGE_TAG) GINKGO_ARGS="$(GINKGO_ARGS)" \
	E2E_TARGET_FOLDER="singlecluster" \
	./hack/e2e-test-ocp.sh
	$(PROJECT_DIR)/bin/ginkgo-top -i $(ARTIFACTS)/$@/e2e.json > $(ARTIFACTS)/$@/e2e-top.yaml

.PHONY: ginkgo-top-ocp
ginkgo-top-ocp:
	$(GO_BUILD_ENV) $(GO_CMD) build -ldflags="$(LD_FLAGS)" -o $(PROJECT_DIR)/bin/ginkgo-top ./pkg/openshift/ginkgo-top

.PHONY: kueuectl-ocp
kueuectl-ocp:
	CGO_ENABLED=$(CGO_ENABLED) $(GO_BUILD_ENV) $(GO_CMD) build -ldflags="$(LD_FLAGS)" -o $(PROJECT_DIR)/bin/kubectl-kueue cmd/kueuectl/main.go
