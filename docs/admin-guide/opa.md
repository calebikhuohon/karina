`karina.yml`
#OPA Gatekeeper

```yaml
gatekeeper:
  version: v3.3.0
  constraints: /path/to/constraints/folder
  templates: /path/to/templates/folder
  whitelistNamespaces:
    - unpolicied-namespace
```

---
**NOTE**

Support for installing vanilla OPA via karina is deprecated.  Use at own risk.

---

Deploying using :

```bash
karina deploy opa -c karina.yml

```

See https://open-policy-agent.github.io/gatekeeper/website/docs/howto/ for general gatekeeper information.

By default, karina deploys gatekeeper with a selection of the default constrainttemplates found in the [gatekeeper example library](https://github.com/open-policy-agent/gatekeeper-library).  These include:

- K8sAllowedRepos 
- K8sBlockNodePort
- K8sContainerLimits
- K8sContainerRatios
- K8sDenyAll
- K8sDisallowedTags
- K8sRequiredAnnotations
- K8sRequiredLabels
- K8sRequiredProbes
- K8sUniqueIngressHost

Additional templates and constraints can be referred to in the deployment configuration.

To test new constraints, configure them with

```yaml
spec:
  enforcementaction: dryrun
```

and monitor the number of violations that would be enforced using:

```bash
# high level overview
karina status violations -c gatekeeper.yaml
# specific violation
kubectl describe constraint $CONSTRAINT
```

https://play.openpolicyagent.org/ is a useful tool for debugging rego operation
