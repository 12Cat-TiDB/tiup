# Multiple Repositories for TiUP

> This is part of a PingCAP Hackthon 2022 project, it may or may not be available in future TiUP releases.

> This documentation is in very early drafting stage, things may break or be modified anytime.

## Implementation Designs

### Metadata Storage

A new `mirrors.yaml` (or `.ini`?) file will be used to store the information of multiple repositories used by TiUP. Main fields of it are:

 - repository list `mirrors = [string]`
   - e.g.: `tiup-mirrors.pingcap.com`
   - an array of repositories, in the same format as what used for `tiup mirror set` now (URL or file path)
   - the natural (literal) order of repositories in the list **is** the order of their priorities, when packages with the same name are available in more than one repos, the first in the list is used

 - package alias list `aliases = [{string: string}]`
   - e.g.: `zstd: "12cat.github.io/zstd"`
   - users can specify an alias for a package name to override repository priority
   - when alias is set, only the specified repository will be searched for the package
   - no `:` allowed in the alias, thus no version binding

### Manifests

Manifests are stored in subdirectories by each repository, where `/`, `\`, and `:` characters in the address are replaced by `_`.

#### Root

Stored in `$TIUP_HOME/trusted/<repo_addr>`.

#### Other Manifests

Stored in `$TIUP_HOME/manifests/<repo_addr>`.

## TUI Design

TBD.
