# Hello, NATS (Streaming) on Dapr (on Kubernetes)!

First, deploy the sample Dapr app.

```
$ kubectl apply -f k8s/
$ kubectl get pods
NAME                              READY   STATUS    RESTARTS   AGE
hello-nats-dapr-78487cd6d-wj7kf   2/2     Running   0          47s
nats-0                            1/1     Running   0          46s
stan-0                            1/1     Running   0          46s
```

Next, we're going to print the messages.

In one terminal, forward local ports to `nats-0` container.

```
$ kubectl port-forward nats-0 4222:4222
Forwarding from 127.0.0.1:4222 -> 4222
Forwarding from [::1]:4222 -> 4222
```

In another terminal, print the messages.

```
$ stan-sub -c stan foo
Connected to nats://127.0.0.1:4222 clusterID: [stan] clientID: [stan-sub]
Listening on [foo], clientID=[stan-sub], qgroup=[] durable=[]
[#1] Received: sequence:1 subject:"foo" data:"{\"id\":\"cff4b937-7780-49c2-af40-79974a732ef1\",\"source\":\"hello-nats-dapr\",\"type\":\"com.dapr.event.sent\"
,\"specversion\":\"1.0\",\"datacontenttype\":\"text/plain\",\"data\":\"Hello 1!\",\"subject\":\"00-7fef701158ea077b622030030ff24a9c-77a96368f8f27ac6-00\",\"to
pic\":\"foo\",\"pubsubname\":\"messagebus\"}" timestamp:1602117353553901836
```
