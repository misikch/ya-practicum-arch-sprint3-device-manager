# ya-practicum-arch-sprint3-device-manager
Device Manager service mvp

### Как устроен сервис
Все доступные api methods описаны в `api.yaml` файле. После изменнеия файла нужно запустить кодогенерацию:
```shell
go generate ./...
```

### Как локально протестировать
Запустить окружение
```shell
docker-compose up
```

Подготовить переменные окружения:
```shell
cp .env.example .env
```

Запустить приложение из директории `/cmd/service`.

### Как остановить и почистить окружение
```shell
docker-compose down --remove-orphan
```
