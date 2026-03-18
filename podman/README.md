# Podman

[Podman](https://podman.io/) is a free container tools, an alternative to docker.

## VSCode

To enable development inside container in VSCode, you need [DevContainers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension (ID: `ms-vscode-remote.remote-containers`).

If you work with Dockerfile, it's also recommended to install [Container Tools](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-containers) extension (ID: `ms-azuretools.vscode-containers`)


Set the VSCode to use podman by pressing `Ctrl + Shift + P` and choose `Preferences: Open Settings (UI)`.

Change `Dev > Containers: Docker Compose Path` value to `podman compose`.

Also change `Dev > Containers: Docker Path` to `podman`.

