# Discount Service

## Discount service is a stateful service.

## we can make it stateless by implementing a new storage for it and inject it.

### It handles the voucher redemptions

### Run

```sh
    go run cmd/discount/discount.go
```

### Endpoints

#### *please also run wallet service on port 8081 for handling transactions*

#### Redeem a voucher by user

    /redeem?voucher-code=test&cellphone=09017806181


#### creates a voucher in storage

    /voucher?voucher-code=test&max-use-time=200&discount-price=1200000

#### get the list of redemptions for a specific voucher

    /voucher/redemptions?voucher-code=test

##### examples

```bash
    curl --url 'http://localhost:8080/redeem?voucher-code=test&cellphone=09017806181'

    curl --url 'http://localhost:8080/voucher?voucher-code=test&max-use-time=200&discount-price=1200000'

    curl --url 'localhost:8080/voucher/redemptions?voucher-code=test'
```
