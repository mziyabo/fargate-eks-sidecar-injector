# fargate-eks-sidecar-injector
Kubernetes mutating webhook for injecting sidecars into AWS Fargate pods.

Works by checking fargate pod annotation `eks.amazonaws.com/fargate-profile` and then injecting sidecars from a `fargate-injector-sidecar-config` ConfigMap.

## Install
To build and run the webhook from source:

```bash
IMAGE_NAME="fargate-eks-sidecar-injector"
TAG="latest"

docker build . -t $IMAGE_NAME:$TAG

helm install helm/chart/fargate-sidecar-injector  \
    --values helm/chart/fargate-sidecar-injector/values.yaml \
    --generate-name \
    --set image.repository=${IMAGE_NAME} \
    --set image.tag=${TAG}
```

This webhook needs to run after `0500-amazon-eks-fargate-mutation.amazonaws.com` therefore if you specify `.Values.nameOverride` make sure to use a name lexicographically greater than the amazon webhook.
## Usage
After installing, To trigger sidecar injection, add a ConfigMap named `fargate-injector-sidecar-config` in the webhook namespace with the below format:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: fargate-injector-sidecar-config
data: 
  default: |-  # <- Fargate Profile
    spec:
      volumes:
      - name: log
        emptyDir: {}
      containers:
      - image: busybox
        name: side-man
        args:
        - /bin/sh
        - -c
        - >
          while true; do
            echo "$(date) INFO hello from main-container" >> /var/log/myapp.log ;
            sleep 1;
          done
        volumeMounts:
        - name: log
          mountPath: /var/log
```

> Note the fargate-profile is used to configure which sidecars are injected. 
  The spec for the containers under the fargate profile is exactly the same as podSpec.Containers - which is the same for Volumes.

Triggers on Pod `CREATE` and `UPDATE` operations.

## Release Notes
WIP, Contributions Welcome

## License
[Apache License, Version 2.0](./LICENSE)

## Known Issues