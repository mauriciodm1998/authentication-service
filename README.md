# HACKATON - Software Architecture - Auth Service

## Description

Este serviço é responsável por gerar tokens para ter autorização para enviar solicitações a outros serviços. Neste processo, este serviço recebe o login inserido e busca o usuário no banco de dados, após validação retorna o token ou um erro. É possível se logar tanto por username quanto pela matrícula.

## Features

- Generate Token


## How To Run Locally

```shell
make local-run
```

### VSCode - Debug
The launch.json file is already configured for debuging. Just hit F5 and be happy.

