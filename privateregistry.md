# Setting up Kubernetes Secrets for Container Registry Authentication (Image Scanning)

Scanning images from container registries which require authentication can be accomplished by following the steps below.

1. Generate base64 encoded strings of your registry, username and password:

    ```bash
    echo -n 'myrepo.azurecr.io' | base64
    echo -n 'myusername' | base64
    echo -n 'mypassword' | base64
    ```

    Replace 'myrepo.azurecr.io', 'myusername' and 'mypassword' placeholders with your actual registry, username and password respectively.

2. Create a Kubernetes secret with a named prefix `seclogic-registry-auth` in the `seclogic` namespace.

    Each entry in the secret should contain the following fields, with the base64 encoded strings from the previous step:
    - `registry`: Registry name without the http/https prefix
    - `username`: Username
    - `password`: Password


    ```yaml
    apiVersion: v1
    kind: Secret
    metadata:
        name: mysecret
        namespace: seclogic
        labels: 
          seclogic.io/registry=creds
    type: Opaque
    data: // stringData
        registry: bXlyZXBvLmF6dXJlY3IuaW8=
        username: bXl1c2VybmFtZQ==
        password: bXlwYXNzd29yZA==
    ```

    Save this content into a file (e.g., `seclogic-registry-auth.yaml`) and apply it using `kubectl`:

    ```bash
    kubectl apply -f seclogic-registry-auth.yaml
    ```

3. That's it. For each image scan, Kubescape will look for all `seclogic-registry-auth` secrets under the `seclogic` namespace.