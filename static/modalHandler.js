// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

function closeModal() {
    var modal = document.getElementById("myModal");
    modal.style.display = "none";
}

span.onclick = function () {
    closeModal()
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
    var modal = document.getElementById("myModal");
    if (event.target == modal) {
        modal.style.display = "none";
    }
}