# TODO CLI Tool

golang + sqlite

| Command                           | Example                            | Behavior          |
| --------------------------------- | ---------------------------------- | ----------------- |
| `add <task>`                        | `go run main.go add "Buy milk"`      | Add a new task    |
| `list`                              | `go run main.go list`                | Show all tasks    |
| `done <id>`                         | `go run main.go done 1`              | Mark task as done |
| `delete <id>`                       | `go run main.go delete 2`            | Delete a task     |
| `update <id> <new task>` (optional) | `go run main.go update 3 "Buy eggs"` | Edit a task       |

Run project:
`go run ./cmd/todo <cmd>`

currently works:
`go run ./cmd/todo add "new task"`
`go run ./cmd/todo list`

