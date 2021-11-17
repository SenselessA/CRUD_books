# CRUD-приложение, предоставляющее Web API к данным

## Запуск
Конфиг с настройками порта и БД в ```configs\config.yaml```
Пароль для бд хранится отдельно от репозитория на локальной машине, запись вида ``` DB_PASSWORD=your_password ``` в корневой папке проекта в файле ```.env```
Swagger doc ```http://localhost:8000/swagger/index.html```

###### SQL Скрипт для создания таблицы book в postgreSQL: 
```
  -- Table: public.book

  -- DROP TABLE public.book;

  CREATE TABLE IF NOT EXISTS public.book
  (
      title text COLLATE pg_catalog."default" NOT NULL,
      isbm character varying(32) COLLATE pg_catalog."default" NOT NULL,
      book_id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 )
  )
  WITH (
      OIDS = FALSE
  )
  TABLESPACE pg_default;

  ALTER TABLE public.book
      OWNER to postgres;
```

## API

##### http://localhost:8000/books
* GET - Получить все книги
Запрос: -  
Ответ JSON вида:
```
  [
      {
          "Id": 1,
          "Title": "The Diary of a Young Girl",
          "Isbm": "0199535566"
      },
      ...
  ]
```

##### http://localhost:8000/book/1
* GET - Получить книгу по id 1
Запрос path: http://localhost:8000/book/:id  
Ответ JSON вида:
```
  {
    "Id": 1,
    "Title": "The Diary of a Young Girl",
    "Isbm": "0199535566"
  }
```

##### http://localhost:8000/book
* POST - Добавить новую книгу  
Запрос JSON вида: 
```
  {
    "Title": "New book 3",
    "Isbm": "0654634534566"
  }
``` 
Ответ JSON вида:
```
  {
    "id": 14
  }
```

##### http://localhost:8000/book
* PUT - Обновить данные по книге  
Запрос JSON вида: 
```
  {
    "Id": 7,
    "Title": "New book 2 Updated",
    "Isbm": "023233253445"
  }
``` 
Ответ JSON вида:
```
  {
    "Id": 7,
    "Title": "New book 2 Updated",
    "Isbm": "023233253445"
  }
```

##### http://localhost:8000/book
* DELETE - Удалить книгу  
Запрос JSON вида: 
```
  {
    "Id": 14
  }
``` 
Ответ JSON вида:
```
  {
    "Id": 14,
    "Title": "New book 3",
    "Isbm": "0654634534566"
  }
```
