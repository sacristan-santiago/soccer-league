document.addEventListener("DOMContentLoaded", function () {
    // When the DOM content is fully loaded, fetch all players from the database
    fetch("http://localhost:3000/players/all")
        .then(response => response.json())
        .then(players => {
            // Populate the players table with the retrieved players
            const playerTableBody = document.getElementById("playerTableBody");
            players.forEach(player => {
                const newRow = document.createElement("tr");
                newRow.classList.add("PlayersTable");
                newRow.innerHTML = `
                    <td>${player.firstName}</td>
                    <td>${player.lastName}</td>
                    <td>${player.rating}</td>
                    <td><button class="deletePlayerButton" id="${player.id}">Delete</button></td>
                    <td><button class="joinTeamButton" id="${player.id}">Join Team</button></td>
                `;
                playerTableBody.appendChild(newRow);
            });

            // Add event listeners to the delete buttons
            const deletePlayerButtons = document.querySelectorAll(".deletePlayerButton");
            deletePlayerButtons.forEach(button => {
                button.addEventListener("click", deletePlayerButtonClickListener(button));
            });

            const joinTeamButtons = document.querySelectorAll(".joinTeamButton");
            joinTeamButtons.forEach(button => {
                button.addEventListener("click", joinTeamButtonClickListener(button.id));

            })
        })
        .catch(error => {
            console.error("Error fetching players:", error);
        });
});


function deletePlayerButtonClickListener(button) {
    return function () {
        const playerId = button.id;
        // Make a DELETE request to delete the player
        fetch(`http://localhost:3000/players/${playerId}`, {
            method: "DELETE"
        })
            .then(response => {
                if (response.ok) {
                    // If successful, remove the row from the table
                    const rowToRemove = button.closest("tr");
                    rowToRemove.remove();
                } else {
                    console.error("Failed to delete player");
                }
            })
            .catch(error => {
                console.error("Error deleting player:", error);
            });
    };
}

function joinTeamButtonClickListener(joinTeamButtonId) {
    return function () {
        // Here you can display a list of teams that the player can join
        var modal = document.getElementById("myModal");
        modal.style.display = "block";

        // Fetch teams list
        fetch("http://localhost:3000/teams/all", {
            method: "GET"
        })
            .then(response => {
                if (response.ok) {
                    // Parse response JSON asynchronously
                    return response.json();
                } else {
                    throw new Error("Failed to get teams");
                }
            })
            .then(data => {
                // Populate teams list (assuming populateTeamsList function exists)
                populateTeamsList(data, joinTeamButtonId);
            })
            .catch(error => {
                console.error("Error getting teams:", error);
            });
    };
}
