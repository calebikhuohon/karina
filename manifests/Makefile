TEMP_DIR = $(shell mktemp -d)

TARGET_CRD_FILES := $(patsubst upstream/crd/%.yaml,crd/%.yaml,$(wildcard upstream/crd/*.yaml))
TARGET_MANIFEST_FILES := $(patsubst upstream/%.yaml,%.yaml,$(wildcard upstream/*.yaml))

build: build-crds build-manifests
build-crds: $(TARGET_CRD_FILES)
build-manifests: $(TARGET_MANIFEST_FILES)

define BUILD_MANIFEST
	@echo "Generating $@..."
	@export TEMP_DIR=$(TEMP_DIR) && \
	cp upstream/$@ $$TEMP_DIR/kustomization.yaml && \
	out=$$(kustomize build $$TEMP_DIR) && \
	echo "# This file was generated by running 'make $@' in manifests directory\n$$out" > $@ && \
	rm -rf $$TEMP_DIR
endef

crd/%.yaml: upstream/crd/%.yaml
	$(BUILD_MANIFEST)

%.yaml: upstream/%.yaml
	$(BUILD_MANIFEST)


.PHONY: build build-crds build-manifests
