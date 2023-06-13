// // // // // // //
// Constants
// // // // // // //
const usernameInputId = "username"
const usernameValidationId = "usernameValidation"
const passwordValidationId = "passwordValidation"
// Regex
const usernameRegex = /^\S{1,24}$/gm
// // // // // // //

function login(username, password) {
    fetch('http://bowling.supergoon.com:8000/api/v1/signin', {
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
    fetch('http://bowling.supergoon.com:8000/api/v1/welcome', {
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
    const formProperties = Object.fromEntries(formData);
    console.log(formProperties)
    let valid = formValidation(formProperties)
    if (valid) {
        login(formProperties.username, formProperties.password)
        // Make it invisible after logging in.
        // loginForm.style.display = "none"
    }
}


function formValidation(formData) {
    let username = formData.username
    let password = formData.password
    let consoleString = "Username: " + username + "\nPassword: " + password
    let usernameTest = usernameInputValid(username)
    if (!usernameTest) {
        return false
    }
    console.log(consoleString)
}

function usernameInputValid(usernameText) {
    let isMatched = usernameRegex.test(usernameText)
    if (!isMatched) {
        usernameSpan.innerHTML = "Invalid Username: less than 24 characters without spaces"
    }
    else {
        usernameSpan.innerHTML = ""
    }
    return isMatched
}

// // // //
// On page load
// // // //
const usernameInput = document.getElementById(usernameInputId);
const usernameSpan = document.getElementById(usernameValidationId)

document.getElementById("request").onclick = checkLogin;
const loginForm = document.getElementById("loginform");
loginForm.addEventListener("submit", handleSubmit)
usernameInputValid()
// // // //