## API Load

!!! This is still a work in progress !!!

`api-load` is a small tool to pipe massive data via API calls. It will try to minimize memory buffering in order to scale (not a priority for first releases)

Initially developed to work with `coog-api`, it should be suitable for any REST API.

### Why?

Our first internal need was to import massively a set of contracts to Coog platform

- data is in `JSON` format
- Coog API ensure all kind of checks on imported data (format, consistency, rates, etc.)
- Coog API contract entrypoint is unitary and not intended manage piping or scheduling
- Serializing calls is not performant and needs scripting

Even if it looks simple, these tasks usually turns into messy development projects

`api-load` has been created to fix that, with focus on

- performance
- robustness (long process for migration)
- clean reporting

### Features

- configurable concurrency
- managed authentication (only cookie for now)
- multi-platform and portability (Go)

### Installation and usage

- Grab a binary from [here](https://github.com/coopengo/api-load/releases)
- copy it to somewhere in your `PATH`

#### example of usage

- Supposing that `coog-api` is running at `http://localhost:3000`
- Prepare your data file (`data.json`)
    ```
    [
        {"name": "company1"},
        {"name": "company2"}
    ]
    ```
- `api-load -url "http://localhost:3000/party/company" -data data.json`

### TODO (urgent first)

- check cookies (passed by reference, go concurrency?)
- make auth more generic
- better logs and result (average call, stats, etc.)
