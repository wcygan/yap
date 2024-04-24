# Real-time Chat Application

## Quick Start

Install [Minikube](https://minikube.sigs.k8s.io/docs/start/) and [Skaffold](https://skaffold.dev/docs/install/#standalone-binary), then run:

```
minikube start && skaffold dev
```

## Connecting to PostgreSQL Locally

Pick a postgres node (`postgres-69c569c6c9-wj2zx`):

```
k get po
NAME                            READY   STATUS    RESTARTS   AGE
auth-service-7b47d6d967-cj82q   1/1     Running   0          26s
postgres-69c569c6c9-wj2zx       1/1     Running   0          28s
yap-api-75dbfb86b5-c9h9d        1/1     Running   0          24s

```

Port forward it:

```
kubectl port-forward postgres-69c569c6c9-wj2zx 5432:5432
```

Use a Database Driver and a connection string similar to the following: 

```
postgres://postgres:your-password-here@postgres:5432/postgres?sslmode=disable
```

I like [DataGrip](https://www.jetbrains.com/datagrip/) or the [IntelliJ Postgres Driver](https://www.jetbrains.com/help/idea/postgresql.html).