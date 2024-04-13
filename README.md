# Instalation

```bash
go mod tidy
```

```bash
cd .\services\jwt\cmd\app\
```

```bash
go run main.go
```

GET request to 

```bash
localhost:4000/auth/login?guid=12345678
```

```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMwMzAyNzcsImd1aWQiOiIxMjM0NTY3OCJ9.3E7vOwbNNJFwhs6knuR3azATziT0fhPtWo4yXBix2MvBT-z1VW0hiwMr8Xc5n-soGgPSTxtycUAzEL-FoeMxzA",
    "refresh_token": "jROsEv1k0EjvsonzqtWqFW--KB62Hmpm0ai2s08ITNV_MouV5R-ZGiLX6VFxhFaNBE-tKrqKz9LH2Rz010WAA5FJlgGdeG6r7zUsL08QOA2rc1GcFOpci366UpQADM9FXsC3E8A2lCRhekqy5xLwPRCHa82jVB-4WM0MQF-DfzeKgOQdZHhAJTFToXy1zOaQAKAsJiPUWgWvE8kr7x7ZZJGsENRDO0nKypLkcukRCdrLanaBSwwFEO8NZlWgg2S83v_9XgyphRI="
}
```

