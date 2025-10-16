## required packages 
- go statick check 
  - go install honnef.co/go/tools/cmd/staticcheck@latest\n
- act 
  - if you want to run github actions locally 






## Github actions locally with act 

first install act, you can install it with brew install act 

then make sure you have added this
`--container-architecture linux/amd64`
to the `~/.actrc`

then you can run `act -l` to see the acts and also `act -j test -r`

test is the name of one of stages in the ci.yaml I have in this repo and -r means it will not remove the docker after the the test job  and will help you to run the CI faster next time 

```
act -s GITHUB_TOKEN="github token " --reuse
````
