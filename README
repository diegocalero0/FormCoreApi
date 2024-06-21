# Teamform API - Una API de prueba para Teamcore

Este proyecto realizado en Go es una API con un único método **/questions**

Este método se encarga de consultar el endpoint proporcionado para la prueba y realizar una transformación devolviendo la información como se muestra a continuación:

```json
{
    "titulo": "Preguntas sobre cultura general",
    "dia": "12-01-2022",
    "info": [
        {
            "pregunta_id": 1,
            "pregunta": "Cuál es uno de los fundadores de apple?"
        },
        {
            "pregunta_id": 2,
            "pregunta": "¿Cuál es uno de los fundadores de Microsoft?"
        },
        {
            "pregunta_id": 3,
            "pregunta": "¿Indique quien jugó en el equipo de chicago bulls?"
        }
    ],
    "api_version": 1
}
```

## ❗⚠️ Variables de entorno ❗⚠️

Para poder desplegar el proyecto en Cloud Run es importante configurar las variables de entorno, esto se realizó por motivos de seguridad y no exponer la información suministrada a internet.

### Configurar .env
Para poder correr y probar el proyecto en Google Cloud se requiere crear en la raíz del proyecto un archivo **.env** con 2 variables de entorno: **TEAMCORE_API_URL** y **TEAMCORE_API_TOKEN**, estos valores se encuentran en el Email enviado.

❗IMPORTANTE: La variable **TEAMCORE_API_URL** debe finalizar en **/api** y NO debe contener el **/questions** ya que esto se configura como ruta.

## Subir proyecto a Cloud Run 
El proyecto fue desarrollado de tal manera que sea fácil desplegar en Cloud Run ya que se encuentra Dockerizado.

### Requisitos para subir a Cloud Run
* Un proyecto con la facturación habilitada (De lo contrario podrían obtener errores al publicar el artefacto).
* Un usuario con acceso al proyecto.

### Autenticar
Este paso es opcional, se puede omitir en caso de que ya se tenga configurada la autenticación y el proyecto al cual se desea subir el proyecto.

Para autenticar puede usar el siguiente comando:

```bash
gcloud auth login
```

Una vez autenticado puede cambiar de proyecto usando el siguiente comando: 

```bash
gcloud config set project PROJECT_ID
```
Donde **PROJECT_ID** corresponde al id del proyecto en Google Cloud.

## Construir, publicar y correr imagen
Para el proceso de construir, publicar y correr la imagen en Cloud Run se hará uso de los siguientes comandos:
### Construir y publicar
```bash
gcloud builds submit -t gcr.io/{PROJECT_ID}/teamform-api:v1 .
```
No olvide reemplazar {PROJECT_ID} por el id de su proyecto en Google Cloud.

### Correr imagen
Para correr la imagen se hará uso del siguiente comando:
```bash
gcloud run deploy teamform-api --image gcr.io/{PROJECT_ID}/teamform-api:v1 --port 3000
```
No olvide reemplazar {PROJECT_ID} por el id de su proyecto en Google Cloud.

### Probar la ejecución del proyecto
Si la ejecución de los 2 comandos anterior fue exitosa, tendrá una respuesta similar a la siguiente:

```
Deploying container to Cloud Run service [teamform-api] in project [tinpos-project] region [us-central1]
✓ Deploying... Done.                                                                                                                        
  ✓ Creating Revision...                                                                                                                    
  ✓ Routing traffic...                                                                                                                      
Done.                                                                                                                                       
Service [teamform-api] revision [teamform-api-00014-9n8] has been deployed and is serving 100 percent of traffic.
Service URL: https://teamform-api-juydqhgrqq-uc.a.run.app
```
Siendo el Service URL, la URL en la cual puede probar la ejecución del programa.

La ruta que se debe probar es **/questions**.
Ejemplo: https://teamform-api-juydqhgrqq-uc.a.run.app/questions

❗IMPORTANTE: Esta API al igual que la API original, tiene autenticación por Token, para probar la API se debe especificar el token de autenticación Bearer enviado al correo. ESTE TOKEN ES DIFERENTE AL ORIGINAL.

## License

[MIT](https://choosealicense.com/licenses/mit/)