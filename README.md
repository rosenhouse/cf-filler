# cf-filler
Generate variables to fill in [cf-deployment](https://github.com/cloudfoundry/cf-deployment) using the fancy new (alpha) [bosh-cli](https://github.com/cloudfoundry/bosh-cli).

## Install
Grab [latest binary release](https://github.com/rosenhouse/cf-filler/releases) and `chmod +x`

OR

```
go get github.com/rosenhouse/cf-filler
```


## Usage
```bash
cf-filler --dnsname my-env.example.com > /tmp/vars.yml

bosh-cli build-manifest --var-errs --var-file=/tmp/vars.yml cf-deployment.yml > /tmp/my-deployment.yml

bosh-cli -e my-director -d cf deploy /tmp/my-deployment.yml
```

## Running the tests

- Locally
  ```bash
  go get github.com/cloudfoundry/bosh-cli
  ./test
  ```

- On a remote [Concourse](http://concourse.ci/)
  ```bash
  fly -t myconcourse execute -x -c ci/test.yml
  ```
