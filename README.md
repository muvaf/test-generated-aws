# AWS Provider

`test-generated-aws` is a Crossplane Provider that's been generated by
[AWS Controllers for Kubernetes (ACK)](https://github.com/aws/aws-controllers-k8s/) generation pipeline.

Please do not use this in production since it's still in very early stages.

## Generation Steps

The first step is cloning ACK:
```console
git clone https://github.com/aws/aws-controllers-k8s.git
cd aws-controllers-k8s
```

Build the `ack-generate` tool:
```console
make build-ack-generate
```

Now let's generate API types for ECR:
```
./bin/ack-generate crossplane apis ecr --provider-dir <path to root of this repo>
``` 

You will see that the `.go` files for ECR API appears under `apis/ecr`
