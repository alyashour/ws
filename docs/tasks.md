# Tasks

Tasks are simple data structures that contain 6 fields:
- an *ID*
- some *Text*
- a *Done* flag
- a *CreatedAt* stamp
- an optional *Due* date

They are stored as entries inside of a list stored in yaml.

## Schema
```yaml
todos:
  - id: t00-01 # 6 digit hex
    text: Fix the backward pass bug
    done: false
    created_at: 2024-01-15
    due: 2024-01-20        # optional
```