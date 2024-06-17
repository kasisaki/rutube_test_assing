
# RUTUBE test assignment

*Read this in other languages: [Russian](README.md)*


Solution to a test problem from Rutube

The task is formulated as follows:

```
Write a service for birthday greetings
• The goal is convenient congratulations to employees
• Obtaining a list of employees in any way (api/ad ldap/direct registration)
• Authorization
• Ability to subscribe/unsubscribe from birthday alerts

• Notification about the BD of the person you are subscribed to

• External interaction (json ari or front or tg bot)
• In case of interaction via a telegram bot (creating a group and adding all subscribers to it)
• In case of interaction through the front, setting the notification time before the birthday by email:

• It is advisable to provide unit tests;

• It is mandatory to add a README file to the repository with information on how to run the application and how to run tests. And also describe the catalog system, what is located and where.
```
## Description
### Authorization:

Used JWT for tokens. The user submits a login and password and receives a token for further authorization.
### Employee management:

API for adding employees, where their name, email and birthday are indicated.
### Alert subscriptions:

The user can subscribe to an employee birthday alert by specifying the number of days before the event to receive the notification.
The user can unsubscribe from the notification.
### Alerts:

A background process that checks employee birthdays daily and sends email alerts to subscribed users.

## Run Locally

Clone the project

```bash
  git clone https://github.com/kasisaki/rutube_test_assing
```

Go to the project directory

```bash
  rutube_test_assing
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run main.go
```


## Running Tests

To run tests, run the following command

```bash
  go test
```

