const maxDepth = 3;
let socket;

document.getElementById("archiveForm").addEventListener("submit", function(event) {
	event.preventDefault(); // Prevent default form submission

	const options = {
		NoAudio: document.getElementById("NoAudio").checked,
		NoCss: document.getElementById("NoCss").checked,
		// NoIframe: document.getElementById("NoIframe").checked,
		NoFonts: document.getElementById("NoFonts").checked,
		NoJs: document.getElementById("NoJs").checked,
		NoImages: document.getElementById("NoImages").checked,
		NoVideo: document.getElementById("NoVideo").checked,
		// NoMetadata: document.getElementById("NoMetadata").checked,
		Targets: document.getElementById("Targets").value.trim(),
		AllowDomains: document.getElementById("AllowDomains").value.trim(),
		BlockDomains: document.getElementById("BlockDomains").value.trim(),
		Output: document.getElementById("Output").value.trim(),
		Cookie: document.getElementById("Cookie").value.trim(),
		Depth: parseDepth(document.getElementById("Depth").value),
		// Mode: document.getElementById("Mode").value.trim(),
		Timeout: parseTimeOut(document.getElementById("Timeout").value)
	};

	// Web Socket Shenanigans
	//------------------------------------------------------------------
	if (socket) {
		socket.close();  // Close any existing connection
	}
	socket = new WebSocket("ws://localhost:8080/messaging");

	socket.onmessage = function(event) {
		message_box = document.getElementById("message");
		message_box.value = message_box.value + event.data;
	};

	socket.onopen = function() {
		console.log("Connected to WebSocket server.");
	};

	socket.onerror = function(error) {
		console.error("WebSocket Error:", error);
	};

	socket.onclose = function() {
		console.log("WebSocket connection closed.");
	};
	//------------------------------------------------------------------

	sendData(options);
});

function parseDepth(value) {
	let num = parseInt(value, 10);
	if (isNaN(num) || num < 0 || num > 255) {
		return 0;
	} else {
		if (num > maxDepth) {
			alert(`We suggest maximum depth of ${maxDepth} if your intention is get tangible epub document.`);
			return 0;
		} else {
			return num;
		}
	};
}

function parseTimeOut(value) {
	let num = parseInt(value, 10);
	if (isNaN(num) || num < 0) {
		return 0
	} else {
		return num
	}
}

function sendData(data) {
	console.log("Sending data:", JSON.stringify(data, null, 2));

	setTimeout(() => {
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
	}, 500);  // Slight delay to allow WebSocket connection
}

function displayResponse(data) {
	const responseInput = document.getElementById("message");
	const downloadButton = document.getElementById("downloadButton");

	responseInput.value = responseInput.value + (data.Message || "[ERROR GETTING MESSAGE]");


	if (data.DownloadUrl) {
		downloadButton.href = data.DownloadUrl; // Set download link
		downloadButton.innerText = "[Download File]";
		downloadButton.setAttribute("download", "");
	}
}
