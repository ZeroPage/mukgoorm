window.onload = function() {
    var modal = document.getElementById('share-modal');
    var closeBtn = document.getElementById('close-modal-btn');
    var shareBtns = document.getElementsByClassName('share-btn');
    var copyToCB = document.getElementById('copy-btn');

    copyToCB.onclick = function(event){
      var link = document.querySelector(".link-holder");
      var sucessful = document.execCommand('copy');
    }

    closeBtn.onclick = function(){
        modal.style.display = "none";
    }

    Array.from(shareBtns).forEach(function(element) {
        element.onclick = function(event) {
            link = event.target.parentElement.getElementsByClassName("file-link")[0].href;
            modal.getElementsByClassName("link-holder")[0].textContent = link;
            modal.style.display = "block";
        }
    });
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
