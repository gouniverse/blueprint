# Blueprint 

![tests](https://github.com/gouniverse/blueprint/workflows/tests/badge.svg)

## Development IDE

<a href="https://gitpod.io/#https://github.com/gouniverse/blueprint" style="float:right:" target="_blank">
    <img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy">
</a>
<a href="https://app.codeanywhere.com/#https://github.com/gouniverse/blueprint/blob/b7b76728d496499527c9bfd6476b3f0784708958/README.md" style="float:right:" target="_blank">
    <img src="https://codeanywhere.com/img/open-in-codeanywhere-btn.svg" alt="Open in Gitpod" loading="lazy">
</a>

## URLS

- https://YOURAPPURL
- https://YOURAPPURL.a.run.app (Dev)
This is a quick start blueprint for an MVC web applications

- Ready to develop in the cloud (Gitpod / Github CodeSpaces)
- Database connection setup (SQLite example)
- Router setup
- Background tasks setup
- Scheduler setup
- Controllers setup
- CMS setup

## Installation

```bash
git clone https://github.com/gouniverse/blueprint
```

## Environment Variables

- Copy the `.env_example` file to `.env`

```bash
cp .env_example .env
```

- Set the dev vault values

```bash
task env-dev
```

- Set the prod vault values

```bash
task env-prod
```


## Local Development

- Just starting
```bash
task dev:init
```

- Run in development mode
```bash
task dev
```

## Development on Gitpod

Use the link on the top of this README

## Testing

Running all tests

```bash
task test
```

-Running individual test

```
go test -run ^TestGuestFunnelTestSuite$
```

## Coverage Report

```bash
task cover
```

## CLI Commands

Deploy Live:

```bash
task deploy:live
```

Deploy Staging:

```bash
task deploy:staging
```

List Routes:

```bash
go run . routes list
```

Run task:

```bash
go run . task run ...
```

Run job:

```bash
go run . job run ...
```
