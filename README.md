Тестовое задание по рефакторингу от TrueConf

Что было сделано:

- Архитектура сервиса переработана так, чтобы соответствовать правилам "чистой архитектуры".
- Выделено 3 слоя (controller, usecases, repository). Благодаря инверсии зависимостей, слои не связаны друг с другом.
- Сервер вынесен в отдельный пакет, реализован graceful shutdown.
- Конфигурация подгружается из переменных окружения.
- Приложение обернуто в docker контейнер.

Как запустить:

- Клонировать репоситорий и перейти в него
- С помощью docker-compose (рекомендуемый):
    - собрать контейнер командой ```docker-compose up --build```
- С использованиеи bash скрипта (для linux):
    - подгрузить зависимости ```go mod download```
    - запусть скрипт ```./scripts/run_app.sh```

Запуск, по умолчанию, осуществляется по адресу http://localhost:8080
    

Что можно улучшить:

- Так как сериализация json процесс не быстрый, 
можно использовать сторонние библиотеки по сериализации json (например github.com/goccy/go-json).
Чтобы увеличить производительность, можно заменить хранилище на SQL (PostgreSQL, MariaBD, MySQL и др.) или noSQL (MongoDB, CouchDB и др.). 
Сделать это не сложно, благодаря разделению на слои, для этого нужно реализовать интерфейс Repository.
- Добавить валидацию данных в слое middleware.
- В результате работы могум возникать внутренние ошибки, слдует добавить логирование.
- Покрыть код тестами, для спокойствия.

Задание:

Вам предстоит выполнить рефакторинг небольшого приложения на Go (200 строк).

Приложение представляет собой API по работе с сущностью User, где хранилищем выступает файл json.

Ограничения:
- Хранилищем должен оставаться файл в json формате.
- Структура пользователя не должна быть уменьшена.
- Приложение не должно потерять существующую функциональность. 

Мы понимаем, что пределу совершенства нет и ожидаем, что объем рефакторинга вы определяете на свое усмотрение.  

После того как вы выполните задание, вы так же можете написать, как бы улучшили проект в перспективе текстом.

Что следует знать:
- В будущем это приложение ожидает увеличение количества функций и сущностей. 
- Вопрос авторизации умышленно опущен, о нем не стоит беспокоиться.
- API еще не выпущено, вы в праве скорректировать интерфейс / форматы ответов.

Работа должна быть оформлена на Github.
