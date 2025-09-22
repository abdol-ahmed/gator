# Welcome to the Blog Aggregator -gator

`gator` CLI tool allows users to add and manage RSS feeds, storing the collected posts in a PostgreSQL database. Users can follow or unfollow feeds and view summaries of aggregated posts—complete with links—directly in their terminal. The project targets users familiar with Go and SQL, and its key learning goals are:
- Integrating a Go application with PostgreSQL
- Practicing SQL skills (using sqlc and goose for typesafe migrations)
- Writing a long-running Go service for continuous RSS feed collection and storage
It’s designed to help you track your favorite blogs, news, and podcasts efficiently with aggregated updates.

## ✅ Prerequisites:

To run this program, two prerequisites must be met: **PostgreSQL** and **Go** must be installed on your system.
- **PostgreSQL**: This program interacts with a database, and PostgreSQL is the specific database system required. Ensure that PostgreSQL is installed and configured, and that you have the necessary credentials to access the database.
- **Go**: The program itself is written in Go. Therefore, the Go programming language environment needs to be installed on your system to compile and execute the program. You can find installation instructions for Go on the official Go website.


## Installation:
To install the gator CLI tool from the specified GitHub repository using go install, follow these steps:
### 💾 Step 1: Install gator
First, ensure you have the Go programming language installed and that your $PATH is configured correctly. 
1. **Open** your terminal or command prompt.

2. **Run** the following `go install` command:
    ```bash
        go install github.com/abdol-ahmed/gator@latest
    ```
   The `go` tool will download the source code, compile the program, and place the executable binary file in your Go bin directory (e.g., `$HOME/go/bin`). 

3. Verify the installation by checking its version:
    ```bash
      gator -v
    ```

### ⚙️ Step 2: Set up the configuration file
The gator tool requires a database to store user and feed data. You will need a PostgreSQL database instance running locally and a configuration file to tell gator how to connect to it. 
Create a configuration file named `.gatorconfig.json` in your home directory (e.g., `~/.gatorconfig.json`). 
Add the database connection details to the file. A basic example looks like this:
    json

    {
      "db_url": "postgresql://[USERNAME]:[PASSWORD]@localhost:5432/[DBNAME]?sslmode=disable",
      "current_user_name": "YOUR_USER_NAME"
    }

ℹ️ _Note_: Replace the bracketed values ([USERNAME], [PASSWORD], and [DBNAME]) with your PostgreSQL credentials. 

### 🤖 Step 3: Run gator commands
Once the installation and configuration are complete, you can start using gator from your terminal. Since you need to be a registered user to do anything, you'll start there. 
1. **Register a user**
    
    Create a new user account with a name of your choice.
    ```bash
        gator register my_username
    ```

2. **Log in**

    Log in as the user you just created.
    ```bash
        gator login my_username
    ```
3. **Add an RSS feed**

    Add a new RSS feed to the database. You will be able to follow this feed later.
    ```bash
        gator addfeed "https://www.example.com/rss.xml"
    ```

4. **Follow a feed**

    Have your user account follow a feed that has been added to the database. 
    ```bash
        gator follow "https://www.example.com/rss.xml"
    ```
5. **Aggregate feeds**

    Fetch new posts from all the feeds your user is following. Duration can be 1s, 1m, 1h, etc. This example tells gator to check for new posts every 30 seconds.
    ```bash
        gator agg 30s
    ```

6. **Browse posts**

    Browse the posts saved for the current user. This command will display the 10 most recent posts. The default display will be 2 recent posts.
    ```bash
        gator browse 10
    ```

ℹ️ There are a few other commands you may need as well:
- **reset:** `gator reset` - Delete all users, feeds, and posts 
- **users:** `gator users` - List all users
- **feeds:** `gator feeds` - List all feeds
- **unfollow:** `gator unfollow <url>` - Unfollow a feed that already exists in the database