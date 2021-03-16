# Сервер для сбора статистики  
## Запуск  
sudo docker-compose build  
sudo docker-compose up  
## Тестирование  
### метод сохранения статистики  
```bash
curl --request POST \
  --url http://127.0.0.1:8080/stat/add \
  --header 'Content-Type: application/json' \
  --data '{
	"date":"2019-08-22",
	"views":"494",
	"clicks":"55556786839",
	"cost":"1.57"
}
```
### метод показа статистики без сортировки по конкретному полю  
```bash
curl --request GET \
  --url 'http://127.0.0.1:8080/stat/show?from=2000-01-01&to=2021-01-01'
```
### метод показа статистики с сортировкой по конкретному полю  
доступные поля для сортировки: event_date, views, clicks, cost, cpc, cpm  
```bash
curl --request GET \
  --url 'http://127.0.0.1:8080/stat/show?from=2000-01-01&to=2025-01-01&sort=clicks'
```
### метод удаления статистики  
```bash
curl --request DELETE \
  --url http://127.0.0.1:8080/stat/del
```

## Выполнено  
Данные сервиса хранятся во внешнем хранилище (PostgreSQL)  
В методе показа статистики можно выбрать сортировку по любому из полей ответа  
Код частично покрыт тестами (вся валидация)  
Для документации я использовал swagger  