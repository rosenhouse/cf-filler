# cf-filler

Generate variables to fill in [cf-deployment](https://github.com/cloudfoundry/cf-deployment) using the fancy new (alpha) [bosh-cli](https://github.com/cloudfoundry/bosh-cli).

## Install
- Grab [latest binary release](https://github.com/rosenhouse/cf-filler/releases) and `chmod +x`

OR

- 

  ```
  go get github.com/rosenhouse/cf-filler
  go install github.com/rosenhouse/cf-filler
  ```


## Usage
```bash
./cf-filler --dnsname my-env.example.com > /tmp/vars.yml

bosh build-manifest --var-files=/tmp/vars.yml --var-errs ~/workspace/cf-deployment/cf-deployment.yml > /tmp/my-deployment.yml

bosh -e my-director -d cf deploy /tmp/my-deployment.yml
```

## Running the tests
Currently it just runs `bosh build-manifest`. 
```bash
fly -t myconcourse execute -c ci/test.yml
```
