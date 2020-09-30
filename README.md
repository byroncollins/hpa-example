# Description

Horizontal Pod Autoscaler automatically scales the number of Pods in a replication controller, deployment, replica set or stateful set based on observed CPU utilization (or, with beta support, on some other, application-provided metrics).

The default hpa-example doesnt run in OpenShift


# Run go-webserver

```bash
oc apply -f https://github/byroncollins/manifests/hpa-example.yaml
```

```
deployment.apps/go-webserver created
service/go-webserver created
```

# Create Horizontal Pod Autoscaler

```bash
oc autoscale deployment go-webserver --cpu-percent=50 --min=1 --max=10
```
```
horizontalpodautoscaler.autoscaling/php-apache autoscaled
```

We may check the current status of autoscaler by running:

```bash
oc get hpa
```

```
NAME         REFERENCE                     TARGET    MINPODS   MAXPODS   REPLICAS   AGE
php-apache   Deployment/php-apache/scale   0% / 50%  1         10        1          18s
```


