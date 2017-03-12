window.onload = function() {
    var modal = document.getElementById('share-modal');
    var closeBtn = document.getElementById('close-modal-btn');
    var shareBtns = document.getElementsByClassName('share-btn');
    var copyToCB = document.getElementById('copy-btn');

    var uploadModal = document.getElementById('upload-modal');
    var closeUploadModalBtn = document.getElementById('close-upload-modal-btn');
    var uploadBtn = document.getElementsByClassName('submit-file')[0];

    closeBtn.onclick = function(){
        modal.style.display = "none";
    }

    closeUploadModalBtn.onclick = function(event) {
        uploadModal.style.display = "none";
    }

    // TODO copy button
    copyToCB.onclick = function(event) {
        var link = document.querySelector(".link-holder-text");
        link.select();
        var successful = document.execCommand('copy');
    }

    uploadBtn.onclick = function(event) {
        uploadModal.style.display = "block";
    }

    Array.from(shareBtns).forEach(function(element) {
        element.onclick = function(event) {
            link = event.target.parentElement.getElementsByClassName("file-link")[0].href;
            modal.getElementsByClassName("link-holder-text")[0].textContent = link;
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
}

function remove(filePath) {
    // TODO Should we consider IE?
    var httpRequest = new XMLHttpRequest;

    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState == 4) {
            if (httpRequest.status == 200) {
                location.reload()
            } else {
                alert("Can't delete file/folder")
            }
        }
    }

    url = "/delete?dir=" + filePath;
    httpRequest.open('DELETE', url);
    httpRequest.send(null);
}
