document.getElementById("playerForm").addEventListener("submit", function (event) {
    event.preventDefault(); // Prevent the default form submission

    // Get form data
    const formData = {
        FirstName: document.getElementById("FirstName").value,
        LastName: document.getElementById("LastName").value,
        Rating: parseInt(document.getElementById("Rating").value)
    };

    // Convert form data to JSON
    const jsonData = JSON.stringify(formData);

    // Make a POST request with JSON data
    fetch("http://localhost:3000/players", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: jsonData
    })
        .then(response => response.json())
        .then(data => {
            console.log("Player added:", data);

            // Add the new player to the table
            const playerTableBody = document.getElementById("playerTableBody");
            const newRow = document.createElement("tr");
            newRow.classList.add("PlayersTable")
            newRow.innerHTML = `
        <td>${formData.FirstName}</td>
        <td>${formData.LastName}</td>
        <td>${formData.Rating}</td>
        <td><button class="deletePlayerButton" id="${data.id}">Delete</button></td>
        <td><button class="joinTeamButton" id="${data.id}">Join Team</button></td>
    `;
            playerTableBody.appendChild(newRow);

            // Clear form fields after adding the player
            document.getElementById("playerForm").reset();

            // Add event listener to the delete button
            const deletePlayerButton = newRow.querySelector(".deletePlayerButton");
            deletePlayerButton.addEventListener("click", deletePlayerButtonClickListener(deletePlayerButton));

            // Add event listener to the Join team button
            const joinTeamButton = newRow.querySelector(".joinTeamButton");
            joinTeamButton.addEventListener("click", joinTeamButtonClickListener(joinTeamButton.id));
        })
        .catch(error => {
            console.error("Error adding player:", error);
        });
});


// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

span.onclick = function () {
    var modal = document.getElementById("myModal");
    modal.style.display = "none";
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
    var modal = document.getElementById("myModal");
    if (event.target == modal) {
        modal.style.display = "none";
    }
}

function populateTeamsList(teams, playerId) {
    var teamList = document.getElementById("teamList");
    // Clear previous list items
    teamList.innerHTML = "";

    teams.forEach(function (team) {
        var listItem = document.createElement("li");
        listItem.textContent = team.name;

        // Create a "Join" button
        var joinButton = document.createElement("button");
        joinButton.textContent = "Join";
        joinButton.classList.add("joinButton");
        joinButton.dataset.teamId = team.id; // Store the team ID as a data attribute

        // Append the "Join" button to the list item
        listItem.appendChild(joinButton);

        // Append the list item to the team list
        teamList.appendChild(listItem);
    });

    // Add event listener to "Join" buttons
    var joinButtons = document.querySelectorAll(".joinButton");
    joinButtons.forEach(function (button) {
        button.addEventListener("click", function () {
            var teamId = button.dataset.teamId;
            // Make API call to add player to the team
            fetch(`http://localhost:3000/teams/player/${teamId}/${playerId}`, {
                method: "POST"
            })
                .then(response => {
                    if (response.ok) {
                        console.log("Player joined the team successfully");
                        // Handle success if needed
                    } else {
                        console.error("Failed to join team");
                        // Handle error if needed
                    }
                })
                .catch(error => {
                    console.error("Error joining team:", error);
                    // Handle error if needed
                });
        });
    });
}