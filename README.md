## Coog API Load

This is a util to call Coog API intensively. The most common use case is small
migrations to import contracts to Coog via API.

Nothing is hardly specific to Coog API, coog-api-load could be used against any
REST API.

### Targets

- configure concurrency
- manage authentication
- multi-platform and portability (Go)

### TODO (in order)

- check cookies (passed by reference, concurrency?)
- make auth more generic
- better logs and result (average call, stats, etc.)
