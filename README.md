[![autotests](https://github.com/lastbyte32/go-metric/actions/workflows/devopstest.yml/badge.svg?branch=iter1)](https://github.com/lastbyte32/go-metric/actions/workflows/devopstest.yml)

[![go vet test](https://github.com/lastbyte32/go-metric/actions/workflows/statictest.yml/badge.svg?branch=iter1)](https://github.com/lastbyte32/go-metric/actions/workflows/statictest.yml)

# go-musthave-devops-tpl


# Обновление шаблона

Чтобы получать обновления автотестов и других частей шаблона, выполните следующую команду:

```
git remote add -m main template https://github.com/yandex-praktikum/go-musthave-devops-tpl.git
```

Для обновления кода автотестов выполните команду:
```
git fetch template && git checkout template/main .github
```
