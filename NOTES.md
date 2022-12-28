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