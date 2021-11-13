# El contenedor base es Alpine en su versión 3.14 (menos de 3MB)
# Con este se crea el contenedor con todo lo necesario para ejecutar los tests
FROM alpine:3.14

# Como root, actualizar todos los paquetes básicos de Alpine
RUN apk update && \
    apk upgrade &&

# Se cambia a un usuario que no sea root
USER ostfind

# Se copia el archivo que contiene las dependencias del host a dentro del docker
COPY go.mod .
# Se copia el archivo que contiene las instrucciones del gestor de tareas del host a dentro del docker
COPY Taskfile.yml .

# Se instala el lenguaje Go, el gestor de tareas y las dependencias
RUN apk add --no-cache go=1.17.3-r0 && \
    go install github.com/go-task/task/v3/cmd/task@latest && \
    task installdeps

# Ejecuta lo que se quiere cuando se inicia el contenedor, en este caso pasa los tests usando el gestor de tareas
CMD ["task", "test"]