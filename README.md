# Cleanup Deployments in Test Environments

This program cleans up deployments in it's namespace based on the value of the label that defines the the time a resource is allowed to live

resource yaml's would need to have a label as below

```
apiVersion: extensions/v1beta1 
kind: Ingress 
metadata: 
    name: "{APP_NAME}" 
    labels:
        life: "{DAYS_TO_LIVE}"    
spec: rules: # DNS name your application should be exposed on
    host: "{APP_NAME}.{TARGET_CLUSTER}.devops.lk"
    http: 

```

DAYS_TO_LIVE is a value which needs to be defined in hours as of now.



or the label can be manually set"
 
 ```
kubectl label ing app-pr-12 life=160h --overwrite

kubectl label deployment app-pr-12 life=160h --overwrite

 ```

WARNING!!....This is only intented to run in a  kubernetes test/development cluster. Never Deploy this to production cluster.

## How to Build

manual:



build application and create a binary named cleaner:

```
go build -o kubemeister
```


## How to run 

execute the binary:
if you build using makefile binary will be in build/osx/ or build/linux folder
./kubemeister
