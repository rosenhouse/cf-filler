# cf-filler
[![Build Status](https://api.travis-ci.org/rosenhouse/cf-filler.png?branch=master)](http://travis-ci.org/rosenhouse/cf-filler)

Generate variables to fill in [cf-deployment](https://github.com/cloudfoundry/cf-deployment) using the fancy new (alpha) [bosh-cli](https://github.com/cloudfoundry/bosh-cli).

## Install
```
go get github.com/rosenhouse/cf-filler
```

## Note
Make sure you're using `bosh-cli` v0.0.107 or higher.  Older versions [don't play nicely](https://github.com/cloudfoundry/bosh-cli/issues/46) with `cf-filler`.

Also, it appears that this functionality is now built the bosh cli!  Check out [cf-deployment#24](https://github.com/cloudfoundry/cf-deployment/pull/24) for an example.

## Usage
You'll need a "recipe" file that describes the variables to generate.
At the time of this writing, `cf-deployment` [includes its own recipe](https://github.com/cloudfoundry/cf-deployment/blob/master/cf-filler/recipe-cf-deployment.yml)
([permalink](https://github.com/cloudfoundry/cf-deployment/blob/197b32f158bd90c56a7d9b410119c551401a3108/cf-filler/recipe-cf-deployment.yml) in case that changes).

```bash
cf-filler -recipe ~/workspace/cf-deployment/recipe-cf-deployment.yml \
          -dnsname my-env.example.com > /tmp/vars.yml

bosh-cli build-manifest --var-errs --var-file=/tmp/vars.yml \
    ~/workspace/cf-deployment/cf-filler/cf-deployment.yml > /tmp/my-deployment.yml

bosh-cli -e my-director -d cf deploy /tmp/my-deployment.yml
```

If you modify `cf-deployment` (say to add extra jobs), you may need to customize your recipe file as well.

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
