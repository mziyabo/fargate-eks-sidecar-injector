# eks-fargate-sidecar-injector
:warning: WIP

Kubernetes mutating webhook for conditionally injecting sidecars into AWS Fargate pods.


## Install
```bash
IMAGE_NAME="eks-fargate-sidecar-injector-webhook"
TAG="latest"

docker build . -t $IMAGE_NAME:$TAG

helm install helm/chart/fargate-sidecar-injector  --values helm/chart/fargate-sidecar-injector/values.yaml --generate-name --set image.repository=${IMAGE_NAME} --set image.tag=${TAG}
```


This webhook needs to run after `0500-amazon-eks-fargate-mutation.amazonaws.com` therefore if you specify `.Values.nameOverride` make sure to use a name lexicographically greater than the amazon webhook.
## Usage
To trigger sidecar injection add a ConfigMap named `fargate-injector-sidecar-config` in the webhook namespace with the below format:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: fargate-injector-sidecar-config
data: 
  default: |-  # <- Fargate Profile
    spec:
      containers:
      - name: datadog
        image: datadog-operator:v1
        ports:
        - containerPort: 80
        volumeMounts:
        - name: data
          mountPath: /path/in/container
      volumes:
      - name: data
        configMap:
          name: myconfig
```

> Note the fargate profile is used to configure which sidecars are injected. 
  The spec for the containers under the fargate profile is exactly the same as podSpec.Containers - which is the same for Volumes.

## Release Notes
WIP, Contributions Welcome

## License
[Apache License, Version 2.0](./LICENSE)

## Known Issues
- `helm uninstall` does not delete the MutatingWebhookConfiguration resource since it's deployed via a helm post-install-hook. This means you have to delete it separately from the helm release - [see here](https://helm.sh/docs/topics/charts_hooks/#hook-resources-are-not-managed-with-corresponding-releases)