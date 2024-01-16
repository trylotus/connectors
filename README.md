# 🐙 Connectors
This repo contains all the Lotus team supported connectors.

Github Actions is set up to build, run tests, and push docker image for each connector individually.

### 🏃‍♀️ To run
```
cd <connector_name>
go run cmd/<connector_name>/main.go
```

### 🚀 To deploy
For Lotus devs, ensure that the image was built successfully (check Github Actions) and pushed to our internal artifact registry, then update Helmfile values.yaml. 

If you would like to build your own connector, start by taking a look at our [documentation](https://docs.trylotus.xyz/docs/intro), and [reach out to our team](mailto:hello@trylotus.xyz) with any questions you may have.