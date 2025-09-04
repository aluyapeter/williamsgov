# Task CLI Manager called Williamsgov

Williamsgov is a simple command-line task manager written in Go.  
This tool lets you add, list, complete, and delete tasks directly from your terminal.

---

## ✨ Features
- Add new tasks
- List all tasks
- Mark tasks as completed
- Delete tasks
- Persistent storage (tasks are saved to a JSON file)

---

## 🚀 Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/aluyapeter/williamsgov.git
   cd williamsgov
2. Build the binary:
   ```bash
   go buil -o williamsgov

---

📖 Usage
After building, you can run commands like this:
## Add a task:
```bash
./willlaimsgov add "Task title" -d "Task description"
```

## Get the list of all non-completed tasks, organized by their ID:
```bash
./williamsgov list
```
Use [williamsgov list --all] to list both completed and non-completed tasks

## Mark a task as completed (using the id of the task):
```bash
./williamsgov done 1
```

## Delete a task by ID:
```bash
./williamsgov delete 1
```

---

# 🛠️ Development

To run the project in development without building a binary:
```bash
go run main.go add "Task title" --description "Task description"
```

---
# ✅ Running Tests (which I strongly believe might only be attempted by senior devs)

If you wrote tests, run them with:
```bash
go test ./...
```

# 🤝 Contributing

Pull requests are welcome! For major changes, open an issue first to discuss what you’d like to change.
