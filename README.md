# 📝 GoTasker

A simple command-line interface (CLI) tool to track and manage your daily tasks. This tool helps you stay productive by keeping your tasks organized right from your terminal.

> 🚀 Project by [Md. Mehedi Hasan](https://github.com/MehediMubin)

---

## 📌 Features

- ✅ Add new tasks
- ✏️ Update existing tasks
- 🗑️ Delete tasks
- 🚧 Mark tasks as _in progress_
- ✔️ Mark tasks as _done_
- 📄 List all tasks
- 🎯 List tasks by status (`todo`, `in-progress`, `done`)
- 🔄 Reset all tasks to `todo` using the `reset` command
- 🧹 Delete all tasks using the `clear` command
- 📘 View command help with `help` and detailed command help with `help <command>`
- 🧭 List tasks sorted by priority (`high` → `medium` → `low`)
- 🗂️ Stores task data in a local JSON file

---

## 🏁 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/MehediMubin/GoTasker.git
cd GoTasker
```

### 2. Build and Install Globally (for Linux-based systems)

Make sure you have [Go](https://go.dev/dl/) installed on your system.

```bash
go build -o task-cli
sudo mv task-cli /usr/local/bin
```

Now you can run the app globally using just `task-cli`.

---

## 🚀 How to Use

### ➕ Add a New Task

```bash
task-cli add "Buy groceries"
task-cli add "Buy groceries" "high"
```

Supported priorities are `high`, `medium`, and `low`.
If priority is not provided, the task defaults to `low`.

**Output:** `Task added successfully (ID: 1)`

### ✏️ Update an Existing Task

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

### 🗑️ Delete a Task

```bash
task-cli delete 1
```

### 🚧 Mark a Task as In Progress

```bash
task-cli mark-in-progress 2
```

### ✔️ Mark a Task as Done

```bash
task-cli mark-done 2
```

### 📄 List Active Tasks

```bash
task-cli list
```

By default, `task-cli list` shows every task except those marked `done` in a spaced, table-style layout.

### 🎯 List Tasks by Status

```bash
task-cli list todo
task-cli list in-progress
task-cli list done
```

### 🔄 Reset All Tasks to Todo

```bash
task-cli reset
```

### 🧹 Delete All Tasks

```bash
task-cli clear
```

### 📘 View Command Help

```bash
task-cli help
task-cli help add
task-cli help delete
```

### 🧭 List Tasks by Priority Order (high → medium → low)

```bash
task-cli list
```

---

## 🧩 Task Structure

Each task is stored in the `tasks.json` file with the following properties:

```json
{
   "id": 1,
   "description": "Example task",
   "status": "todo",
   "priority": "medium",
   "createdAt": "2025-06-08T10:00:00Z",
   "updatedAt": "2025-06-08T12:30:00Z"
}
```

---

## 📁 Storage Details

- All tasks are stored in a fixed per-user file at `~/.config/task-cli/tasks.json` on Linux
- The app creates the directory automatically on first write
- The app uses Go's native `os` and `encoding/json` packages to read/write the file
- No external libraries or dependencies required

---

## 🧪 Error Handling

The app gracefully handles:

- Invalid command usage (e.g., missing arguments)
- Non-existent task IDs during update, delete, or mark
- Corrupted or unreadable `tasks.json` file

---

## 🛠 Tech Stack

- Language: [Go (Golang)](https://go.dev/)
- File Format: JSON
- No third-party packages

---

## ✅ Final Words

This project is a simple but effective way to get comfortable with:

- Command-line interfaces
- File handling
- JSON serialization
- Basic structuring in Go

---

## 📖 License

MIT License

---

## 🌱 Contributing

Pull requests are welcome! If you find a bug or want to improve something, feel free to open an issue or submit a PR.

---

### 🔗 Official Project Page

> 📌 [https://roadmap.sh/projects/task-tracker](https://roadmap.sh/projects/task-tracker)
