# Marshalling k8s protos

Run:

```
$ git clone https://github.com/lukaszx0/k8s-protos-issue $GOPATH/src/k8s-protos-issue
$ go run $GOPATH/src/k8s-protos-issue/main.go
```

Output:

```
$ go run $GOPATH/src/k8s-protos-issue/main.go
deployment.Marshal(): [10 18 10 0 18 0 26 0 34 0 42 0 50 0 56 0 66 0 122 0 18 60 26 50 10 18 10 0 18 0 26 0 34 0 42 0 50 0 56 0 66 0 122 0 18 28 26 0 50 0 66 0 74 0 82 0 88 0 96 0 104 0 130 1 0 138 1 0 154 1 0 194 1 0 34 2 10 0 40 0 56 0 26 12 8 0 16 0 24 0 32 0 40 0 56 0]
proto: no encoder for TypeMeta v1.TypeMeta [GetProperties]
proto: no coders for v1.ObjectMeta
proto: no encoder for ObjectMeta v1.ObjectMeta [GetProperties]
proto: no coders for v1.DeploymentSpec
proto: no encoder for Spec v1.DeploymentSpec [GetProperties]
proto: no coders for v1.DeploymentStatus
proto: no encoder for Status v1.DeploymentStatus [GetProperties]
proto.Marshal(embededDeployment): [10 6 102 111 111 98 97 114 18 96 10 18 10 0 18 0 26 0 34 0 42 0 50 0 56 0 66 0 122 0 18 60 26 50 10 18 10 0 18 0 26 0 34 0 42 0 50 0 56 0 66 0 122 0 18 28 26 0 50 0 66 0 74 0 82 0 88 0 96 0 104 0 130 1 0 138 1 0 154 1 0 194 1 0 34 2 10 0 40 0 56 0 26 12 8 0 16 0 24 0 32 0 40 0 56 0]
```

Why am I getting all those `proto: no (de|en)coder` errors (warnings?) ?

Notes:

* k8s uses gogo for protos but using `protoc-gen-gogo` for `embedded_deployment.proto` gives the same output