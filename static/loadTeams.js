document.addEventListener("DOMContentLoaded", function () {
    // When the DOM content is fully loaded, fetch all teams from the database
    fetch("http://localhost:3000/teams/all")
        .then(response => response.json())
        .then(teams => {
            // Populate the teams table with the retrieved teams
            const teamTableBody = document.getElementById("teamTableBody");
            teams.forEach(team => {
                const newRow = document.createElement("tr");
                newRow.classList.add("TeamsTable");
                newRow.innerHTML = `
                    <td>${team.name}</td>
                    <td><button class="deleteTeamButton" id="${team.id}">Delete</button></td>
                `;
                teamTableBody.appendChild(newRow);
            });

            // Add event listeners to the delete buttons
            const deleteTeamButton = document.querySelectorAll(".deleteTeamButton");
            deleteTeamButton.forEach(button => {
                button.addEventListener("click", deleteTeamButtonClickListener(button));
            });
        })
        .catch(error => {
            console.error("Error fetching teams:", error);
        });
});


function deleteTeamButtonClickListener(button) {
    return function () {
        const teamId = button.id;
        // Make a DELETE request to delete the team
        fetch(`http://localhost:3000/teams/${teamId}`, {
            method: "DELETE"
        })
            .then(response => {
                if (response.ok) {
                    // If successful, remove the row from the table
                    const rowToRemove = button.closest("tr");
                    rowToRemove.remove();
                } else {
                    console.error("Failed to delete team");
                }
            })
            .catch(error => {
                console.error("Error deleting team:", error);
            });
    };
}