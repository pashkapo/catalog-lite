# catalog-lite
Приложение доступно по [ссылке](https://catalog-lite.herokuapp.com/api/firms). Первый запрос может выполняться не быстро, сервис отключается после 30m простоя и стартует при обращении

## Общая информация
 - [описание api](https://pashkapo.github.io/catalog-lite/#/)
 
## Структура проекта
- `config` - переменные окружения и константы
- `db` - взаимодействие с бд
- `dist` - папка со статикой для отображения swagger.ui
- `handler` - обработчики
- `model` - модели
- `resources` - дамп бд

## Запуск dev окружения
```bash
docker-compose up -d --build
```
