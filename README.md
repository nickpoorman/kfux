# kfux

Is a kubernetes output unfuxer. For some reason, configmaps are json encoded in
a value even if they aren't json.

```
kubectl get configmap nats-config -ojson | jq '.data."nats.conf"' | kfux -p
```

You can accomplish something similiar with jq's `fromjson` function, but it will
choke on the first line of output that isn't json. Maybe there's a way to get jq
not to do that? Let me know!

```
kubectl get configmap nats-config -ojson | jq '.data."nats.conf" | fromjson`
> jq: error (at <stdin>:27): Invalid numeric literal at line 1, column 2 (while parsing '# PID file shared with configuration reloader.
```
