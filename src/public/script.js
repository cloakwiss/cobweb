document.getElementById("archiveForm").addEventListener("submit", function(event) {
	event.preventDefault(); // Prevent default form submission

	const options = {
		NoAudio: document.getElementById("NoAudio").checked,
		NoCss: document.getElementById("NoCss").checked,
		NoIframe: document.getElementById("NoIframe").checked,
		NoFonts: document.getElementById("NoFonts").checked,
		NoJs: document.getElementById("NoJs").checked,
		NoImages: document.getElementById("NoImages").checked,
		NoVideo: document.getElementById("NoVideo").checked,
		NoMetadata: document.getElementById("NoMetadata").checked,
		Targets: document.getElementById("Targets").value.trim(),
		AllowDomains: document.getElementById("AllowDomains").value.trim(),
		BlockDomains: document.getElementById("BlockDomains").value.trim(),
		Output: document.getElementById("Output").value.trim(),
		Cookie: document.getElementById("Cookie").value.trim(),
		Depth: parseUint8(document.getElementById("Depth").value),
		// Mode: document.getElementById("Mode").value.trim(),
		Timeout: parseUint64(document.getElementById("Timeout").value)
	};

	sendData(options);
});

function parseUint8(value) {
	let num = parseInt(value, 10);
	return isNaN(num) || num < 0 || num > 255 ? 0 : num;
}

function parseUint64(value) {
	let num = parseInt(value, 10);
	return isNaN(num) || num < 0 ? 0 : num;
}

function sendData(data) {
	console.log("Sending data:", JSON.stringify(data, null, 2));

	fetch('/archive', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(data)
	})
		.then(response => response.json())
		.then(data => {
			displayResponse(data);
		})
		.catch(error => {
			console.error('Error:', error);
			document.getElementById("message").innerHTML = `<p style="color: red;">An error occurred.</p>`;
		});
}

function displayResponse(data) {
	const responseInput = document.getElementById("message");
	const downloadButton = document.getElementById("downloadButton");

	responseInput.value = data.Message || "[ERROR GETTING MESSAGE]"; // Set the text box value


	if (data.DownloadUrl) {
		downloadButton.href = data.DownloadUrl; // Set download link
		downloadButton.innerText = "[Download File]";
		downloadButton.setAttribute("download", "");
	}

	downloadButton.href = data.DownloadUrl; // Update download link
}
