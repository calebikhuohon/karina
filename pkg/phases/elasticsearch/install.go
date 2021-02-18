package elasticsearch

import (
	"github.com/flanksource/karina/pkg/constants"
	"github.com/flanksource/karina/pkg/platform"
	"github.com/flanksource/karina/pkg/types"
	"github.com/pkg/errors"
)

const Namespace = "eck"

func Deploy(p *platform.Platform) error {
	if p.Elasticsearch == nil || p.Elasticsearch.Disabled {
		p.Elasticsearch = &types.Elasticsearch{Mem: &types.Memory{Limits: "1Gi", Requests: "1Gi"}, Persistence: &types.Persistence{Enabled: true}}
		if err := p.DeleteSpecs(Namespace, "elasticsearch.yaml"); err != nil {
			p.Warnf("failed to delete specs: %v", err)
		}
		return nil
	}

	if err := p.CreateOrUpdateNamespace(Namespace, nil, nil); err != nil {
		return err
	}
	if p.Elasticsearch.Mem == nil {
		p.Elasticsearch.Mem = &types.Memory{Limits: "3G", Requests: "2G"}
	}
	if p.Elasticsearch.Replicas == 0 {
		p.Elasticsearch.Replicas = 3
	}

	if !p.LogsExporter.IsDisabled() {
		if err := p.ApplySpecs(constants.PlatformSystem, "logs-exporter.yaml"); err != nil {
			return errors.Wrap(err, "failed to deploy logs exporter")
		}

		if err := p.ApplySpecs(constants.PlatformSystem, "logs-exporter-elasticlogs.yaml"); err != nil {
			return errors.Wrap(err, "failed to deploy logs exporter elasticlogs")
		}
	}

	return p.ApplySpecs(Namespace, "elasticsearch.yaml")
}
