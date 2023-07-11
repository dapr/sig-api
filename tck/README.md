# Technology Compatibility Kit

This project aims to test the Open API specifications against different implementations. 

This project uses Portman, Postman and Newman to create and perform tests against the Dapr APIs.

## Findings

- Contract Tests defined by default here: https://github.com/apideck-libraries/portman/blob/main/portman-config.default.json 
    - Failing on content type return for empty bodys (which defaults to text/plain)
    - Schema validation for empty bodys return