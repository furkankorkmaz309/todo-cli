# todo-cli

**todo-cli** is a command-line application written in Go that allows you to manage your to-do tasks directly from the terminal. It provides a user-friendly interface to add, list, update, delete, and archive tasks using a simple text-based menu system.

---

## Features

- ✅ Add new tasks with titles, descriptions, and categories  
- 📋 List tasks by status or category  
- 📝 Update existing tasks  
- ❌ Delete tasks  
- 📂 Archive completed tasks  
- 🔎 Sort tasks alphabetically or by creation date  
- 💾 Stores data locally in JSON files (no external database required)

---

## Installation

> Ensure you have Go 1.18 or later installed on your system.

Clone the repository and build the application:

```bash
git clone https://github.com/furkankorkmaz309/todo-cli.git
cd todo-cli
go build -o todo-cli ./cmd/todo-cli
