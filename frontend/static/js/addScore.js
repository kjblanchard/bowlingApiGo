// // // // // // //
// Constants
// // // // // // //
const addScoreFormId = "addScore"
const scoreValidationId = "scoreValidation"
// // // // // // //


function handleAddScoreSubmit(event) {
    event.preventDefault();
    const formData = new FormData(event.target);
    const formProperties = Object.fromEntries(formData);
    let score = formData.score
    let valid = addScoreSubmitValidation(score)
    if (valid) {
        addScore(score)
    }
}

function addScoreSubmitValidation(score) {
    let valid = (score >= 0 && score <= 300)
    addScoreSpan.innerHTML = (valid) ? "" : "Invalid Score: 1-300"
    return valid
}

function addScore(score, userId) {
    // fetch('http://bowling.supergoon.com:8000/api/v1/users/' + userId + '/scores', {
    //     method: 'POST',
    //     headers: {
    //         'Accept': 'application/json',
    //         'Content-Type': 'application/json'
    //     },
    //     credentials: 'include',
    //     body: JSON.stringify({ "score": score }),
    // })
}

// // // //
// On page load
// // // //
const addScoreSpan = document.getElementById(scoreValidationId)
const addScoreLoginForm = document.getElementById(addScoreFormId);
addScoreLoginForm.addEventListener("submit", handleAddScoreSubmit)
// // // //