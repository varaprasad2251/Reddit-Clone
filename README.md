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