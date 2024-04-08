document.getElementById("teamForm").addEventListener("submit", function (event) {
    event.preventDefault(); // Prevent the default form submission

    // Get form data
    const formData = {
        Name: document.getElementById("Name").value,
    };

    // Convert form data to JSON
    const jsonData = JSON.stringify(formData);

    // Make a POST request with JSON data
    fetch("http://localhost:3000/teams", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: jsonData
    })
        .then(response => response.json())
        .then(data => {
            console.log("Team added:", data);

            // Add the new team to the table
            const teamTableBody = document.getElementById("teamTableBody");
            const newRow = document.createElement("tr");
            newRow.classList.add("TeamsTable")
            newRow.innerHTML = `
                <td>${formData.Name}</td>
                <td><button class="deleteTeamButton" id="${data.id}">Delete</button></td>
            `;
            teamTableBody.appendChild(newRow);

            // Clear form fields after adding the team
            document.getElementById("teamForm").reset();

            // Add event listener to the delete button
            const deleteTeamButton = newRow.querySelector(".deleteTeamButton");
            deleteTeamButton.addEventListener("click", deleteTeamButtonClickListener(deleteTeamButton));
        })
        .catch(error => {
            console.error("Error adding team:", error);
        });
});