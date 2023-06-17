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

function handleSubmit(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const formProperties = Object.fromEntries(formData);
    console.log(formProperties)
    let valid = formValidation(formProperties)
    if (valid) {
        login(formProperties.username, formProperties.password)
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
    let isMatched = testObjectValues.regex.test(stringText)
    testObjectValues.getSpan().innerHTML = (isMatched) ? testObjectValues.validString : testObjectValues.invalidString
    return isMatched
}

// // // //
// On page load
// // // //
const usernameSpan = document.getElementById(usernameValidationId)
const passwordSpan = document.getElementById(passwordValidationId)
const loginForm = document.getElementById("loginform");
loginForm.addEventListener("submit", handleSubmit)
// // // //