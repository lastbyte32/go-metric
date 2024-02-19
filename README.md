[![autotests](https://github.com/lastbyte32/go-metric/actions/workflows/devopstest.yml/badge.svg?branch=iter8)](https://github.com/lastbyte32/go-metric/actions/workflows/devopstest.yml)
[![go vet test](https://github.com/lastbyte32/go-metric/actions/workflows/statictest.yml/badge.svg?branch=iter8)](https://github.com/lastbyte32/go-metric/actions/workflows/statictest.yml)
[![codecov](https://codecov.io/gh/lastbyte32/go-metric/branch/iter3/graph/badge.svg?token=JGW4NDIJR0)](https://codecov.io/gh/lastbyte32/go-metric)

# Сервис сбора метрик и алертинга

>Проект на курсе «Продвинутый Go‑разработчик» в Яндекс.Практикум 

## Описание 
Агент собирает runtime-метрики и передает их серверу используя **gRPC/REST** для хренения.

Сервер и агент конфигурируются через флаги запуска, переменные среды ОС и конфигурационный файл.

При передаче метрик реализована возможность проверки данных метрик с помощью цифровой подписи, а также применение шифрования для передачи данных.

Сервер может использовать оперативную память, файловое хранилище или базу данных **PostgreSQL** в качестве места для хранения данных.

### Использованы следующие технологии/библиотеки

* PostgreSQL (pgx/sqlx)
* gRPC
* Resty
* ZAP
* Chi
* Testify
