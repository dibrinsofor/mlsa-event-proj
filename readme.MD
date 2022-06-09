# mlsa event 3

a quick user onboarding flow written in Go (Gin) to demonstrate some use cases for queueing systems in data intensive applications

## key features

This go app:

- uses redis as a backend data store to persist user info
- leverages queues (azure service bus) to handle sending of onboarding emails

## available endpoints and errors

- `/` needs jwt authorization else redirects to `/createUser`

- `/health` endpoint to ensure the server is up and running

- `/createUser` endpoint handles creating user accounts and processing onboarding emails to be sent via the azure service bus