package utils

import "fmt"

func DisplayOption() {
	fmt.Println(`╔══════════════════════════════════════════════════════╗
║              📋 TASK MANAGER CLI TOOL                ║
║     Efficient. Minimal. Built for Developers.        ║
╚══════════════════════════════════════════════════════╝

Choose an action to perform:
────────────────────────────────────────────
1️⃣  Create a New Task        → Add a task with title, description, and priority.
2️⃣  View All Tasks           → Display a list of existing tasks.
3️⃣  Update an Existing Task  → Modify the details of a chosen task.
4️⃣  Delete a Task            → Permanently remove a task by ID.
5️⃣  Exit                     → Close the CLI Task Manager.
────────────────────────────────────────────`)
}
