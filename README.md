# API Specification

This repository contains the documents, discussions and implementations of the API specification for Dapr.

## Structure

The Open API definitions for Dapr Components can be found in the following directories: 
- [Dapr Statestore Component APIs](statestore-api/)
- [Dapr PubSub Component APIs](pubsub-api/) 

This repository also includes a Technology Compatibility Kit (TCK for short) which tests the Open APIs specification by creating Postman collections including Contract, Content and Variation tests using [Portman](https://github.com/apideck-libraries/portman) and then run all the tests against `daprd` which is started using [Docker Compose](tck/docker-compose.yaml). These tests are automated using GitHub Actions that you can find inside the [`.github/workflows`](.github/workflows/) directory, and use [Newman](https://github.com/postmanlabs/newman) to execute the generated tests.



## Contact

If you have a general question or suggestion regarding the work done in this special interest group, please open a GitHub issue in this repository.

* [API Spec Discord Channel](https://discord.com/channels/778680217417809931/935578420589522984)
* [Owners](CODEOWNERS)
