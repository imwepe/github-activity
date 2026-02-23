# GitHub Activity CLI (Golang)

A simple Command Line Interface (CLI) application written in Go to fetch and display recent GitHub user activities.

This project was inspired by the GitHub Activity project idea from roadmap.sh.

## 🚀 Features

- Input GitHub username interactively
- Fetch public activity from GitHub API
- Filter activities by event type:
  - PushEvent
  - PullRequestEvent
  - IssuesEvent
  - WatchEvent (Stars)
  - All events
- Display up to 5 latest events
- Human-readable activity output (e.g., "Pushed 3 commits to repo-name")
- Graceful error handling (invalid username, API failure)
- Interactive menu with Exit option
- Built using only Go standard library (no external dependencies)

---

## 📦 Technologies Used

- Go (Golang)
- net/http
- encoding/json
- time
- fmt
- os
- strings

---

## 🔗 GitHub API Endpoint Used
https://api.github.com/users/<username>/events


---

## 🛠 Installation & Usage
<h2>Make SURE you have GOLANG installed in your device!</h2>

### 1. Clone the repository
<code>git clone https://github.com/imwepe/github-activity.git</code>

<code>cd github-activity</code>


### 2. Run the program
<code>go run main.go</code>

### 3. Follow the interactive prompts

---

## 🧠 How It Works

1. The program asks for a GitHub username.
2. It fetches public events using GitHub's REST API.
3. Events are parsed into Go structs.
4. The user selects a filter from the menu.
5. The program formats the activity into human-readable output.
6. The program continues running until the user selects Exit.

---

## ⚠️ Error Handling

- Username cannot be empty
- Invalid GitHub username returns "user not found"
- API errors are handled gracefully
- Invalid menu input does not crash the program

---

## 🎯 Learning Goals

This project demonstrates:

- Working with REST APIs in Go
- JSON decoding using struct tags
- Using loops and switch statements
- Building interactive CLI applications
- Error handling best practices
- Working with time formatting in Go

---

## 📌 Project Inspiration

This project idea comes from:

https://roadmap.sh/projects/github-user-activity

---
