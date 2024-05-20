# READING LIST API

## My motivation

Throughout the day, I read a lot of material. Sometimes, at work, I don't have time to read it immediately. Therefore, I thought it would be beneficial to have an app that allows sharing articles between various devices. This is the backend component of the app.

I have no prior experience with Go. I decided to use it for this project because I haven't used it before xd.

## Getting Started

Before running the app, you need to create a `.env` file. You can copy the example file and fill in the values:

```bash
$ cp .env.example .env
```

## Running the app

Then you can run the app with the following commands:

```bash
#build
$ make dev-build
#start
$ make dev-start
```

## Stop

```bash
# stop
$ make dev-down
```

### All ENV list ###

| Variable            | Description                  | Type     | Required | Default |
|---------------------|------------------------------|----------|----------|---------|
| `API_PORT`          | The port the API listens on. | `number` | true     | 4000    |
| `AUTH_USERNAME`     | The username for Basic Auth. | `string` | true     | -       |
| `AUTH_PASSWORD`     | The password for Basic Auth. | `string` | true     | -       |
| `DATABASE_URL`      | URL for postgres connection. | `string` | true     | -       |
| `POSTGRES_USER`     | The postgres password.       | `string` | false    | -       |
| `POSTGRES_PASSWORD` | The postgres username.       | `string` | false    | -       |
| `POSTGRES_DB_NAME`  | The postgres db name.        | `string` | false    | -       |
| `POSTGRES_PORT`     | The postgres port.           | `string` | false    | -       |
| `POSTGRES_HOST`     | The postgres host.           | `string` | false    | -       |

Env variables marked as `false` are for local development purposes and they are present in the `.env.example` file.
For production, you should use the `DATABASE_URL` variable.