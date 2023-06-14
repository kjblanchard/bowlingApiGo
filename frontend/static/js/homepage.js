// // // // // // //
// Constants
// // // // // // //
const usernameInputId = "username"
const usernameValidationId = "usernameValidation"
const passwordValidationId = "passwordValidation"
/**
 * Validation function object, used when iterating through the things we need regex for.
 */
const validationMap = {
    username: {
        regex: /^\S{1,24}$/,
        validString: "",
        invalidString: "Invalid Username, 1-24 characters and no spaces",
        getSpan: () => {return usernameSpan}
    },
    password: {
        regex: /^\S{1,24}$/,
        validString: "",
        invalidString: "Invalid Password, 1-24 characters and no spaces",
        getSpan: () => {return passwordSpan}
    }
}
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
    let usernameTest = regexTest(username, "username")
    let passwordTest = regexTest(password, "password")
    if (usernameTest && passwordTest) {
        return true
    }
    return false
}

function regexTest(stringText, testKey)  {
    testObjectValues = validationMap[testKey]
    console.log("Testing value: " + stringText)
    let isMatched = testObjectValues.regex.test(stringText)
    console.log("Match: ", + isMatched)
    testObjectValues.getSpan().innerHTML = (isMatched) ? "" : "Invalid Password: 1-24 Chars without spaces"
    return isMatched
}

// // // //
// On page load
// // // //
const usernameSpan = document.getElementById(usernameValidationId)
const passwordSpan = document.getElementById(passwordValidationId)

document.getElementById("request").onclick = checkLogin;
const loginForm = document.getElementById("loginform");
loginForm.addEventListener("submit", handleSubmit)
// // // //