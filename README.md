# connectors
This repo contains all the Nakji team supported connectors.

Github Actions is set up to build, run tests, and push docker image for each connector individually.

### To run
```
cd <connector_name>
go run cmd/<connector_name>/main.go
```

### To deploy
Ensure that the connector's docker image has been built successfully and pushed to the docker repo in either Nakji [Slack](https://blepai.slack.com/archives/C03DL2A73SB) or [Google Artifact Registry](https://console.cloud.google.com/artifacts/docker/blep-core/asia/docker?project=blep-core).

Update Helmfile to use the new docker image.
