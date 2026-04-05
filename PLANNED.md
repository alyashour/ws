
# PLANNED RELEASES

This is the counterpart to CHANGELOG.md

## [0.6.0] Multi-user
- `ws task <id> loop @other-user`
- Needs some more design on how user auth & setup will work. Ideally decentralized.

## [0.5.0] Task Splitting and Editing
- Tasks now form a tree
- `ws task split <id> -c <title 1> -c <title 2>`. This returns a list of all the children's IDs.
- `ws task edit <id> [--parent <id>] [--title <new title>] [--due-date <new due date>] ...`
- No cyclical dependencies are allowed.

## [0.4.0] API
- `ws api start` — local HTTP server
- `GET /projects`, `/projects/:id`
- `GET /tasks`, `/todos?project=<id>`
- `GET /journal`, `/notes`
- `ws api install-service` — writes systemd/launchd unit file

## [0.3.0] Journal & Notes
- `ws journal add` — timestamped daily entries
- `ws journal list` — reverse chronological
- `ws journal list --project` — filter by project
- `ws note add`, `ws note list`, `ws note show`, `ws note remove`

## [0.2.1] Tasks linked to projects
- `ws task add --project <id>`
- `ws task list --project <id>`
- `ws project show` now displays linked tasks

## [0.2.0] Projects
- `ws project new` — interactive prompt
- `ws project list` — with `--status` and `--category` filters
- `ws project show <id>` — full detail view
- `ws migrate <path>` — scan existing project folder, generate YAML stub

## [0.1.4] Tasks upgrades
- Add notifications
    - `ws task notif <id> <notif time>`
    - `ws task notifs` and `ws notifs`

## [0.1.3] Customization and QOL
- Add path config
    - `ws tasks path -l/--list`
    - `ws tasks path -n/--new <new path>`
- Add easy reset
    - `ws task reset -f/--force`
- Improve help menus (both dev user and dev experience)
- Add `ws --version\-v`
- Clean up code TODOs

