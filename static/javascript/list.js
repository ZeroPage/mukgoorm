window.onload = function() {
    var modal = document.getElementById('share-modal');
    var shareBtn = document.getElementById('share-btn');
    var closeBtn = document.getElementById('close-modal-btn');

    shareBtn.onclick = function(){
        link = event.target.parentElement.getElementsByClassName("file-link")[0].href
        modal.getElementsByClassName("link-holder")[0].textContent = link
        modal.style.display = "block";
    }

    closeBtn.onclick = function(){
        modal.style.display = "none";
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