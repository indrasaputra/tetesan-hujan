# Tetesan Hujan

[![Go Report Card](https://goreportcard.com/badge/github.com/indrasaputra/tetesan-hujan)](https://goreportcard.com/report/github.com/indrasaputra/tetesan-hujan)
[![Workflow](https://github.com/indrasaputra/tetesan-hujan/workflows/Test/badge.svg)](https://github.com/indrasaputra/tetesan-hujan/actions)
[![codecov](https://codecov.io/gh/indrasaputra/tetesan-hujan/branch/master/graph/badge.svg?token=R17RPYS094)](https://codecov.io/gh/indrasaputra/tetesan-hujan)
[![Maintainability](https://api.codeclimate.com/v1/badges/e2e45026960fb8cf7725/maintainability)](https://codeclimate.com/github/indrasaputra/tetesan-hujan/maintainability)
[![Go Reference](https://pkg.go.dev/badge/github.com/indrasaputra/tetesan-hujan.svg)](https://pkg.go.dev/github.com/indrasaputra/tetesan-hujan)

## Description

Tetesan Hujan is a bot to connect [Telegram](https://telegram.org/) with [Raindrop](https://raindrop.io/).
Its main purpose is currently to save a bookmark for my own purpose.

## Owner

[Indra Saputra](https://github.com/indrasaputra)

## Development

- Install Go

    This project uses version 1.16. Follow [Golang installation guideline](https://golang.org/doc/install).

- Create your Telegram bot

    Read [https://telegram.org/blog/bot-revolution](https://telegram.org/blog/bot-revolution).

    Follow [https://core.telegram.org/bots](https://telegram.org/blog/bot-revolution) for developer guide.

- Create Raindrop.io account

    Get the Raindrop's token in [settings](https://app.raindrop.io/settings/integrations).

    Follow [https://developer.raindrop.io/](https://developer.raindrop.io/) for more guidance.

- Clone the project (use one of the two methods below)

    Use SSH
    ```
    $ git@github.com:indrasaputra/tetesan-hujan.git
    ```
    
    Use HTTP
    ```
    https://github.com/indrasaputra/tetesan-hujan.git
    ```

- Go to project folder

    Usually, it would be `cd go/src/github.com/indrasaputra/tetesan-hujan`.

- Fill in the environment variables

    Copy the sample env file.
    ```
    cp env.sample .env
    ```
    Then, fill the values according to your setting in `.env` file.

- Download the dependencies

    ```
    make dep-download
    ```
    or run this command if you don't have `make` installed in your local.
    ```
    go mod download 
    ```

- Run the application

    ```
    go run cmd/bot/main.go
    ```

- Expose your localhost to the internet

    Usually, I use [https://ngrok.com/](https://ngrok.com/)
    ```
    ngrok http $PORT
    ```

- Send some messages to your bot in Telegram.

## Deployment

Currently, this project is deployed in [Google Cloud Run](https://cloud.google.com/run).
The deployment process definiton is stated and ruled in [Github Actions](https://github.com/indrasaputra/tetesan-hujan/blob/master/.github/workflows/test-and-deploy.yml).