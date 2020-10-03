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
go-webserver   Deployment/go-webserver/scale   0% / 50%  1         10        1          18s
```


# Generate some load

Now, we will see how the autoscaler reacts to increased load. We will start a container, and send an infinite loop of queries to the php-apache service (please run it in a different terminal):


```bash
oc run -it --rm load-generator --image=busybox /bin/sh
```

Hit enter for command prompt

```bash
while true; do wget -q -O- http://go-webserver:8080; done
```

watch hpa and pods in another terminal

```bash
watch -n3 "echo "oc get hpa"; oc get hpa ;  echo "oc get deployment go-webserver" ;oc get deployment go-webserver ; echo "oc get pods -l run=go-webserver" ; oc get pods -l run=go-webserver"
```

```bash
Every 3.0s: echo oc get hpa; oc get hpa ;  echo oc get deployment go-webserver ;oc get de...  byroncserver: Sat Oct  3 15:57:20 2020

oc get hpa
NAME           REFERENCE                 TARGETS   MINPODS   MAXPODS   REPLICAS   AGE
go-webserver   Deployment/go-webserver   0%/50%    1         10        1          2d18h
oc get deployment go-webserver
NAME           READY   UP-TO-DATE   AVAILABLE   AGE
go-webserver   1/1     1            1           2d18h
oc get pods -l run=go-webserver
NAME                            READY   STATUS    RESTARTS   AGE
go-webserver-5879bd85df-n7mjd   1/1     Running   0          2d6h
```

Note: It may take a few minutes to stabilize the number of replicas. Since the amount of load is not controlled in any way it may happen that the final number of replicas will differ from this example.

# Stop load

We will finish our example by stopping the user load.

In the terminal where we created the container with busybox image, terminate the load generation by typing "\<Ctrl\> + C".

Then we will verify the result state (after a minute or so):

