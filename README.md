# Cloud Build Ops

Cloud Build Ops is a tool that lets you manage Cloud Build pipeline configuration from yaml files which makes managing Cloud Build much easier and faster.

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

```
./cloudbuildops push -c pipelines/*
```
