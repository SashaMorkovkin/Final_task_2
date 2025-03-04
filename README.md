# Веб-калькулятор
## Структура проекта:
+ cmd/ - директория с файлом main.go(точка входа в программу)
+ internal/ - директория где храниться сервер
+ internal/api - Здесь находится обработчик программы
+ internal/calculator - здесь находится код для вычислений
+ internal/task_manager - это директория для работы с задачами
+ internal/worker - получает задачи от оркестратора
+ test/ - тесты программы
+ pkg/rpn/ - директория где храниться код калькулятора

## Запуск
1. Склонируйте проект с github git clone ```https://github.com/SashaMorkovkin/Final_task_2```
2. Перейдите в папку с проектом и запустите сервер ```go run ./cmd```
> PS: Сервер работает на порте 8080

## Примеры запросов
+ Пример №1
    + Команда:
        >curl --location 'http://localhost:8080/api/v1/calculate' \
        >--header 'Content-Type: application/json' \
        >--data '{"expression": "2+2*2"}'

    + Ответ:
        >{"id": "1","result": 6}
+ Пример №2
    + Команда:
        >curl --location 'http://localhost:8080/api/v1/expressions'
    + Ответ:
        >{"expressions":[{"ID":"1","Status":"completed","Result":6,"Expression":"ваш_пример"},{"ID":"2","Status":"completed","Result":6,"Expression":"ваш_пример"}]}
+ Пример №3
    + Команда:
        >curl --location 'http://localhost:8080/api/v1/expressions/1'
    + Ответ:
        > {"ID":"1","Status":"completed","Result":3,"Expression":"ваш_пример"}
+ Если калькулятор не может посчитать :
    >{"Error calculating expression"}
+ В других случаях :
    >{"Expression not found"}

>Так же ошибки пишутся более подробно в логи
