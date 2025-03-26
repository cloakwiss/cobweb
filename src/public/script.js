document.getElementById("archiveForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Prevent form submission

    let a = parseInt(document.getElementById("inputA").value);
    let b = parseInt(document.getElementById("inputB").value);

    sendData(a, b);
});

function sendData(a, b) {
    console.log("Sending data:", { A: a, B: b });

    fetch('/archive', {  // Updated route
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ A: a, B: b })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById("message").innerText = "Response: " + data.message;
    })
    .catch(error => {
        console.error('Error:', error);
        document.getElementById("message").innerText = "An error occurred.";
    });
}
