## HowTo

You need to source .envrc: `source .envrc`

To launch graphql: `make graphql`

To launch project microservice: `make project`

To launch user microservice: `make user`

Then go to `localhost:8080` and make queries.

Available queries are:

```graphql
query {
    allProjects {
        id
        name
    }
}
```

```graphql
query {
    allUsers {
        id
        email
    }
}
```
