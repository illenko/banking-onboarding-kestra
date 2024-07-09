## Direct calls to Kestra API:

### Start onboarding:
```shell
curl --location 'http://localhost:8080/api/v1/executions/company.myteam/onboarding-agreement' \
--form 'firstName="John"' \
--form 'lastName="Kestra"' \
--form 'email="johnkestra@gmail.com"' \
--form 'city="Kharkiv"' \
--form 'accountType="current"' \
--form 'currency="UAH"' \
--form 'sessionId="1231243214112412412"'
```

### Get onboarding status:
```shell
curl --location 'http://localhost:8080/api/v1/executions/34WB0vw3SDMKtCUuW7VtA4'
```

### Sign agreement:
```shell
curl --location 'http://localhost:8080/api/v1/executions/company.myteam/onboarding-signature' \
--form 'initial_flow="34WB0vw3SDMKtCUuW7VtA4"' \
--form 'account_id="3534d4b1-8532-44a4-a497-db19b4bba7fb"' \
--form 'agreement_id="8c6431e5-ed54-49ff-b410-01b69119060d"' \
--form 'signature="test_signature"'
```