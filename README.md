# ws — workspace

A personal engineering workspace CLI. Centralizes projects, todos, journal entries, and notes into a single git-trackable data store.

## Philosophy

- **Files are the source of truth.** Everything lives in `data/` as human-readable YAML. Edit what you want, write what you want. Nothing breaks ws.
- **Git is the backup.** The sync system is just git. Write, edit, and branch commits as you need.

## Installation

Install the tool with `go install github.com/alyashour/ws@latest`.

## Data layout

```
~/.ws/data/
├── todos/
│   └── personal.yaml
# UNIMPLEMENTED (WIP)
├── projects/
│   └── <id>.yaml
├── notes/
│   └── <slug>.yaml
└── journal/
    └── <date>.yaml
```

## Usage

### Todos

See help menu for what's been implemented!

```bash
ws todo add "text"               # add a new todo
ws todo list                     # list open todos
ws todo list --all               # list all todos including done
ws todo done <id>                # mark done
ws todo edit <id> "new text"     # edit text
ws todo remove <id>              # delete
```

### Projects

```bash
ws project new                   # interactive prompt
ws project list                  # list all projects
ws project list --status active  # filter by status
ws project list --category academic
ws project show <id>             # full detail view
ws migrate <path>                # migrate existing project folder
```

### Journal

```bash
ws journal add "text"            # add entry (timestamped today)
ws journal list                  # reverse chronological
ws journal list --project <id>   # entries linked to a project
```

### Notes

```bash
ws note add "title" "body"       # add a note
ws note list                     # list all notes
ws note show <id>                # show full note
ws note remove <id>              # delete
```

### API

```bash
ws api start                     # start local HTTP server (default :8080)
ws api start --port 3001
ws api install-service           # write systemd/launchd config
```

### Sync

```bash
ws sync init
ws sync pull/push                          # git pull + push on data/
```

---

## API endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/projects` | all projects |
| GET | `/projects/:id` | single project |
| GET | `/todos` | all todos |
| GET | `/todos?project=<id>` | todos for a project |
| GET | `/journal` | all journal entries |
| GET | `/notes` | all notes |

---

## Project YAML schema

```yaml
id: neural-net-optimizer
title: Neural Net Optimizer
description: Gradient descent optimizer for sparse networks
category: academic        # academic | personal | club | startup
status: active            # active | archived | idea
tags:
  - ml
  - c
links:
  github: https://github.com/you/nno
  demo: ~
  paper: https://arxiv.org/abs/...
timeline:
  start: 2024-09
  end: ~
collaborators:
  - Alice
repo_path: ~/projects/neural-net-optimizer
todos:
  - id: t001
    text: Write unit tests
    done: false
    due: ~
```

---

## License

MIT