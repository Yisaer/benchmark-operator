# How to use

## Prerequisites

You need a running mysql-compatible database (eg. mysql, tidb) in your k8s cluster.

If you don't have one, you can use [demo-mysql.yaml](https://github.com/Yisaer/benchmark-operator/blob/master/manifests/demo-mysql.yaml) to deploy one:

```bash
kubectl apply -f https://raw.githubusercontent.com/Yisaer/benchmark-operator/master/manifests/demo-mysql.yaml
```

## Install crd

Just apply the [crd.yaml](https://github.com/Yisaer/benchmark-operator/blob/master/manifests/crd.yaml) file:

```bash
kubectl apply -f https://raw.githubusercontent.com/Yisaer/benchmark-operator/master/manifests/crd.yaml
```

## Install controller

Just apply the [controller.yaml](https://github.com/Yisaer/benchmark-operator/blob/master/manifests/controller.yaml) file:

```bash
kubectl apply -f https://raw.githubusercontent.com/Yisaer/benchmark-operator/master/manifests/controller.yaml
```

## Do benchmark

### Prepare data

The first stage of benchmark is load benchmark the data into the target database, you need to deploy a `DataBaseBenchmarkPrepare` resource to achieve this.

You can refer to [crd.yaml](https://github.com/Yisaer/benchmark-operator/blob/master/manifests/crd.yaml) to find out the fields you need in the yaml file for this.

We suggest using the [sample yaml](https://github.com/Yisaer/benchmark-operator/blob/master/config/samples/benchmark.cloud_v1alpha1_databasebenchmarkprepare.yaml)
and change the values in this yaml to adapt your own database.

```bash
curl https://raw.githubusercontent.com/Yisaer/benchmark-operator/master/config/samples/benchmark.cloud_v1alpha1_databasebenchmarkprepare.yaml > prepare.yaml
# edit prepare.yaml to adapt your own database
kubectl apply -f prepare.yaml
```

### Run benchmark

Then you can run the benchmark, you need to deploy a `DataBaseBenchmarkRun` resource to achieve this.

You can also refer to [crd.yaml](https://github.com/Yisaer/benchmark-operator/blob/master/manifests/crd.yaml) to find out the fields you need in the yaml file for this.

We suggest using the [sample yaml](https://github.com/Yisaer/benchmark-operator/blob/master/config/samples/benchmark.cloud_v1alpha1_databasebenchmarkrun.yaml)
and change the values in this yaml to adapt your own database.

```bash
curl https://raw.githubusercontent.com/Yisaer/benchmark-operator/master/config/samples/benchmark.cloud_v1alpha1_databasebenchmarkrun.yaml > run.yaml
# edit run.yaml to adapt your own database
kubectl apply -f run.yaml
```

### Check result

The result is in the log of the pod which `DataBaseBenchmarkRun` created.

Just `kubectl log` to view it.
