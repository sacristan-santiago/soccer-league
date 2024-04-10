document.getElementById("gamesForm").addEventListener("submit", function (event) {
    event.preventDefault(); // Prevent the default form submission

    const teamASelect = document.getElementById("TeamA");
    const teamBSelect = document.getElementById("TeamB");
    const scoreAInput = document.getElementById("ScoreA");
    const scoreBInput = document.getElementById("ScoreB");

    // Check if both TeamA and TeamB have been selected
    if (teamASelect.value && teamBSelect.value) {
        // Both teams have been selected, make the API call
        const formData = {
            TeamA: parseInt(teamASelect.value),
            TeamB: parseInt(teamBSelect.value),
            ScoreA: scoreAInput.value ? parseInt(scoreAInput.value) : null,
            ScoreB: scoreBInput.value ? parseInt(scoreBInput.value) : null
        };

        console.log(formData)

        // Make the API call to /match endpoint
        fetch("http://localhost:3000/match", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(formData)
        })
            .then(response => response.json())
            .then(game => {
                console.log("Game added:", game);

                // Add the new player to the table
                const gamesTableBody = document.getElementById("gamesTableBody");
                const newRow = document.createElement("tr");
                newRow.classList.add("GamesTable")
                newRow.innerHTML = `
                <td>${teamASelect.selectedOptions[0].textContent}</td>
                <td id=cellScoreA${game.id}>${formData.ScoreA ? formData.ScoreA : `<input class='width100' type='number' id='inputScoreA${game.id}' name='ScoreA'>`}</td>
                <td id=cellScoreB${game.id}>${formData.ScoreB ? formData.ScoreB : `<input class='width100' type='number' id='inputScoreB${game.id}' name='ScoreB'>`}</td>
                <td>${teamASelect.selectedOptions[0].textContent}</td>
                <td><button class="deleteGameButton" id="${game.id}">Delete</button></td>
                <td><button class="updateGameButton" id="${game.id}">Update Score</button></td>
            `;


                gamesTableBody.appendChild(newRow);

                // Clear form fields after adding the player
                document.getElementById("gamesForm").reset();

                // Add event listener to the delete button
                const deleteGamesButton = newRow.querySelector(".deleteGameButton");
                deleteGamesButton.addEventListener("click", deleteGameButtonClickListener(deleteGamesButton));

                //Add event listener to update button
                const updateGameButton = newRow.querySelector(".updateGameButton");
                updateGameButton.addEventListener("click", updateGameButtonClickListener(updateGameButton));
            })
            .catch(error => {
                console.error("Error adding match:", error);
            });
    }
});