# cf-filler

work in progress: generate variables to fill in cf-deployment


## to test
```bash
bosh-cli build-manifest --var-files=<(go run main.go) --var-errs fixtures/cf-deployment/cf-deployment.yml
```

```bash
fly -t myconcourse execute -c ci/test.yml
```
