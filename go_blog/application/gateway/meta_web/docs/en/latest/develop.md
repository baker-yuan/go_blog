---
title: Development Guide
---

<!--
-->

The Dashboard contains both `manager-api` and `web` parts, so you need to start the development environment separately.

## Prerequisites

Before development, refer to this [guide](./install.md) to install dependencies.

## Clone the project

```sh
$ git clone -b release/3.0 https://github.com/apache/apisix-dashboard.git
```

## Start developing

```sh
$ cd apisix-dashboard
```

### manager-api

1. Please change the configuration in `api/conf/conf.yaml`.

2. In the root directory, launch development mode.

```sh
$ make api-run
```

3. In the root directory, stop development mode.

```sh
$ make api-stop
```

4. Please refer to the [FAQ](./FAQ.md) about the problem of displaying exception in the dashboard after adding custom plugins or modifying plugin's schema.

5. If writing an back end E2E test, please refer to the [Back End E2E Writing Guide](./back-end-tests.md)

### web

1. Go to the `web` directory.

```sh
$ cd ./web
```

2. Please change the `manager-api` address in the `web/.env` file. If you follow this guidelines, the address may need to be set as below.

> All commands here are for Linux environment, other systems please use the corresponding commands for your platform. You are also welcome to contribute your own methods.

```bash
echo "SERVE_URL_DEV=http://localhost:9000" > web/.env
```

If you don't want to create the file, you can also export the variable.

```bash
export SERVE_URL_DEV=http://localhost:9000
```

3. Launch development mode

```sh
$ yarn install

$ yarn start
```

> If there is an error about gyp during yarn install, please ignore it and go ahead!

4. If writing an front end E2E test, please refer to the [Front End E2E Writing Guide](./front-end-e2e.md)
