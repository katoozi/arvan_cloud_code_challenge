# Wallet Service

## this service will handle the users wallets

### store and retrieve the transactions and current balance of their wallets

#### RUN

```bash
    go run cmd/wallet/wallet.go
```

#### Requests

##### start a trx

```bash
    curl --url 'http://localhost:8081/trx?cellphone=09017806181&amount=45000&timestamp=2021-12-12T12:12:12Z'
```

##### Get Balance

```bash
    curl --url 'http://localhost:8081/balance?cellphone=09017806181'
```
