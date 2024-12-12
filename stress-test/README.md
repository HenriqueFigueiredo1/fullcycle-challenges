# Desafio goexpert Stress Test

1. Na pasta raiz do projeto

2. Compilar imagem docker
```shell
docker build -t st .
```

3. Executar a aplicação
```shell
docker run st:latest --url=http://google.com --requests=1000 --concurrency=10
```
