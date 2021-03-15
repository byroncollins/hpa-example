# Description

Horizontal Pod Autoscaler automatically scales the number of Pods in a replication controller, deployment, replica set or stateful set based on observed CPU utilization (or, with beta support, on some other, application-provided metrics).

Based on https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/ but implemented in Go as a scratch container and doesn't run as root.


# Deploy go-webserver

Deploy straight from github or locally.

Prequisites: 
1. OpenShift must be able to access `docker.io` images.
2. (Optional) Create new project/namespace such as: `oc new-project hpa-example`

### 1. Deploy straight from github

```bash
oc apply -f https://raw.githubusercontent.com/byroncollins/hpa-example/master/manifests/hpa-example.yaml
```
### 2. Deploy locally

```bash
git clone https://github.com/byroncollins/hpa-example.git
oc apply -f hpa-example/manifests/hpa-example.yaml
```

### Successful deploy

```
deployment.apps/go-webserver created
service/go-webserver created
```

# Create Horizontal Pod Autoscaler

Configure horizontal pod autoscaler (must be run in same namespace as `go-webserver`)
```bash
oc autoscale deployment go-webserver --cpu-percent=50 --min=1 --max=10
```
```
horizontalpodautoscaler.autoscaling/go-webserver autoscaled
```

Check the current status of autoscaler by running the following in the same namespace:

```bash
oc get hpa
```

```
NAME         REFERENCE                     TARGET    MINPODS   MAXPODS   REPLICAS   AGE
go-webserver   Deployment/go-webserver/scale   0% / 50%  1         10        1          18s
```

# Generate some load

Now, we will see how the autoscaler reacts to increased load. We will start a new busybox container, and send an infinite loop of queries to the go-webserver service.

Open a new terminal window and login to openshift as before, then run the following to enter the command prompt of the busybox:

```bash
oc run -it --rm load-generator --image=busybox /bin/sh
```

Run the performance load:

```bash
while true; do wget -q -O- http://go-webserver:8080; done
```

Watch hpa and pods in another terminal:

```bash
watch -n3 "echo "oc get hpa"; oc get hpa ;  echo "oc get deployment go-webserver" ;oc get deployment go-webserver ; echo "oc get pods -l run=go-webserver" ; oc get pods -l run=go-webserver"
```

The following should be displayed and refreshed every 3 seconds:
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

Stop the performance load by typing `ctrl+c` in the terminal used to build the busybox container and run the load.

