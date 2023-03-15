# Telegram Bot 
## Описание 
Для начала укажите токен бота в конфиге(config/config.yaml)  
Затем начните тестировать бота.  
## Пример работы 
Бот поддерживает две команды  
```
/find [argument]. Example: /find weather
/stats - get your personal stats
```
В ТЗ было указано, что можно взять любой источник информации я же взял ```https://api.publicapis.org/```   
Юзер отправляет какую категорию он хочет найти например погода ```/find weather``` вот ответ бота  
```
API Name: 7Timer!
Description: Weather, especially for Astroweather
Auth: key
Https: false
Cors: unknown
Link: http://www.7timer.info/doc.php?lang=en
Category: Weather
```
Получить статистику ```/stats```  
```
First request time: 2023-03-15 14:11:21.174458 +0000 +0000
Requests count: 1
```
