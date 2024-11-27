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
2. Run `go run main.go`

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

## Future Improvements

- Implement a REST API for web client integration
- Add support for images and markdown in posts
- Implement a more sophisticated karma system
- Enhance the simulation to better mimic real-world usage patterns, including Zipf distribution for subreddit membership
