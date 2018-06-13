## API Load

:warning: This is still a work in progress :warning:

`api-load` is a small tool to pipe massive data via API calls. It will try to minimize memory buffering in order to scale (not a priority for first releases)

Initially developed to work with `coog-api`, it should be suitable for any REST API.

### Why?

Our first internal need was to import massively a set of contracts to Coog platform

- Data is in `JSON` format
- Coog API ensure all kind of checks on imported data (format, consistency, rates, etc.)
- Coog API contract entrypoint is unitary and not intended manage piping or scheduling
- Serializing calls is not performant and needs scripting

Even if it looks simple, these tasks usually turns into messy development projects

`api-load` has been created to fix that, with focus on

- Performance
- Robustness (long process for migration)
- Clean reporting

### Features

- Configurable concurrency
- Managed authentication (only cookie for now)
- Multi-platform and portability (Go)

**Public API will be frozen soon for v1.0.0.**

### Installation

`api-load` is a standalone binary

- Install: copy the binary somewhere in your `PATH`
- Uninstall: remove the binary

Binaries are hosted [here](https://github.com/coopengo/api-load/releases)

For Gophers, just `go get github.com/coopengo/api-load`

### Usage example

- Supposing that `coog-api` is running at `http://localhost:3000`
- Prepare your data file (`data.json`)
    ```
    [
        {"name": "company1"},
        {"name": "company2"}
    ]
    ```
- Launch data loading
    ```
    api-load \
        -auth "cookie <username>:<password>@http://localhost:3000/auth/login" \
        -url "http://localhost:3000/party/company" \
        data.json
    ```

### TODO (urgent first)

- [X] Some cleaning (Check Go concurrency on cookies, set url and method on job model)
- [X] Scalability (support huge file size)
- [X] Better logs and result (average call, stats, API errors)
- [X] Make auth more generic (the way we format username, password, other methods)
