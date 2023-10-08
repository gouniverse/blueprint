# Blueprint 

<a href="https://gitpod.io/#https://github.com/gouniverse/blueprint" style="float:right;" target="_blank"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

![tests](https://github.com/gouniverse/blueprint/workflows/tests/badge.svg)

This is a quick start blueprint for an MVC web applications

- Redy to develop in the cloud (Gitpod / Github CodeSpaces)
- Router setup
- Controllers setup
- Database connection setup (SQLite example)
- CMS setup

## Installation

```bash
git clone https://github.com/gouniverse/blueprint
```

## Local Development

```bash
task dev:init
```

```bash
task dev
```

## Testing

```bash
task test
```

## Coverage Report

```bash
task cover
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
