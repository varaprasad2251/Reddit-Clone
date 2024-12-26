# COP5615 Project 4
Project 4 - Reddit Clone in GO

## Description

This project implements a Reddit-like engine with core functionalities and a client simulator. The engine supports user registration, subreddit management, posting, commenting, voting, and direct messaging. The simulator tests the engine by simulating multiple users performing various actions concurrently.

## Group Members

* Member 1: Chenna Kesava Varaprasad Korlapati (UFID: 4836-8778)
* Member 2: Phalguna Peravali (UFID: 3753-9361)

## Setup and Prerequisites

### Prerequisites
For this project, you need to have GO installed on your machine. Here are the steps to install GO on macOS:

### Installing GO on macOS

1. Visit the official GO downloads page: https://golang.org/dl/

2. Download the macOS installer package for the latest GO version.

3. Open the downloaded package and follow the prompts to install GO.

4. After installation, open a terminal and verify the installation by running:
    ~~~
    go version
    ~~~
    This should display the installed Go version.

5. Set up Go workspace by adding the following to ~/.bash_profile or ~/.zshrc file:
    ~~~
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    ~~~
6. Reload shell configuration:
    ~~~
    source ~/.bash_profile
    ~~~
    or
    ~~~
    source ~/.zshrc
    ~~~

Now you have Go installed and configured on your macOS system.

### Project Setup
1. **Clone the repository**:
    ```
    git clone https://github.com/varaprasad2251/COP5615_Project4.git
    ```

2. Change the directory to Project root directory
    ```
    cd COP5615_Project4
    ```

3. Install project dependencies:
   ~~~
   go mod tidy
   ~~~

4. Execute the program
    ```
    go run main.go -users <num_users>
    ```
    Here `<num_users>` is the maximum number of users to be simulated

    For `<num_users> > 50` , it is safe to use the below command that writes the output to `output.txt` file since the output will be large number of lines.
    ```
    go run main.go -users <num_users> >output.txt
    ```

# 4.1

## Overview

### Engine (Engine/engine.go)

The Reddit engine is implemented using the Actor model with the following key features:

- User registration and management
- Subreddit creation, joining, and leaving
- Posting and commenting in subreddits
- Upvoting and downvoting posts
- Karma computation
- Direct messaging between users

### Simulation (simulation/simulation.go)

The simulation creates multiple concurrent users (default: 10) who perform random actions on the Reddit clone. 

Each user can perform the below actions:

- Register an account
- Join subreddits
- Create posts
- Comment on posts
- Upvote and downvote posts
- Send and reply to direct messages

The simulation runs for a fixed number of actions per user (default: 20) with random delays between actions to simulate real-world usage patterns.

## Modules

### 1. Engine (Engine/engine.go)

The core engine that handles all Reddit-like functionalities:

- `RegisterUser`: Creates a new user account
- `SubredditSpecificOp`: Handles joining and leaving subreddits
- `CreatePost`: Creates a new post in a subreddit
- `ReplyToComment`: Adds a comment to a post or another comment
- `SendDMtoUser`: Sends a direct message to another user
- `ReplyToAllDMs`: Replies to all received direct messages
- `UpvoteRandomPost` and `DownvoteRandomPost`: Simulates voting on posts

### 2. Simulation (simulation.go)

Simulates user activities:

- `SimulateUser`: Performs a series of random actions for each simulated user
- Utilizes a variety of message types (e.g., `RegisterUser`, `UserJoinSubReddit`, `CreatePost`) to interact with the engine

### 3. Main (main.go)

Initializes the ActorSystem, creates the Engine, and spawns multiple simulated users.

## Stats

The engine tracks various statistics:

- Total number of users
- Total number of posts
- Total number of subreddits
- Total number of messages
- Total Simulation Time 

These stats are collected throughout the simulation and printed at the end using the `PrintStats` method of the Engine.


## Running the Simulation

To run the simulation:

1. Ensure all dependencies are installed
2. Run `go run main.go -users <num_users>`

The simulation will output detailed logs of user actions and their results. At the end of the simulation, comprehensive statistics will be displayed.

## Simulation Statistics

| Total Users | Total Posts | Total Subreddits | Total Messages | Total Simulation Time |
|-------------|-------------|------------------|----------------|-----------------------|
| 10          | 131         | 106              | 113            | 40.5556424s           |
| 50          | 603         | 541              | 596            | 42.121259s            |
| 100         | 1186        | 1111             | 1141           | 43.4298415s           |
| 500         | 5407        | 5448             | 5323           | 44.0380639s           |
| 1000        | 10423       | 10779            | 10683          | 44.5763764s           |
| 5000        | 56089       | 56706            | 54101          | 45.0230307s           |

## Future Improvements (In Part II)

- Implement a REST API for web client integration
- Implement a more sophisticated karma system
- Enhance the simulation to better mimic real-world usage patterns, including Zipf distribution for subreddit membership


# 4.2

# REST API Interface 

In this part, we extended the Reddit engine developed in 4.1 by implementing a REST API interface and a client to interact with it. The system demonstrates concurrent user interactions and core Reddit like functionalities.

## Features

1. REST API interface for the Reddit engine
2. Simple client implementation to interact with the API
3. Multi-client simulation to demonstrate concurrent functionality

## Packages Used

- `github.com/gin-gonic/gin`: Web framework used to create the REST API
- `net/http`: Standard Go package for HTTP client and Server implementation
- `encoding/json`: Used for JSON encoding and decoding
- `sync`: Provides synchronization primitives, used for coordinating goroutines
- `time`: Used for adding delays between client actions

## Components

### 1. API (api/api.go)

The API is implemented using the Gin web framework and provides endpoints for core Reddit-like functionalities:

- Register an account
- Join, Leave subreddits
- Create posts
- Comment on posts
- Upvote and downvote posts
- Send and reply to direct messages

List of  Endpoints:

| Endpoint                   | Method | Description                        |
|----------------------------|--------|------------------------------------|
| `/api/register`            | POST   | Register a new user               |
| `/api/user/:username`      | GET    | Get user information              |
| `/api/user/:username/join` | POST   | Join a subreddit                  |
| `/api/user/:username/leave`| POST   | Leave a subreddit                 |
| `/api/subreddit`           | POST   | Create a new subreddit            |
| `/api/subreddit/:name`     | GET    | Get subreddit information         |
| `/api/submit`              | POST   | Create a new post                 |
| `/api/posts/:id/upvote`    | POST   | Upvote a post                     |
| `/api/posts/:id/downvote`  | POST   | Downvote a post                   |
| `/api/comment`             | POST   | Create a comment                  |
| `/api/message/compose`     | POST   | Send a direct message             |
| `/api/message/inbox`       | GET    | Get user's direct messages        |
| `/api/feed`                | GET    | Get user's feed                   |

---

Here are some sample API requests that can be made while the server is running:

1. To Register a User:
    ~~~
    curl -X POST http://localhost:8080/api/register -H "Content-Type: application/json" -d '{"username":"testuser"}'
    ~~~
2. To Create a subreddit:
    ~~~
    curl -X POST http://localhost:8080/api/subreddit -H "Content-Type: application/json" -d '{"name":"testsubreddit"}'
    ~~~
3. To Join a subreddit:
    ~~~
    curl -X POST http://localhost:8080/api/user/testuser/join -H "Content-Type: application/json" -d '{"name":"testsubreddit"}'
    ~~~
4. To Create a Post:
    ~~~
    curl -X POST http://localhost:8080/api/submit -H "Content-Type: application/json" -d '{"username":"testuser", "subreddit":"testsubreddit", "content":"This is a test post"}'
    ~~~

## 2. Client Implementation (client.go)

The client provides a simple interface to interact with the API. It supports all core functionalities, allowing users to:

- Register users
- Create subreddits
- Join/leave subreddits
- Create posts
- Upvote/downvote posts
- Comment on posts
- Send direct messages
- Retrieve user information

## 3. Multi-Client Simulation (multi_client.go)

The multi-client simulation demonstrates concurrent interactions with the API. It simulates multiple clients performing actions such as registering users, creating subreddits, posting content, and interacting with each other's data.

### Key Features:
- Creates multiple RedditClient instances.
- Simulates various user actions concurrently.
- Demonstrates the system's ability to handle multiple clients simultaneously.


## Usage

### Start the Server

Run the Reddit-like engine server:
~~~
go run main.go
~~~

### To Run Single-Client
To create a single client and call the api endpoints to test their functionality. (in another terminal window)
~~~
go run client.go
~~~

### To Run Multi-Client Simulation
Run the multi-client simulation to test concurrent requests to the api server. For this, we have created two clients and made different endpoint calls from the two clients. (in another terminal window)
~~~
go run multi_client.go
~~~
 

### Demo Link

Here's the link to the demo video for 4.2 : [link](https://drive.google.com/file/d/1C_gcwL3DYXVF1fidh32jhFCbvlZ6dqe_/view?usp=sharing)


## Conclusion
This project showcases a Reddit-like system with a REST API interface and demonstrates its ability to handle concurrent client interactions. It supports core functionalities such as user registration, subreddit management, posting, voting, commenting, and messaging. The multi-client simulation validates that the system can handle multiple clients interacting with server by making requests concurrently while maintaining data integrity.