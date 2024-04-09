document.addEventListener("DOMContentLoaded", function () {
    // When the DOM content is fully loaded, fetch all players from the database
    fetch("http://localhost:3000/players/all")
        .then(response => response.json())
        .then(players => {
            // Populate the players table with the retrieved players
            const playerTableBody = document.getElementById("playerTableBody");
            players.forEach(player => {
                const teamsList = document.createElement("ul");
                teamsList.dataset.playerId = player.id

                fetch(`http://localhost:3000/player/teams/${player.id}`)
                    .then(response => response.json())
                    .then(teams => {
                        if (teams != null) {
                            teams.forEach(team => {
                                const teamRow = document.createElement("li")
                                teamRow.innerHTML = team.name
                                teamsList.appendChild(teamRow)
                            });
                        }

                        // Create table row and append to table after playerTeams are fetched
                        const newRow = document.createElement("tr");
                        newRow.classList.add("PlayersTable");
                        newRow.innerHTML = `
                            <td>${player.firstName}</td>
                            <td>${player.lastName}</td>
                            <td>${player.rating}</td>
                            <td class="teams-column"></td>
                            <td><button class="deletePlayerButton" id="${player.id}">Delete</button></td>
                            <td><button class="manageTeamsButton" id="${player.id}">Manage Teams</button></td>
                        `;
                        const teamsColumn = newRow.querySelector(".teams-column");
                        teamsColumn.appendChild(teamsList);

                        playerTableBody.appendChild(newRow);

                        // Add event listeners to the delete and join team buttons inside this block
                        const deletePlayerButton = newRow.querySelector(".deletePlayerButton");
                        const manageTeamsButton = newRow.querySelector(".manageTeamsButton");

                        deletePlayerButton.addEventListener("click", deletePlayerButtonClickListener(deletePlayerButton));
                        manageTeamsButton.addEventListener("click", manageTeamsButtonClickListener(manageTeamsButton.id));
                    })
                    .catch(err => {
                        console.log("Error fetching player teams: ", err);
                    });
            });
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

function manageTeamsButtonClickListener(manageTeamsButtonId) {
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
                populateTeamsList(data, manageTeamsButtonId);
            })
            .catch(error => {
                console.error("Error getting teams:", error);
            });
    };
}

function populateTeamsList(teams, playerId) {
    // Clear previous list items
    teamList.innerHTML = "";

    fetch(`http://localhost:3000/player/teams/${playerId}`)
        .then(response => {
            return response.json();
        })
        .then(playerTeams => {
            teams.forEach(function (team) {
                var listItem = document.createElement("li");
                listItem.textContent = team.name;

                var button;
                if (playerTeams && playerTeams.some(playerTeam => playerTeam.id === team.id)) {
                    // Team is joined, create a "Leave" button
                    button = createButton("Leave", "leaveButton", team.id, playerId, true);
                } else {
                    // Team is not joined, create a "Join" button
                    button = createButton("Join", "joinButton", team.id, playerId, false);
                }

                // Append the button to the list item
                listItem.appendChild(button);

                // Append the list item to the team list
                teamList.appendChild(listItem);
            });
        })
        .catch(err => {
            console.error("Error getting player teams:", err);
        });
}

// Function to create a button
function createButton(text, className, teamId, playerId, isLeaveButton) {
    var button = document.createElement("button");
    button.textContent = text;
    button.classList.add(className);
    button.dataset.teamId = teamId; // Store the team ID as a data attribute

    // Add event listener to the button
    button.addEventListener("click", function () {
        var method = isLeaveButton ? "DELETE" : "POST";
        var apiUrl = `http://localhost:3000/teams/player/${teamId}/${playerId}`;

        // Make API call to join or leave the team
        fetch(apiUrl, {
            method: method
        })
            .then(response => {
                if (response.ok) {
                    if (isLeaveButton) {
                        console.log("Player left the team successfully");
                    } else {
                        console.log("Player joined the team successfully");
                    }
                    //UPDATE THE teamList element corresponding to the playerId with the playerTeams
                    const teamListToUpdate = document.querySelector(`ul[data-player-id="${playerId}"]`);
                    teamListToUpdate.innerHTML = ""
                    fetch(`http://localhost:3000/player/teams/${playerId}`)
                        .then(response => response.json())
                        .then(teams => {
                            if (teams != null) {
                                teams.forEach(team => {
                                    const teamRow = document.createElement("li")
                                    teamRow.innerHTML = team.name
                                    teamListToUpdate.appendChild(teamRow)
                                });
                            }
                        })
                        .catch(error => {
                            console.error("Error performing action:", error);
                        });
                    closeModal()

                } else {
                    console.error("Failed to perform action");
                }
            })
            .catch(error => {
                console.error("Error performing action:", error);
            });
    });

    return button;
}