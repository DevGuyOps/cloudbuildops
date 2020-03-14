# Cloud Build Ops

Cloud Build Ops is a tool that lets you manage Cloud Build pipeline configuration from yaml files which makes managing Cloud Build much easier and faster.

You will still need to add your repository to Cloud Build via the console (i.e. Connect Repository). However, all of the management after adding the repository can be managed by this tool.

## Functions

### Get

Write all existing cloud build pipelines to file

#### Flags

| Flag            | Description                              | Required |
| --------------- | ---------------------------------------- | -------- |
| -projectid / -p | Project ID of the GCP project            | TRUE     |
| -output / -o    | Output directory to publish config files | TRUE     |

#### Example usage

```
./cloudbuildops get -p my-first-project -o pipelines
```

### Push

Create/Update cloud build pipelines from the proveided config files

#### Flags

| Flag         | Description                               | Required |
| ------------ | ----------------------------------------- | -------- |
| -config / -c | Path to config files (Supports wildcards) | TRUE     |

#### Example usage

```
./cloudbuildops push -c pipelines/*
```

## Using the Container

For convenience you can use the published Docker container.

Container: **guywatson/cloudbuildops:latest**

### IAM Permissions

The service account will need the below permissions

- Cloud Build Editor
- Cloud Build Viewer

### Example usage

```
docker run \
    -e GOOGLE_APPLICATION_CREDENTIALS=/service_account.json \
    -v $PWD/service_account.json:/service_account.json \
    -v $PWD/output:/output \
    guywatson/cloudbuildops:latest \
    get --projectid MY_PROJECT --output /output
```
