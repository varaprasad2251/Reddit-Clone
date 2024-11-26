# COP5615_Project4
Reddit Clone and a client tester/simulator


Download Go
Visit the official Go website: https://go.dev/dl/.
Download the macOS ARM64 installer (.pkg file) for M1/M2 chips.
3. Install Go
Open the downloaded .pkg file.
Follow the on-screen instructions to complete the installation.
By default, Go will be installed in /usr/local/go.
4. Add Go to Your PATH
To ensure Go is accessible from the terminal:

Open your terminal and edit your shell configuration file. Depending on your shell:

For zsh (default on macOS):
bash
Copy code
nano ~/.zshrc
For bash:
bash
Copy code
nano ~/.bash_profile
Add the following line at the end of the file:

bash
Copy code
export PATH=$PATH:/usr/local/go/bin
Save and close the file (Ctrl+O, then Enter, then Ctrl+X in nano).

Reload the configuration file:

bash
Copy code
source ~/.zshrc  # or ~/.bash_profile




Build and run the program:
bash
Copy code
go run main.go



# Reddit Clone

## Description

This project implements a Reddit-like engine with core functionalities and a client simulator. The engine supports user registration, subreddit management, posting, commenting, voting, and direct messaging. The simulator tests the engine by simulating multiple users performing various actions concurrently.

## Overview

### Engine (redditEngine/engine.go)

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

### 1. RedditEngine (redditEngine/engine.go)

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

Initializes the ActorSystem, creates the RedditEngine, and spawns multiple simulated users.

## Stats

The engine tracks various statistics:

- Total number of users
- Total number of posts
- Total number of subreddits
- Total number of messages
- User-specific stats (karma, post count)
- Simulation duration

These stats are collected throughout the simulation and printed at the end using the `PrintStats` method of the RedditEngine.

## Running the Simulation

To run the simulation:

1. Ensure all dependencies are installed
2. Run `go run main.go`

The simulation will output detailed logs of user actions and their results. At the end of the simulation, comprehensive statistics will be displayed.

## Future Improvements

- Implement a REST API for web client integration
- Add support for images and markdown in posts
- Implement a more sophisticated karma system
- Enhance the simulation to better mimic real-world usage patterns, including Zipf distribution for subreddit membership