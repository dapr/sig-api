# API Specification

This repository contains the documents, discussions and implementations of the API specification for Dapr.

## Structure

The Open API definitions for Dapr Components can be found in the following directories: 
- [Dapr Statestore Component APIs](statestore-api/)
    - HTTP-1 (Manual spec for endpoints to generate clients)
    - GRPC (Generated Schemas from .proto)
- Dapr PubSub Component APIs (TBD)

This repository also includes a Technology Compatibility Kit (TCK for short) which tests the Open APIs specification by generating clients and interact with the real Dapr APIs. You can find the TCK test inside the `tck/` directory. 



## Contact

If you have a general question or suggestion regarding the work done in this special interest group, please open a GitHub issue in this repository.

* [API Spec Discord Channel](https://discord.com/channels/778680217417809931/935578420589522984)
* [Owners](CODEOWNERS)
