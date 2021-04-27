# Services Pong
Servicio **pong** es un servicio que será deployado en kubernetes. 

## Pre-requisitos

Tener instalado: 
* Golang. 
* Descargar las dependencias. 

En la ruta donde está el proyecto 
```bash 
go mod download
```

## Ejecutar

Solo para ejecutar 
```bash 
go run pong.go
```

Para compilar en Windows
```bash 
go build -o path/pong.exe pong.go
```