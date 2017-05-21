window.onload = function() {
	var modal = document.getElementById('share-modal');
	var shareBtns = document.getElementsByClassName('share-btn');
	var closeBtn = document.getElementById('close-btn');
	var copyToCB = document.getElementById('copy-btn');
	var uploadModal = document.getElementById('upload-modal');
	var closeUploadModalBtn = document.getElementById('close-upload-modal-btn');
	var uploadBtn = document.querySelector('.submit-file-btn');
	var uploadSelectedFileBtn = document.querySelector('.upload-btn');
	var remoteDownBtn = document.querySelector(".remote-down-btn");

	closeBtn.onclick = function(){
		modal.style.display = "none";
	}

	closeUploadModalBtn.onclick = function(event) {
		uploadModal.style.display = "none";
		return false;
	}

	copyToCB.onclick = function(event) {
		var link = document.querySelector(".link-holder-text");
		link.select();
		var successful = document.execCommand('copy');
	}

	uploadBtn.onclick = function(event) {
		uploadModal.style.display = "block";
	}

	uploadSelectedFileBtn.onclick = function(event) {
		var isSelected = document.getElementById('uploaded-file').files.length > 0;
		if(!isSelected){
		  alert("You need to select a file.");
		}
		return isSelected;
	}

	Array.from(shareBtns).forEach(function(element) {
		element.onclick = function(event) {
			link = event.target.parentElement.querySelector(".file-link").href;
			modal.querySelector(".link-holder-text").textContent = link;
			modal.style.display = "block";
		}
	});

	// Search Button
	var searchBtn = document.querySelector('.search-btn');
	searchBtn.onclick = function(event) {
		var query = event.target.parentElement.querySelector(".search-input").value;
		url = "/search?q=" + query;
		location.href = url;
	}

	remoteDownBtn.onclick = function(event) {
		downModal = document.querySelector("#remote-download-modal");
		downModal.style.display = "block";
	}

	var remoteDownloadCancelBtn = document.querySelector(".remote-down-cancel-btn");
	remoteDownloadCancelBtn.onclick = function(event) {
		modal = document.querySelector("#remote-download-modal");
		modal.style.display = "none";
		return false;
	}

	var remoteDownSubmitBtn = document.querySelector(".remote-down-submit-btn");
	remoteDownSubmitBtn.onclick = function(event) {
		input = event.target.parentElement.querySelector("input");
		if (input.value.length === 0) {
			alert("Please enter url.");
			return false;
		}
	}
}

function remove(filePath) {
	// TODO Should we consider IE?
	var httpRequest = new XMLHttpRequest;

	httpRequest.onreadystatechange = function() {
		if (httpRequest.readyState == 4) {
			if (httpRequest.status == 200) {
				location.reload();
			} else {
				alert("Can't delete file/folder");
			}
		}
	}

	url = "/delete?dir=" + filePath;
	httpRequest.open('DELETE', url);
	httpRequest.send(null);
}
