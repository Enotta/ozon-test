# Ozon-test

### Руководство пользователя

Данное руководство покрывает все функции разработанной системы

0. Собрать проект при помощи docker-compose

Конфигурация для хранения данных в БД Postgres:
```console
docker-compose -f .\docker-compose-postgres.yml up
```

Конфигурация для хранения данных локально:
```console
docker-compose -f .\docker-compose-inmemory.yml up
```

1. Перейти на localhost:8080
2. Начать посылать graphql запросы

Первым делом (если конечно мы не хотим проверить обработку ошибок) стоит создать автора:
```graphql
mutation {
  createAuthor(
    input: {
      username: "AuthorName"
    }
  ) {
    id,
    username
  }
}
```

Затем необходимо, чтобы нащ автор написал свой первый пост:
```graphql
mutation {
  createPost(
    input: {
      title: "Title",
      content: "Lorem Ipsum",
      author: "1",
      commentsEnabled: true
    }
  ) {
    id,
    title,
    createdAt
  }
}
```

К посту можно писать комментарии (parentId: null для корневых комментариев):
```graphql
mutation {
  createComment(
    input: {
      content: "Like!",
      author: "1",
      postId: "1",
      parentId: null
    }
  ) {
    id,
    content,
    createdAt
  }
}
```

3. Создать необходимое количество авторов, постов и комментариев
4. Написать запросы

Запрос на получение поста и корневых комментариев к нему
```graphql
query {
  post(id: "1") {
    title,
    content
  }
  
  comments(postId: "1", parentId: null) {
    id,
    content
  }
}
```

Запрос на получение всех постов
```graphql
query {
  posts {
    id,
    title,
    createdAt
  }
}
```

### Замечания:
Запросы к серверу не будут отличаться в зависимости от типа хранения данных
