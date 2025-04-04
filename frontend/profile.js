// profile.js
document.addEventListener('DOMContentLoaded', async () => {
    const usernameDisplay = document.getElementById('usernameDisplay')
    const userIdDisplay = document.getElementById('userIdDisplay')
    const logoutButton = document.getElementById('logoutButton')
    const backendUrl = 'http://localhost:3000'

    // 1. Verificar si tenemos un userId guardado
    const userId = localStorage.getItem('userId')
    const username = localStorage.getItem('username') // Recuperamos username también

    if (!userId) {
        // Si no hay ID, no está "logueado", redirigir a index.html
        alert('No han iniciado sesión. Redirigiendo a login...')
        window.location.href = 'index.html'
        return // Detener ejecución
    }

    // 2. Mostrar la información básica que ya tenemos
    userIdDisplay.textContent = userId
    usernameDisplay.textContent = username || 'Usuario' // Usar el guardado o un default

    // (Opcional) Podríamos verificar con el backend si el usuario aún existe
    // try {
    //     const response = await fetch(`${backendUrl}/users/${userId}`)
    //     if (!response.ok) {
    //         throw new Error('Usuario no encontrado en backend')
    //     }
    //     const userData = await response.json()
    //     usernameDisplay.textContent = userData.username // Actualizar con dato del backend
    // } catch (error) {
    //     console.error("Error verificando usuario:", error)
    //     alert('Error al verificar sesión. Redirigiendo a login.')
    //     localStorage.removeItem('userId')
    //     localStorage.removeItem('username')
    //     window.location.href = 'index.html'
    //     return
    // }


    // 3. Configurar el botón de Logout
    logoutButton.addEventListener('click', () => {
        // Limpiar localStorage
        localStorage.removeItem('userId')
        localStorage.removeItem('username')
        // Redirigir a login
        alert('Sesión cerrada.')
        window.location.href = 'index.html'
    })
})
