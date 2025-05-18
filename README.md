# weather-api
Weather API Application for Genesis Software Engineering School
Allows users to subscribe for weather updates for a selected city.

## Demo
html-page-link
Api host https://weather-api-dojf.onrender.com

## How to run localy
1. Copy this repo
2. Create in repo root directory files `postgres.env` and `weather-api.env` with folowing variables
> DISCLAMER
>
> All credential below are fake, and used only for demonstration purpuse.
> For real testing replace them with your own
>
> Weather API key can be obtained from [www.weatherapi.com](https://www.weatherapi.com)
>
> For SMTP server can be used services like [Mailtrap](https://mailtrap.io/) or [Mailjet](https://www.mailjet.com/), if you want to send real emails
### postgres.env
```
POSTGRES_DB=test_database

POSTGRES_PASSWORD=test_password

POSTGRES_USER=test_user
```
### weather-api.env
```
DB_HOST=postgres
DB_PORT=5432
DB_USER=test_user
DB_PASSWORD=test_password
DB_NAME=test_database

API_KEY=example-api-key

SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USER=example_user
SMTP_PASSWORD=example_password
SMTP_FROM=email@example.com

HOST=http://localhost:8080
PORT=8080
```
3. Execute `docker-compose up`

## Technology used
1. Go
2. PostgreSQL
3. Docker & Docker compose
4. [Mailjet SMTP](https://www.mailjet.com/)
5. [Weather API](https://www.weatherapi.com)
6. [Render](https://render.com/)

## Implementation details
### API
The API fully corresponds with the provided `swagger.yaml`. For most endpoints where the response structure was not specified, a default response like `{"message": "message-content"}` was used.
### Web
The [Gin framework](https://gin-gonic.com/) was used for HTTP request handling, which speeds up development process. Each endpoint has separate handler, all registered in 'routes.go'
### Database
[GORM](https://gorm.io/index.html) was used for database connection and interaction, providing easy-to-use abstractions and migrations.
### Email
The standart `net/smtp` package was used for sending emails, allowing flexibility in choosing the email provider. Email templates are stored in separate files and embedded into the binary using Go\`s `embed` package.
### Cron
The package [github.com/robfig/cron](https://github.com/robfig/cron) is used for scheduling cron jobs. The cron service runs in its own goroutine with two jobs: one for hourly emails, and one for daily emails.
### Docker & Docker Compose
The `Dockerfile` implements a standard multi-stage build of Go application. The `docker-compose.yaml` describes two services: weather-api and postgres. There is a volume for persisting the database data and a helthcheckto ensure weather-api starts only after postgres is fully set-up and ready. (with just `depends-on: - postgres` it starts too early, can`t connect and exits)

## Potential improvments
1. Migrate swagger.yaml to OpenAPI 3.0 and specify schemas for all possible responses (it is not allowed to change contract in this task, but in general very useful)
2. Add proper logging with tags and importancy levels (can be done with modules like https://github.com/rs/zerolog)
3. Test: all kind of test, from unit, to integrational and e2e.
4. CI/CD for building Docker image and running tests. Now it is partialy implemented with Render`s build in solution, but could be improved significantly with Github Actions, for example.
5. Proper git workflow: since it is a test project, and I am an only developer, it is not so important, but in case of scaling it is better to have branch for each feature and pull requests with code reviews.
6. Implement email queuing with retries to improve delivery reliability.
7. Add expiration for confirmation tokens and a cleanup job to remove unconfirmed subscriptions from the database.