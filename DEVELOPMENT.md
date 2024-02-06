# Development

Build a container image with the CNI plugin:

```sh
podman build -f Containerfile.cni
```

Build and publish CNI container:

```sh
podman build -f Containerfile.cni -t quay.io/phoracek/passt-kubevirt-network-binding-plugin-cni:latest
podman push quay.io/phoracek/passt-kubevirt-network-binding-plugin-cni:latest
```
