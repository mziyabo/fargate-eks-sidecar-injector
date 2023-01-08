CRD - SideCarConfig
MWH - Detect Fargate - Read Label (Optional) - Inject Container

https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/



https://google.github.io/styleguide/go/best-practices


operator-sdk create api --group cache --version v1alpha1 --kind FargateSidecarConfig --resource --controller


operator-sdk create webhook --group cache --version v1alpha1 --kind FargateSidecarConfig --defaulting --programmatic-validation
operator-sdk create webhook --group cache --version v1alpha1 --kind Pod --defaulting

kubebuilder create webhook --group batch --version v1 --kind Pod --defaulting --programmatic-validation

-----------


https://bmiguel-teixeira.medium.com/helm-certify-your-path-to-victory-6c974a3d9f15


# Webhook
cat EOF << | kubectl apply -f -
apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  annotations:
    cert-manager.io/inject-ca-from: default/client
    meta.helm.sh/release-name: fargate-sidecar-injector
    meta.helm.sh/release-namespace: default
  creationTimestamp: "2022-06-25T19:14:12Z"
  generation: 1
  labels:
    app.kubernetes.io/managed-by: Helm
  name: fargate-sidecar-injector
  resourceVersion: "934881"
  uid: b665d60c-6296-4c7b-a438-50a13ba93db1
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    caBundle: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURjVENDQWxtZ0F3SUJBZ0lVSjlIYzBGS210VHJRWHhOaE5sQzhKQy9VbTNNd0RRWUpLb1pJaHZjTkFRRUwKQlFBd2dZb3hDekFKQmdOVkJBWVRBbHBCTVJrd0Z3WURWUVFJRXhCWFpYTjBaWEp1SUZCeWIzWnBibU5sTVJJdwpFQVlEVlFRSEV3bERZWEJsSUZSdmQyNHhFekFSQmdOVkJBb1RDbTF5WkdWc2FYWmxjbmt4RkRBU0JnTlZCQXNUCkMwVnVaMmx1WldWeWFXNW5NU0V3SHdZRFZRUURFeGhtYjI5a0xXbHRZV2RsTFd4aFltVnNMWGRsWW1odmIyc3cKSGhjTk1qSXdOakl6TVRneU1qQXdXaGNOTWpNd05qSXpNVGd5TWpBd1dqQWpNU0V3SHdZRFZRUURFeGhtYjI5awpMV2x0WVdkbExXeGhZbVZzTFhkbFltaHZiMnN3V1RBVEJnY3Foa2pPUFFJQkJnZ3Foa2pPUFFNQkJ3TkNBQVNaCkxBMUZibGlmblhtTVFYWFBPZTA4cW82dWZJQTVhOU9KdHVIZmRXdDVHZWJjQmZJSWJIanc3cldqaDIwOUticEkKK09pYWFNWW1lSThPUnd3SVJ5MjlvNEgvTUlIOE1BNEdBMVVkRHdFQi93UUVBd0lGb0RBZEJnTlZIU1VFRmpBVQpCZ2dyQmdFRkJRY0RBUVlJS3dZQkJRVUhBd0l3REFZRFZSMFRBUUgvQkFJd0FEQWRCZ05WSFE0RUZnUVVVSzZlClBKSUk4ZEVvMHBiSHZJb2UvRFlRMzVjd0h3WURWUjBqQkJnd0ZvQVVMTmJxVUR0bG9DM08vbjREdkFNbVpjZ0IKMU1rd2ZRWURWUjBSQkhZd2RJSVlabTl2WkMxcGJXRm5aUzFzWVdKbGJDMTNaV0pvYjI5cmdpUm1iMjlrTFdsdApZV2RsTFd4aFltVnNMWGRsWW1odmIyc3VaR1ZtWVhWc2RDNXpkbU9DTW1admIyUXRhVzFoWjJVdGJHRmlaV3d0CmQyVmlhRzl2YXk1a1pXWmhkV3gwTG5OMll5NWpiSFZ6ZEdWeUxteHZZMkZzTUEwR0NTcUdTSWIzRFFFQkN3VUEKQTRJQkFRQVkybkE2eU5LeGUxYWt4eWVIdXdIQm9ibDJlZU5BbEV5d0RXVHZLSktFdzNpQ2hPOVJXMmxXZloxSQovNGx5UEJNWUN2SUJ5ZVNOYStKRXpLbjBpdUZVbDJCV0w0TzhYQmFsQkRoTGk2T0hMeHZwOW56VFU3UDZEaU81CmpkU09tZTduVmdIaXlmanpRNUZFNjlGRStKOW5sNWkxWUlIRWlUcnZieGlMRUZ2SGRJeVZIRCszQVN3MTRrZmgKN3VRRnIrUFRONC91ZDZteFdwSk5WaFJRMU11VG5HUVl6M0EyV3c5Z1NGT1k0ekxyT2labzlBdVB1TXRpQmJUVQp4VTdxY2MzRU9BZ0c2bGIxRTlPYnpqZEJrREFmaEtwSmFuVVVNUzdNOFMrUkdjV3QzTUc4SjZLNDB5bEtENGFWCmthRW9GVVlxYzlDZGZLVHRGOGQ2TDVodEJWNE0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    service:
      name: fargate-sidecar-injector
      namespace: default
      path: /mutate
      port: 443
  failurePolicy: Fail
  matchPolicy: Equivalent
  name: fargate-sidecar-injector.default.svc
  namespaceSelector: {}
  objectSelector: {}
  reinvocationPolicy: Never
  rules:
  - apiGroups:
    - ""
    apiVersions:
    - '*'
    operations:
    - CREATE
    - UPDATE
    resources:
    - pods
    scope: '*'
  sideEffects: None
  timeoutSeconds: 10
EOF

# deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "1"
  generation: 1
  labels:
    app.kubernetes.io/instance: fargate-sidecar-injector
    app.kubernetes.io/name: fargate-sidecar-injector
    app.kubernetes.io/version: 1.16.0
  name: fargate-sidecar-injector
  namespace: default
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app.kubernetes.io/instance: fargate-sidecar-injector
      app.kubernetes.io/name: fargate-sidecar-injector
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        app.kubernetes.io/instance: fargate-sidecar-injector
        app.kubernetes.io/name: fargate-sidecar-injector
    spec:
      containers:
      - image: fargate-sidecar-injector:test
        imagePullPolicy: Never
        name: fargate-sidecar-injector
        ports:
        - containerPort: 8443
          name: http
          protocol: TCP
        resources: {}
        securityContext: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      serviceAccount: fargate-sidecar-injector
      serviceAccountName: fargate-sidecar-injector
      terminationGracePeriodSeconds: 30

# service

apiVersion: v1
kind: Service
metadata:
  labels:
    app.kubernetes.io/instance: fargate-sidecar-injector-webhook
    app.kubernetes.io/name: fargate-sidecar-injector-webhook
    app.kubernetes.io/version: 1.16.0
    helm.sh/chart: fargate-sidecar-injector-webhook-0.1.0
  name: fargate-sidecar-injector-webhook
  namespace: default
spec:
  internalTrafficPolicy: Cluster
  ipFamilies:
  - IPv4
  ipFamilyPolicy: SingleStack
  ports:
  - name: http
    port: 8443
    protocol: TCP
    targetPort: 443
  - name: https
    port: 443
    protocol: TCP
    targetPort: 443
  selector:
    app.kubernetes.io/instance: fargate-sidecar-injector-webhook
    app.kubernetes.io/name: fargate-sidecar-injector-webhook
  sessionAffinity: None
  type: ClusterIP



-------

#  old build process
```
docker build . -t fargate-sidecar-injector-webhook:test && minikube image load fargate-sidecar-injector-webhook:test && helm install fargate-sidecar-injector-webhook ./helm/chart/fargate-sidecar-injector -n default && k delete mutatingwebhookconfigurations/fargate-sidecar-injector-webhook && helm upgrade fargate-sidecar-injector-webhook ./helm/chart/fargate-sidecar-injector-webhook
```
# build and push to minikube
`docker build . -t fargate-sidecar-injector-webhook:test && minikube image load fargate-sidecar-injector-webhook:test`

# running it locally for test

``` sh

docker run  -v /Users/nash.singwango/source/repos/mziyabo/fargate-sidecar-injector:/etc/fargatesidecarinjector/ -it fargate-sidecar-injector-webhook:test

docker run  -v /Users/nash.singwango/source/repos/mziyabo/fargate-sidecar-injector:/etc/fargatesidecarinjector/ -p 8443:8443 -it fargate-sidecar-injector-webhook:test

```

