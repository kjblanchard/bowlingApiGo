function login() {
    fetch('http://localhost:8000/api/v1/signin', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ "username": "user1", "password": "password1" }),
        credentials: 'include',
    })
    .then(response => document.cookie = response.cookie)
        // .then(response => response.json())
        // .then(response => console.log(JSON.stringify(response)))
}
function checkLogin() {
    fetch('http://localhost:8000/api/v1/welcome', {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        credentials: 'include',
    })

}

document.getElementById("login").onclick = login;
document.getElementById("request").onclick = checkLogin;