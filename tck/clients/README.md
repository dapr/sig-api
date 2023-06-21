# Generated Clients

DO NOT change the generated client code. 
To regenerate run from inside this directory, to generate the Statestore client code: 


```
oapi-codegen -generate types,client,spec -package statestore ../../statestore-api/http-1/statestore.yaml > statestore.gen.go
```

