# Development using the Demo

The created kind cluster in the demo comes with a local registry, so that the demo can be used within your development flow. For that you have to build the Docker image with your changes and push it into the local registry:

```sh
docker build -f ./cmd/kobs/Dockerfile -t localhost:5000/kobs:dev .
docker push localhost:5000/kobs:dev
```

When you have pushed your custom image, you have to adjust the patch file for the kobs deployment. Change the Docker image in the [`kobs.yaml`](https://github.com/kobsio/kobs/blob/main/deploy/demo/observability/kobs.yaml) file to the following:

```sh
spec:
  template:
    spec:
      containers:
        - name: kobs
          image: localhost:5000/kobs:dev
```

Now you can deploy the files for the `observability` namespace again:

```sh
kustomize build deploy/demo/observability | kubectl apply -f -
```

Finally you can check if the kobs Pod is using your image, with the following command:

```sh
k get pods -n observability -l app.kubernetes.io/name=kobs -o yaml | grep "image: localhost:5000/kobs:dev"
```