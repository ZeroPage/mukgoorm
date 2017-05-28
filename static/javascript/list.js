function httpRequest(method, url, body) {
    return new Promise(function(resolve, reject){
        var req = new XMLHttpRequest;
        req.open(method, url);

        req.onload = function() {
            if (req.status == 200) {
                resolve(req.response);
            } else {
                reject(req.response);
            }
        };

        req.onerror = function() {
            reject(Error("Fail to httpRequest " + url));
        };

        req.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        req.send(body);
    });
};

window.onload = function() {
	var shareModal = document.getElementById('share-modal');
	var shareBtns = document.getElementsByClassName('share-btn');
	var closeBtn = document.getElementById('close-btn');
	var copyToCB = document.getElementById('copy-btn');
	var uploadModal = document.getElementById('upload-modal');
	var closeUploadModalBtn = document.getElementById('close-upload-modal-btn');
	var uploadBtn = document.querySelector('.submit-file-btn');
	var uploadSelectedFileBtn = document.querySelector('.upload-btn');
	var remoteDownBtn = document.querySelector(".remote-down-btn");

	closeBtn.onclick = function(){
		shareModal.style.display = "none";
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
			var link = event.target.parentElement.querySelector(".file-link").href;
			shareModal.querySelector(".link-holder-text").textContent = link;
			shareModal.style.display = "block";
		}
	});

	// Search Button
	var searchBtn = document.querySelector('.search-btn');
	searchBtn.onclick = function(event) {
		var query = event.target.parentElement.querySelector(".search-input").value;
		location.href = "/search?q=" + query;
	}

	remoteDownBtn.onclick = function(event) {
		var downModal = document.querySelector("#remote-download-modal");
		downModal.style.display = "block";
	}

	var remoteDownloadCancelBtn = document.querySelector(".remote-down-cancel-btn");
	remoteDownloadCancelBtn.onclick = function(event) {
		var modal = document.querySelector("#remote-download-modal");
		modal.style.display = "none";
		return false;
	}

	var remoteDownSubmitBtn = document.querySelector(".remote-down-submit-btn");
	remoteDownSubmitBtn.onclick = function(event) {
		var input = event.target.parentElement.querySelector("input");
		if (input.value.length === 0) {
			alert("Please enter url.");
			return false;
		}
	}

	var deleteBtns = document.getElementsByClassName('delete-btn');
	Array.from(deleteBtns).forEach(function(elem){
		elem.onclick = function(event) {
			var path = event.target.parentElement.getAttribute("path");

			httpRequest('DELETE', "/delete?dir=" + path).then(
				response => {
					location.reload();
				},
				error => {
					alert(JSON.parse(error).error);
				});
		}
	})
}
