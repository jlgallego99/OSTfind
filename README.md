[![Build Status](https://app.travis-ci.com/jlgallego99/OSTfind.svg?branch=main)](https://app.travis-ci.com/jlgallego99/OSTfind)

# OSTfind
Sistema colaborativo en el cual los usuarios pueden buscar dónde y cuándo suena una canción de una banda sonora de una película, serie o videojuego, y recibir recomendaciones en base a sus preferencias. Desarrollado como proyecto de la asignatura Cloud Computing del Máster en Ingeniería Informática. 

## Descripción del problema
[Descripción](./docs/descripcion.md) más completa del problema a resolver en este proyecto.

## Lógica de negocio
Puedes leer sobre la lógica de negocio de este proyecto [aquí](./docs/logica_negocio.md).

## Desarrollo
- Mediante [historias de usuario](./docs/hu.md) se define el funcionamiento de la aplicación.

- Estas historias de usuario definen una serie de [objetos de valor, entidades y agregados](./docs/ddd.md) según el domain driven design.

- En [user journeys](./docs/userjourney.md) se representan los roles de los usuarios de la aplicación.

- Los [milestones](./docs/milestones.md) indican los distintos productos mínimos viables que se alcanzarán en cada fase del desarrollo.

### Tareas
Las [tareas](./docs/tareas.md) de instalación, ejecución de tests, etc se hacen mediante el uso del gestor de tareas Task.

### Contenedores
Se ha creado un [contenedor](./docs/contenedores.md) Docker para aislar y ejecutar los tests en un entorno controlado. También se ha documentado la actualización automática del contenedor.

### Integración Continua
Se han configurado sistemas de [Integración Continua](./docs/ci.md) para la ejecución automática de los tests en la rama principal y asegurar la calidad del software.

### Framework web
:exclamation: Se han estudiado distintos [frameworks web](./docs/web_frameworks.md) para crear un API REST en Go mediante la cual acceder por internet a los recursos de la lógica de negocio. También se ha documentado la elección de la biblioteca para los tests de integración del API REST y la configuración distribuida.

## Documentación adicional
:exclamation: Se ha seguido una [guía de buenas prácticas para estructurar aplicaciones web en Go](https://github.com/bnkamalesh/goapp). Se documentan las decisiones tomadas para la [estructura del proyecto](./docs/estructura.md).

Se ha documentado la [configuración inicial de Git](./docs/configuracion_entorno.md).

Estudio y elección de [framework y bibliotecas de testing en Go](./docs/test_frameworks.md).

Estudio y elección de [gestores de tareas](./docs/gestores_tareas.md).