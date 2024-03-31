# fiber-authentication

API REST para manejar la autenticación de usuarios, implementando funcionalidades de registro (signup) y inicio de sesión (login) utilizando el framework Fiber para gestionar las solicitudes HTTP. La base de datos PostgreSQL se integra mediante GORM, y el paquete validator se utiliza para validar los datos de entrada.

## Características principales

- **Registro de usuarios:** Permite a los usuarios crear nuevas cuentas proporcionando un nombre de usuario, correo electrónico y contraseña.
- **Inicio de sesión:** Permite a los usuarios iniciar sesión con sus credenciales previamente registradas.
- **Hash de contraseñas:** Las contraseñas se almacenan en la base de datos utilizando técnicas de hash para mayor seguridad.
- **JWT (JSON Web Tokens):** Se emiten tokens JWT una vez que los usuarios inician sesión para autenticar y autorizar las solicitudes posteriores.
- Gestión de sesiones mediante **cookies** para mantener la autenticación del usuario.

## Requisitos previos

- Go 1.21 o superior
- Docker para ejecutar los servicios deseados, base de datos PostgreSQL
- **Paquetes Go:** Fiber, GORM, validator

## Instalación y Uso

1. Clonar el repositorio:

   ```bash
   git clone https://github.com/Mayer-04/fiber-authentication.git
   ```

2. Instalar las dependencias:

    ```bash
   go mod tidy
   ```

3. Clonar el archivo **.env.example** a **.env** para configurar las variables de entorno. Credenciales de la base de datos y clave secreta JWT
4. Configurar el **docker-compose.yml** y ejecutar:

   ```bash
   docker-compose up -d
   ```

5. Ejecutar la aplicación:

    ```bash
   go run cmd/main.go
   ```

6. Accede a la API desde: `http://localhost:5000/api/auth/[register|login]`