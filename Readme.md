# Proyecto Golang

## 📄 Descripción del Proyecto

Este es un proyecto desarrollado en **Go** que proporciona una API para gestionar productos. El proyecto fue construido para ser escalable, con una arquitectura bien estructurada, utilizando **Docker** y siguiendo las mejores prácticas de ingeniería de software.

### 🛠️ Pasos para Configuración del Entorno

1. Cloná el repositorio:

   ```sh
   git clone git@github.com:daviresio/products_server.git
   cd products_server
   ```

2. Configurá el entorno:
   Utilizá el script `setup` que se encuentra en la carpeta `scripts` para configurar el entorno local. Este script hará lo siguiente:

   - Instalar **Homebrew**, **Docker**, **Go**, **Air** y **Migrate**, en caso de que no estén instalados.
   - Instalar las dependencias del proyecto usando `go mod tidy`.
   - El script está destinado solo para macOS y debe ejecutarse una sola vez.

   ```sh
   ./scripts/setup
   ```

3. Variables de Entorno

   Las variables de entorno necesarias deben ser cargadas directamente en el entorno del sistema, debés exportarlas para que estén disponibles en la terminal al ejecutar el script de migración o correr la aplicación. Ejemplos de variables incluyen:

   - `DB_HOST=localhost`
   - `DB_PORT=5432`
   - `DB_USER=postgres`
   - `DB_PASSWORD=postgres`
   - `DB_NAME=products_db`

4. Inicializá la aplicación usando **Docker** Compose:

   ```sh
   docker compose up --build
   ```

   Esto creará los contenedores necesarios para la aplicación y la base de datos.

5. Cargá las variables de entorno necesarias en el entorno de tu sistema y aplicá las migraciones de la base de datos:

   ```sh
   ./scripts/migration_up
   ```

6. La aplicación estará corriendo en el puerto `8081` (**Docker**). Una base de datos PostgreSQL también estará corriendo en el puerto `5432`. También podés correr la aplicación localmente usando el comando `air`, que la ejecutará en el puerto `8080`.

### 📊 Migraciones de la Base de Datos

Para aplicar las migraciones de la base de datos, utilizá el script `migration_up` que realizará las migraciones:

```sh
./scripts/migration_up
```

## 📌 Endpoints

- **Health Check**: `GET /health` - Verifica si el servicio está funcionando correctamente.
- **Productos**:
  - `GET /products?page=&search=` - Lista los productos disponibles, con soporte para paginación y filtrado por texto.
  - `GET /products/{id}` - Devuelve los detalles de un producto específico.

## 📜 Scripts

Los scripts disponibles ayudan en la configuración y desarrollo del proyecto:

- `create_migration`: Crea una nueva migración para la base de datos.
- `migration_up`: Aplica las migraciones pendientes a la base de datos.
- `setup`: Configura el entorno para el desarrollo local, instalando dependencias y preparando la base de datos.

## 🧾 Deuda Técnica

Actualmente, el proyecto tiene algunas áreas de deuda técnica que necesitan ser resueltas:

- Falta de pruebas automatizadas para la aplicación, lo que compromete la calidad del código.
- Falta de aplicación de las migraciones en la pipeline de CI, lo que puede causar inconsistencias en la base de datos durante el despliegue.
- Falta de implementación de observabilidad para monitorear la aplicación e identificar problemas en producción.
- Falta de carga automática del archivo `.env` al iniciar la aplicación para agilizar el desarrollo local.

## 🚀 CI/CD

El proyecto utiliza una pipeline de integración continua (CI) para garantizar la calidad del código. La pipeline automatizada ejecuta las siguientes etapas:

1. **Build**: Construye la aplicación y verifica posibles errores.
2. **Deploy**: Utiliza **Docker** para empaquetar la aplicación y realizar el despliegue.

La pipeline realiza un cálculo automático del versionado SemVer. Para esto, la rama debe apuntar a `main` y comenzar con `major/`, `feature/` o `bugfix/`. Después del merge, se calcula el SemVer, se genera la imagen y se envía al artifact de GCP, y la aplicación se despliega automáticamente en Google Cloud Run.

## Contacto 📧

Si tenés dudas o sugerencias, ponete en contacto al email: [daviresio@gmail.com](mailto:daviresio@gmail.com).
