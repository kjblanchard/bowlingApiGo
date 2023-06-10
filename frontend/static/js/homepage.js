function login(username, password) {
    fetch('http://localhost:8000/api/v1/signin', {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ "username": username, "password": password }),
    })
    .then(response => document.cookie = response.cookie)
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

function handleSubmit(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const formProps = Object.fromEntries(formData);
    formValidation()
    login(formProps.username, formProps.password)
    loginForm.style.display = "none"
  }

function formValidation() {

}

document.getElementById("request").onclick = checkLogin;
const loginForm = document.getElementById("loginform");
loginForm.addEventListener("submit", handleSubmit)