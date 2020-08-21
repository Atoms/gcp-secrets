Read Google cloud secrets from Secret Manager using service account.

## HOWTO

* Create service account and download json key
* Add in IAM role `Secret Manager Secret Accessor` role
* export env variable `GOOGLE_APPLICATION_CREDENTIALS` with your service account key file

```shell
gcp-secret -p project-id -s secretname
```

## building
```
go build
```

## Currently supported

* Secret value for latest version of secret

## TODO:

* Add versioning support
* Add flag to list all secrets available
* Add possibility to save new version of secret
