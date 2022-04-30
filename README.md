# Temporal-iamBinding-Google-workflow

`DESCRIPTION`

This project is used to create an IAM binding in GCP with specific Role.
This project uses temporal and Cloud SDK's for IAM Binding.
This project uses  Gin FrameWork.


```shell
docker-compose -f .local/quickstart.yml up --build --force-recreate -d
export TEMPORAL_HOSTPORT=localhost:7233
export GOOGLE_APPLICATION_CREDENTIALS={{path of your SPN File}}
go run main.go
```

The Application by default runs in port `http://localhost:8080`
The Temporal UI runs in port `http://localhost:8088`
