# YOUR PROJECT TITLE
#### Video Demo: https://youtu.be/V791HxCerjo
#### Description:

# Soccer League Manager

This Soccer League project provides functionality for managing a soccer league. It allows you to handle new players, teams, games, load results, follow team positions, tournaments, and more.

## Features

- **Player Management**: Add, update, and delete player information.
- **Team Management**: Create, edit, and remove teams from the league.
- **Game Management**: Schedule and manage games between teams.
- **Results Loading**: Record game results and update team standings accordingly.


Here's an improved version:

## Project Files Overview

1. **main.go**: This file launches the server, serving both a backend REST API and a simple front-end with HTML and Javascript files.

2. **handlers.go**: Each endpoint from `main.go` is routed to the appropriate functions in the data access layer files, where data input/output is processed.

3. **matchDAL.go, teamsDAL.go, playerDAL.go**: These files manage interactions with the SQLite database, handling data operations for matches, teams, and players respectively.

4. **sqliteInit.go**: Responsible for setting up the database by creating the necessary tables required for the project.

5. **structs.go**: Contains all the necessary structs for the project, including definitions for Player, Match, Team, and data transfer objects.

6. **errors.go**: Houses a simple error handling function for Go.

7. **go.mod, go.sum, soccer-league.exe**: Files generated during the building/running of the project.

8. **teams.go**: Provides additional functionalities related to team creation.

9. **/static**: This directory contains all the HTML and Javascript files necessary for running the UI.

10. **/static/index.html**: Basic structure for the user interface.

11. **/static/loadGamesTable.js, /static/loadPlayersTable.js, /static/loadTeamsTable.js**: These files populate tables with data by making calls to the backend API. They also implement dynamic buttons and functionalities.

12. **/static/modalHandler.js, /static/playersHandler.js, /static/gamesHandler.js, /static/teamsHandler.js**: Responsible for managing dynamic interactions with the UI, updating the HTML as needed.

13. **/static/styles.css**: A simple stylesheet for designing the UI.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/sacristan-santiago/soccer-league.git
    ```

2. Navigate to the project directory:

    ```bash
    cd soccer-league
    ```

3. Install dependencies:

    ```bash
    go install github.com/mattn/go-sqlite3
    ```

4. Start the application:

    ```bash
    go build .
    ./soccer-league.exe
    ```

5. Open your web browser and visit `http://localhost:3000` to access the Soccer League ERP.

## Usage

1. **Player Management**:
   - Navigate to the Players section to add, update, or delete player information.

2. **Team Management**:
   - Go to the Teams section to create, edit, or remove teams from the league.

3. **Game Management**:
   - Schedule games between teams in the Games section.

4. **Results Loading**:
   - Record game results and update team standings accordingly.

## Contributing

Contributions are welcome! Please feel free to submit bug reports, feature requests, or pull requests.

## License

This project is licensed under the [MIT License](LICENSE).

## Support

For any questions or issues, please contact [sacristan-santiago@gmail.com](mailto:sacristan-santiago@gmail.com).
