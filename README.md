# API-Implementation-Exercise
API Implementation Exercise

Assumptions:
alert_id and alert_ts are assumed to be uniqie so they are generated by the system and hence removed from the payload.

Steps to run this application.

1. Save the .env.example as .env
Enter an API key in API_KEY=
This is the key you will need to send as a header "X-API-KEY" in every request.

2. Run "go mod tidy".
3. Run "go run *.go".
4. See file requests.txt for API docs.
5. Testing, in a separate terminal run "go test -v".