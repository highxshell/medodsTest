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

response:
```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMwMzAyNzcsImd1aWQiOiIxMjM0NTY3OCJ9.3E7vOwbNNJFwhs6knuR3azATziT0fhPtWo4yXBix2MvBT-z1VW0hiwMr8Xc5n-soGgPSTxtycUAzEL-FoeMxzA",
    "refresh_token": "jROsEv1k0EjvsonzqtWqFW--KB62Hmpm0ai2s08ITNV_MouV5R-ZGiLX6VFxhFaNBE-tKrqKz9LH2Rz010WAA5FJlgGdeG6r7zUsL08QOA2rc1GcFOpci366UpQADM9FXsC3E8A2lCRhekqy5xLwPRCHa82jVB-4WM0MQF-DfzeKgOQdZHhAJTFToXy1zOaQAKAsJiPUWgWvE8kr7x7ZZJGsENRDO0nKypLkcukRCdrLanaBSwwFEO8NZlWgg2S83v_9XgyphRI="
}
```

GET request to 

```bash
localhost:4000/auth/refresh
```

with body:
```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMwMzAyNzcsImd1aWQiOiIxMjM0NTY3OCJ9.3E7vOwbNNJFwhs6knuR3azATziT0fhPtWo4yXBix2MvBT-z1VW0hiwMr8Xc5n-soGgPSTxtycUAzEL-FoeMxzA",
    "refresh_token": "jROsEv1k0EjvsonzqtWqFW--KB62Hmpm0ai2s08ITNV_MouV5R-ZGiLX6VFxhFaNBE-tKrqKz9LH2Rz010WAA5FJlgGdeG6r7zUsL08QOA2rc1GcFOpci366UpQADM9FXsC3E8A2lCRhekqy5xLwPRCHa82jVB-4WM0MQF-DfzeKgOQdZHhAJTFToXy1zOaQAKAsJiPUWgWvE8kr7x7ZZJGsENRDO0nKypLkcukRCdrLanaBSwwFEO8NZlWgg2S83v_9XgyphRI="
}
```

response:
```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMxODQ4OTMsImd1aWQiOiIxMjM0NTY3OCJ9.Ankdr5r7BJL8O3vzp_DUeljA5nEawDC7NHd6ntMdaPa4m3jbYBAJZGwxzQn7e8qd3ffBgZRIym-YvaczOURbhQ",
    "refresh_token": "TRyMbVDkEgcYq-fdhVYC_BhqAZ6GYLTxxqCnEvo6vfwWU1krwm8vsehYS3msDa9XYB4nMuSf4hqBzHIyFEwCBRDvGvHDh1I9lATx0WUudFyPNZEq0tEoLRaEC-zKcvX6WZBfGzSlmKTY4KTmUnws6Frtf6WhTZ2pMZf0gEdCEPLm6V5OHFtEWtXZkH033H8Gber2m8Oltb---V0BdnP_FL4QgJMZIqYIcosQKM-HlunuhlKBw2SbcHkYHJ6bWouQ9RCk-ZgXUDU="
}
```

GET request to 

```bash
localhost:4000/auth/refresh
```

with INCORRECT REFRESH TOKEN body:
```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMwMzAyNzcsImd1aWQiOiIxMjM0NTY3OCJ9.3E7vOwbNNJFwhs6knuR3azATziT0fhPtWo4yXBix2MvBT-z1VW0hiwMr8Xc5n-soGgPSTxtycUAzEL-FoeMxzA",
    "refresh_token": "jROsEv1k0EjvsonzqtWqFW--KB62Hmpm0ai2s08ITNV_MouV5R-ZGiLX6VFxhFaNBE-tKrqKz9LH2Rz010WAA5FJlgGdeG6r7zUsL08QOA2rc1GcFOpci366UpQADM9FXsC3E8A2lCRhekqy5xLwPRCHa82jVB-4WM0MQF-DfzeKgOQdZHhAJTFToXy1zOaQAKAsJiPUWgWvE8kr7x7ZZJGsENRDO0nKypLkcukRCdrLanaBSwwFEO8NZlWgg2S83v_9XgyphR="
}
```

response:
```json
{
    "message": "invalid token"
}
```


