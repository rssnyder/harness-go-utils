# harness-utils-go

helpful harness utils for go applications

## config

authentication and scoping

generall all helpers here will use the following for authentication and scoping:

- HARNESS_PLATFORM_API_KEY: harness nextgen api key
- HARNESS_ACCOUNT_ID: harness account id
- HARNESS_PLATFORM_ORGANIZATION: organization id
- HARNESS_PLATFORM_PROJECT: project id

leaving out project will result in org level resources
leaving out organization will result in account level resources

## secrets

set a text secret in harness
