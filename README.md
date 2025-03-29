# Tic-Tac-Toe in Golang

## 📌 Overview
This is a command-line Tic-Tac-Toe game implemented in **Golang**. It follows software design patterns like **Singleton** (for the board) and **Factory** (for player creation).

## 🚀 Features
- Play Tic-Tac-Toe on a **3x3** board.
- **Factory Pattern** for dynamic player creation.
- **Singleton Pattern** ensures a single board instance.
- Supports **Human vs AI** gameplay.

## 🏗️ Project Structure
```
📂 tic-tac-toe-go
│── 📂 factory         # Factory for creating players
│── 📂 game            # Game logic and management
│── 📂 models          # Board and player models
│── main.go           # Entry point
│── go.mod            # Go module file
│── .gitignore        # Ignored files
│── README.md         # Project documentation
```

## 🛠️ Setup & Installation
### **Step 1: Clone the Repository**
```sh
git clone https://github.com/your-username/tic-tac-toe-go.git
cd tic-tac-toe-go
```

### **Step 2: Initialize Go Modules**
```sh
go mod init tic-tac-toe-go
go mod tidy
```

### **Step 3: Run the Game**
```sh
go run main.go
```

## 🎮 How to Play
- The game starts with a **3x3** empty board.
- Players take turns entering row and column numbers (e.g., `0 1`).
- The game ends when a player wins or the board is full.

## 🏆 Design Patterns Used
- **Singleton Pattern** → Ensures only one board instance.
- **Factory Pattern** → Creates players dynamically (Human/AI).

## 📜 License
This project is open-source and free to use.

## 🤝 Contributing
Feel free to contribute! Open issues, suggest features, or create pull requests.

## 📧 Contact
For any queries, reach out at **your-email@example.com**.

