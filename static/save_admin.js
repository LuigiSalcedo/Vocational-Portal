const saveAdmin = () => {
    let _name = document.getElementById("name_input").value
    let _password = document.getElementById("password_input").value

    const data = {
        name: _name,
        password: _password
    } 

    fetch('/admins/guardar', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }).then(response => {
        if(response.status != 201){
            alert('Error al guardar el administrador')
        } else {
            alert('Administrador guardado con éxito')
        }
    })
}

let button = document.getElementById('save_button')

button.onclick = saveAdmin