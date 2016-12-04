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
