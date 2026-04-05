# Changelog

All notable changes to `ws` will be documented here.
Format loosely follows [Keep a Changelog](https://keepachangelog.com/).

See PLANNED.md for future releases.

## [0.1.2] Move ws binary to project root.
- Installation is now `github.com/alyashour/ws@latest` instead of `.../ws/cmd/ws@latest`.

## [0.1.1] Sync commands and QOL

### Added
- `ws sync init`
- `ws sync status`. Prints the data path, checks if dir exists, if repo exists, then runs `git status`
- `ws sync pull [--hard]`
- `ws sync push [-f]`

Also for tasks
- `ws todo path`. Prints the todo lists current path.

## [0.1.0] Tasks - 13-03-2026

### Added
- `ws task add "text"`
- `ws task list` / `ws todo list --all`
- `ws task done <id>`
- `ws task edit <id> "text"`
- `ws task remove <id>`