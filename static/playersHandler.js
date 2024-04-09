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
                <td class="teams-column"><ul data-player-id=${data.id}></ul></td>
                <td><button class="deletePlayerButton" id="${data.id}">Delete</button></td>
                <td><button class="manageTeamsButton" id="${data.id}">Manage Teams</button></td>
            `;
            playerTableBody.appendChild(newRow);

            // Clear form fields after adding the player
            document.getElementById("playerForm").reset();

            // Add event listener to the delete button
            const deletePlayerButton = newRow.querySelector(".deletePlayerButton");
            deletePlayerButton.addEventListener("click", deletePlayerButtonClickListener(deletePlayerButton));

            // Add event listener to the Join team button
            const manageTeamsButton = newRow.querySelector(".manageTeamsButton");
            manageTeamsButton.addEventListener("click", manageTeamsButtonClickListener(manageTeamsButton.id));
        })
        .catch(error => {
            console.error("Error adding player:", error);
        });
});