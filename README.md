# SpacedLearn

## Descripción

**SpacedLearn** es una aplicación CLI de repetición espaciada diseñada para ayudarte a aprender y recordar palabras de manera eficiente. Utiliza técnicas de repetición espaciada para mejorar la retención a largo plazo de las palabras, presentando palabras más difíciles con mayor frecuencia y palabras fáciles con menos frecuencia.

Este proyecto fue desarrollado de manera individual como parte de un reto de programación, con base en una serie de historias de usuario que guían el desarrollo de las funcionalidades clave.

## Características

1. **Registro y Login**: 
    - Los usuarios pueden registrarse con un nombre de al menos 5 caracteres y luego iniciar sesión.
    - Una vez registrados, pueden acceder directamente al inicio de sesión en futuras ejecuciones.

2. **Aprendizaje de Palabras**:
    - El usuario es presentado con una palabra, puede escribir su significado y recibir retroalimentación inmediata sobre su respuesta.
    - Si el usuario responde correctamente, se registra su progreso. Si responde incorrectamente, el sistema muestra la respuesta correcta.

3. **Progreso del Usuario**:
    - El progreso del usuario se guarda durante la sesión y se le presenta una puntuación de acuerdo a sus respuestas correctas.
    - Las palabras se presentan en base al algoritmo de repetición espaciada, mostrando con mayor frecuencia las palabras que el usuario no ha memorizado correctamente.

4. **Sesiones Continuas**:
    - El usuario puede continuar respondiendo palabras indefinidamente o elegir finalizar la sesión en cualquier momento.
    - Al finalizar la sesión, el sistema limpia los datos de la sesión actual y muestra un mensaje de finalización.

## Historias de Usuario

### 1. Como usuario, quiero conocer el nombre y el propósito de la aplicación de repetición espaciada.

**Criterios de aceptación**:
- Al ejecutar la CLI de SpacedLearn, el usuario verá el nombre de la aplicación y una breve descripción de lo que hace.

---

### 2. Como usuario, quiero crear un registro para poder usar la aplicación.

**Criterios de aceptación**:
- Al iniciar SpacedLearn por primera vez, se le pedirá al usuario registrarse con un nombre de al menos 5 caracteres.
- Si la información es correcta, el usuario será dirigido a la pantalla de login.
- Si hay un error, se mostrará un mensaje adecuado y se le pedirá que ingrese la información correcta.

---

### 3. Como usuario, quiero poder iniciar sesión para usar la aplicación.

**Criterios de aceptación**:
- El usuario puede omitir el registro si ya tiene una cuenta y proceder directamente a iniciar sesión.
- Si el nombre existe, la app le dará la bienvenida con un mensaje.
- Si el nombre es incorrecto, la app lo rechazará con un mensaje claro para que lo intente nuevamente o regrese al registro.

---

### 4. Como usuario, quiero aprender el significado de una palabra.

**Criterios de aceptación**:
- Al iniciar sesión, el usuario verá una palabra que quiere aprender.
- Puede escribir el significado y enviar su respuesta presionando Enter.

---

### 5. Como usuario, quiero recibir retroalimentación de mis respuestas para aprender correctamente.

**Criterios de aceptación**:
- Después de enviar su respuesta, el usuario recibe retroalimentación indicando si su respuesta fue correcta o incorrecta.
- Si la respuesta es incorrecta, se mostrará la respuesta correcta para aprenderla.

---

### 6. Como usuario, quiero avanzar a la siguiente palabra para seguir aprendiendo.

**Criterios de aceptación**:
- El usuario puede pasar a la siguiente palabra después de recibir la retroalimentación.

---

### 7. Como usuario, quiero que mi progreso se registre para saber qué palabras he aprendido o no.

**Criterios de aceptación**:
- Cada vez que el usuario acierte, su progreso será registrado.
- Si se equivoca, su puntuación no cambiará.

---

### 8. Como usuario, quiero ver las palabras basadas en el algoritmo de repetición espaciada.

**Criterios de aceptación**:
- Si el usuario acierta una palabra, esta volverá a aparecer más tarde.
- Si se equivoca, la palabra aparecerá con mayor frecuencia.

---

### 9. Como usuario, quiero seguir aprendiendo tantas palabras como quiera.

**Criterios de aceptación**:
- El usuario puede seguir avanzando indefinidamente entre palabras hasta que decida finalizar.
- Si no elige finalizar, la sesión continuará.

---

### 10. Como usuario, quiero ver mi progreso general para saber qué palabras necesito aprender más.

**Criterios de aceptación**:
- El usuario puede revisar su progreso y ver qué palabras ha acertado más y cuáles necesita practicar más.

---

### 11. Como usuario, quiero poder finalizar la sesión y limpiar los datos.

**Criterios de aceptación**:
- Cuando el usuario elija finalizar, los datos de la sesión actual se limpiarán y se mostrará un mensaje de finalización.
