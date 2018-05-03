# Happy App Server
[![Circle CI](https://img.shields.io/circleci/project/github/UnalignedByte/happy-app-server.svg?style=plastic&label=CircleCI)](https://circleci.com/gh/UnalignedByte/happy-app-server)

This is the server for the Happy App project. It is written in Go.

## API
`/api/happiness (POST)`

```
{
  percentage: Int (0-100)
}
```
`/api/happiness (GET)`

```
{
  overallPercentage: Int (0-100)
}
```

## Author
The project has been created by Rafał Grodziński.