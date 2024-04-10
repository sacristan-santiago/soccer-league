document.addEventListener("DOMContentLoaded", function () {
    // Fetch teams from the server and populate the options list for TeamA and TeamB inputs
    fetch("http://localhost:3000/teams/all")
        .then(response => response.json())
        .then(teams => {
            let teamsList = teams
            const teamOptions = teams.map(team => `<option value="${team.id}">${team.name}</option>`).join('');
            document.getElementById('TeamA').innerHTML = teamOptions;
            document.getElementById('TeamB').innerHTML = teamOptions;

            //Fetch all games from the database
            fetch("http://localhost:3000/match/all")
                .then(response => response.json())
                .then(games => {
                    // Populate the games table with the retrieved games
                    const gamesTableBody = document.getElementById("gamesTableBody");
                    games.forEach(game => {
                        // Find the team object with the corresponding game id
                        const teamA = teamsList.find(team => team.id === game.teamA);
                        const teamB = teamsList.find(team => team.id === game.teamB);

                        // Check if the teams are found
                        const teamAName = teamA ? teamA.name : 'Team Not Found';
                        const teamBName = teamB ? teamB.name : 'Team Not Found';

                        // Create table row and append to table after playerTeams are fetched
                        const newRow = document.createElement("tr");
                        newRow.classList.add("GamesTable");
                        newRow.innerHTML = `
                            <td>${teamAName}</td>
                            <td id=cellScoreA${game.id}>${game.scoreA ? game.scoreA : `<input class='width100' type='number' id='inputScoreA${game.id}' name='ScoreA'>`}</td>
                            <td id=cellScoreB${game.id}>${game.scoreB ? game.scoreB : `<input class='width100' type='number' id='inputScoreB${game.id}' name='ScoreB'>`}</td>
                            <td>${teamBName}</td>
                            <td><button class="deleteGameButton" id="${game.id}">Delete</button></td>
                            ${game.scoreA && game.scoreB ? "" : `<td><button class="updateGameButton" id="${game.id}">Update Score</button></td>`}
                        `;


                        gamesTableBody.appendChild(newRow);

                        // Add event listeners to the delete and join team buttons inside this block
                        const deleteGameButton = newRow.querySelector(".deleteGameButton");
                        const updateGameButton = newRow.querySelector(".updateGameButton");

                        deleteGameButton.addEventListener("click", deleteGameButtonClickListener(deleteGameButton));
                        game.scoreA && game.scoreB ? null : updateGameButton.addEventListener("click", updateGameButtonClickListener(updateGameButton));
                    });
                })
                .catch(error => {
                    console.error("Error fetching players:", error);
                });
        })
        .catch(error => {
            console.error("Error fetching teams:", error);
        });
});

function deleteGameButtonClickListener(button) {
    return function () {
        const gameId = button.id;
        // Make a DELETE request to delete the game
        fetch(`http://localhost:3000/match/${gameId}`, {
            method: "DELETE"
        })
            .then(response => {
                if (response.ok) {
                    // If successful, remove the row from the table
                    const rowToRemove = button.closest("tr");
                    rowToRemove.remove();
                } else {
                    console.error("Failed to delete game");
                }
            })
            .catch(error => {
                console.error("Error deleting game:", error);
            });
    };
}

function updateGameButtonClickListener(button) {
    return function () {
        const gameId = button.id;
        const row = button.closest("tr");
        const inputScoreA = row.querySelector(`#inputScoreA${gameId}`);
        const inputScoreB = row.querySelector(`#inputScoreB${gameId}`);
        const formData = {
            ScoreA: inputScoreA.value ? parseInt(inputScoreA.value) : console.error("Input field cannot be null to update score"),
            ScoreB: inputScoreB.value ? parseInt(inputScoreB.value) : console.error("Input field cannot be null to update score")
        };

        // Make a PUT request to update the game
        fetch(`http://localhost:3000/match/${gameId}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(formData)
        })
            .then(response => response.json())
            .then(updatedGame => {
                console.log("Game updated:", updatedGame);
                // If successful, update the row from the table

                if (row) {
                    // Update the scoreA and scoreB columns with new values
                    const cellScoreA = row.querySelector(`#cellScoreA${updatedGame.id}`);
                    const cellScoreB = row.querySelector(`#cellScoreB${updatedGame.id}`);

                    if (cellScoreA && cellScoreB) {
                        cellScoreA.innerHTML = updatedGame.scoreA || '';
                        cellScoreB.innerHTML = updatedGame.scoreB || '';

                        //Remove update button after updating table
                        const buttonParent = button.closest("td")
                        button.remove()
                        buttonParent.remove()
                    } else {
                        console.error("Failed to find score cells in table row");
                    }
                }


            })
            .catch(error => {
                console.error("Error updating game:", error);
            });
    };
}