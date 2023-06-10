function login() {
    alert("external fn clicked");
    console.log("What")
    // fetch('http://localhost:8000/api/v1/signin', {
    //     method: 'POST',
    //     headers: {
    //         'Accept': 'application/json',
    //         'Content-Type': 'application/json'
    //     },
    //     body: JSON.stringify({ "username": "user1", "password": "password1" })
    // })
    //     .then(response => response.json())
    //     .then(response => console.log(JSON.stringify(response)))
}

document.getElementById("click").onclick = login;