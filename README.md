本项目使用[GitHub - 99designs/gqlgen: go generate based graphql server library](https://github.com/99designs/gqlgen)

参考GraphQL 

```
query getUser {
  user(id: "1000"){
    id
    username
    email
  }
}
query getTask {
  task(id:"1000"){
    title
    content
    status
  }
}

query getTasks {
  tasks{
    title
    content
  }
}

mutation createTask {
  createTask(title: "test title", content: "test content", status: true) {
    id
    title
    content
    status
  }
}

mutation updateTask {
  updateTask(id:"1000" title: "test title updated", content: "test content updated", status: true) {
    id
    title
    content
    status
  }
}

mutation register {
  register(userName: "jiuxia211", password: "123456", email: "2064166368@qq.com") {
    id
    username
  }
}

mutation login {
  login(userName: "jiuxia211", password: "123456") {
    username
    token
  }
}
```

