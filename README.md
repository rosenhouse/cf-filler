# cf-filler
[![Build Status](https://api.travis-ci.org/rosenhouse/cf-filler.png?branch=master)](http://travis-ci.org/rosenhouse/cf-filler)

Generate variables to fill in [cf-deployment](https://github.com/cloudfoundry/cf-deployment) using the fancy new (alpha) [bosh-cli](https://github.com/cloudfoundry/bosh-cli).

## Install
```
go get github.com/rosenhouse/cf-filler
```

## Note
The latest version of the `bosh-cli` [doesn't play nicely](https://github.com/cloudfoundry/bosh-cli/issues/46) with `cf-filler`.

The workaround is to use an older version
```
git -C $GOPATH/src/github.com/cloudfoundry/bosh-cli checkout 810c591
go install github.com/cloudfoundry/bosh-cli
```

## Usage
```bash
cf-filler -dnsname my-env.example.com -recipe recipe-cf-deployment.yml > /tmp/vars.yml

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
