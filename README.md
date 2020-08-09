![logo](logo.png)

_rabbitmkube_ is an autoscaler for k8s deployments that need to base their number of replicas on
the number of messages contained in a queue.

It is a simpler alternative to exporting `custom-metrics` to Kubernetes and creating _HorizontalPodAutoScalers_.

## Requirements
You need to have RabbitMQ Manager plugin so that _rabbitmkube_ can fetch RabbitMQ like [this](https://pulse.mozilla.org/api/).

## Step by Step
1. Create a `RoleBinding` as shown in the command below:

`kubectl create clusterrolebinding rabbitmkube --clusterrole=cluster-admin --serviceaccount=default:default`

2. Use the `deployment.yaml` file in this repository by changing its environment variable values.

3. Add _rabbitmkube_ annotations to your deployments as shown in the example below. 
   `messages-per-pod` refers to the number of messages that a pod can take at once.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-deployment
  labels:
    app: example
  annotations:
    rabbitmkube/min-replicas: "0"
    rabbitmkube/max-replicas: "3"
    rabbitmkube/messages-per-pod: "1"
    rabbitmkube/queue-name: test-autoscaler
spec:                       
  selector:
    matchLabels:
      app: example
  template:
    metadata:
      labels:
        app: example
    spec:
      containers:
      - name: example
        image: example/example
```

Done! Every 10 seconds _rabbitmkube_ will look for deployments annotated with _rabbitmkube_ values
and add/remove new replicas to it.
