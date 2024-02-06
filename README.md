# Passt KubeVirt network binding plugin

Passt network binding plugin configures VMs Passt interface using Kubevirts
hook sidecar interface.

It will be used by Kubevirt to offload Passt networking configuration.

> _NOTE_:
> Passt network binding is supported for pod network interfaces only.

# Usage

Deploy the `passt` binding plugin binary on the cluster:

```sh
TODO
```

Register the `passt` binding plugin with its sidecar image:

```yaml
apiVersion: kubevirt.io/v1
kind: KubeVirt
metadata:
  name: kubevirt
  namespace: kubevirt
spec:
  configuration:
    network:
      binding:
        passt:
          sidecarImage: registry:5000/kubevirt/network-passt-binding:devel
  ...
```

In the VM spec, set interface to use `passt` binding plugin:

```yaml
apiVersion: kubevirt.io/v1
kind: VirtualMachineInstance
metadata:
  name: vmi-passt
spec:
  domain:
    devices:
      interfaces:
      - name: passt
        binding:
          name: passt
  ...
  networks:
  - name: passt-net
    pod: {}
  ...
```

# Development

See [DEVELOPMENT.md](DEVELOPMENT.md) to find some basic commands to interact
with the project. To learn more, read the code. A good place to start would be
TODO.

# License

Software of this project is distributed under the terms of the Apache license
version 2. See [LICENSE](LICENSE) for details.

# Changelog

Read the [CHANGELOG.md](CHANGELOG.md) to learn about changes introduced in each
release.
