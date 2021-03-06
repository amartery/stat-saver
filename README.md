# **Сервер для сбора статистики** 
# Запуск   
sudo docker-compose up  
# Тестирование  
## С помощью Insomia (Open Source API Client)  
в репозитории есть файл **insomnia_avito**  
его можно импортировать в локально установленную Insomia,  
и с помощью GUI потестировать проект  

## С помощью CURL  
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

# Выполнено  
- delivery, usecase, repository покрыты тестами (gomock, sqlmock)
- данные хранятся во внешнем хранилище (PostgreSQL)  
- в методе показа статистики можно выбрать сортировку по полям: 
event_date, views, clicks, cost, cpc, cpm    
- есть документация swagger (файл api-swagger.yaml)  
- настроил CI (Github Actions)