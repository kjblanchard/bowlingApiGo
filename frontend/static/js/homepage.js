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
// function loginForm(data) {
//     console.log("hi")
//     console.log(data)
// }

function handleSubmit(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const formProps = Object.fromEntries(formData);
    console.log(formProps)
    console.log(formData)
  }

document.getElementById("login").onclick = login;
document.getElementById("request").onclick = checkLogin;
// document.getElementById("loginsubmitbutton").onclick = loginForm;

const loginForm = document.getElementById("loginform");
loginForm.addEventListener("submit", handleSubmit)