# Go Sigterm Ignore

Simple Golang web app, using the standard library which will receive a SIGTERM but then ignore it. Used to demonstrate what happens if an app does not handle SIGTERM correctly.

## Pushing to Cloud Foundry

Push the app to the current space based on the manifest:

```
cf push
```