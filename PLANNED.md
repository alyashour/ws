
# PLANNED RELEASES

This is the counterpart to CHANGELOG.md

---

## [0.4.0] API
- `ws api start` — local HTTP server
- `GET /projects`, `/projects/:id`
- `GET /tasks`, `/todos?project=<id>`
- `GET /journal`, `/notes`
- `ws api install-service` — writes systemd/launchd unit file

---

## [0.3.0] Journal & Notes
- `ws journal add` — timestamped daily entries
- `ws journal list` — reverse chronological
- `ws journal list --project` — filter by project
- `ws note add`, `ws note list`, `ws note show`, `ws note remove`

---

## [0.2.1] Tasks linked to projects
- `ws task add --project <id>`
- `ws task list --project <id>`
- `ws project show` now displays linked tasks

---

## [0.2.0] Projects
- `ws project new` — interactive prompt
- `ws project list` — with `--status` and `--category` filters
- `ws project show <id>` — full detail view
- `ws migrate <path>` — scan existing project folder, generate YAML stub

---

## [0.1.2] Tasks upgrades
- Add notifications
    - `ws task notif <id> <notif time>`
    - `ws task notifs` and `ws notifs`

---

## [0.1.1] Sync commands

### Added
Both of these fail on any conflicts unless --hard or -f is present.
- `ws pull [--hard]`
- `ws push [-f]`