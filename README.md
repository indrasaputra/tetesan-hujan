# Tetesan Hujan

[![Go Report Card](https://goreportcard.com/badge/github.com/indrasaputra/tetesan-hujan)](https://goreportcard.com/report/github.com/indrasaputra/tetesan-hujan)
[![Workflow](https://github.com/indrasaputra/tetesan-hujan/workflows/Test/badge.svg)](https://github.com/indrasaputra/tetesan-hujan/actions)
[![codecov](https://codecov.io/gh/indrasaputra/tetesan-hujan/branch/main/graph/badge.svg?token=R17RPYS094)](https://codecov.io/gh/indrasaputra/tetesan-hujan)
[![Maintainability](https://api.codeclimate.com/v1/badges/e2e45026960fb8cf7725/maintainability)](https://codeclimate.com/github/indrasaputra/tetesan-hujan/maintainability)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=indrasaputra_tetesan-hujan&metric=alert_status)](https://sonarcloud.io/dashboard?id=indrasaputra_tetesan-hujan)
[![Go Reference](https://pkg.go.dev/badge/github.com/indrasaputra/tetesan-hujan.svg)](https://pkg.go.dev/github.com/indrasaputra/tetesan-hujan)

## Description

Tetesan Hujan is a bot to connect [Telegram](https://telegram.org/) with [Raindrop](https://raindrop.io/).
Its main purpose is currently to save a bookmark for my own purpose.

## Owner

[Indra Saputra](https://github.com/indrasaputra)

## Usage

Send exactly two strings separated by a space to [@tetesan_hujan_bot](t.me/tetesan_hujan_bot) in Telegram.
The first string represents the link and the second represents the collection/category in which the link will be saved in [Raindrop.io](https://raindrop.io/).

For example, the link is `https://queue.acm.org/detail.cfm?id=3197520` and collection is `learning`. Then, the message in Telegram will be:
![Tetesan Hujan Example](https://user-images.githubusercontent.com/4661221/110271587-efffc380-7ffa-11eb-830c-2b18d62133b5.png)

## Caveat

To use this bot, make sure you have the access to use the bot (it is set by `TELEGRAM_OWNER_ID` environment variable).
Since the original purpose is my own usage only, the [@tetesan_hujan_bot](t.me/tetesan_hujan_bot) is only available for me.
If you want to use the bot, please follow the [Development](#development) and [Deployment](#deployment) sections.

The bot is heavily depends on [Raindrop API](https://developer.raindrop.io/).
Thus, any latency will be depends on [Raindrop API](https://developer.raindrop.io/) latency.

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
    $ https://github.com/indrasaputra/tetesan-hujan.git
    ```

- Go to project folder

    Usually, it would be
    ```
    $ cd go/src/github.com/indrasaputra/tetesan-hujan
    ```

- Fill in the environment variables

    Copy the sample env file.
    ```
    $ cp env.sample .env
    ```
    Then, fill the values according to your setting in `.env` file.

- Download the dependencies

    ```
    $ make dep-download
    ```
    or run this command if you don't have `make` installed in your local.
    ```
    $ go mod download 
    ```

- Run the application

    ```
    $ go run cmd/bot/main.go
    ```

- Expose your localhost to the internet

    Usually, I use [https://ngrok.com/](https://ngrok.com/)
    ```
    $ ngrok http $PORT
    ```

- Send some messages to your bot in Telegram.

## Deployment

Currently, this project is deployed in [Heroku](https://www.heroku.com/).
The deployment process definiton is stated and ruled in [Github Actions](https://github.com/indrasaputra/tetesan-hujan/blob/main/.github/workflows/test-and-deploy.yml).