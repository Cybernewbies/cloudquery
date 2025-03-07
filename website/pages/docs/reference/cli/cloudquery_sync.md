---
title: "sync"
---
## cloudquery sync

Sync resources from configured source plugins to destinations

### Synopsis

Sync resources from configured source plugins to destinations

```
cloudquery sync [files or directories] [flags]
```

### Examples

```
# Sync resources from configuration in a directory
cloudquery sync ./directory
# Sync resources from directories and files
cloudquery sync ./directory ./aws.yml ./pg.yml

```

### Options

```
  -h, --help   help for sync
```

### Options inherited from parent commands

```
      --cq-dir string          directory to store cloudquery files, such as downloaded plugins (default ".cq")
      --log-console            enable console logging
      --log-file-name string   Log filename (default "cloudquery.log")
      --log-format string      Logging format (json, text) (default "text")
      --log-level string       Logging level (default "info")
      --no-log-file            Disable logging to file
      --no-telemetry           disable telemetry collection
```

### SEE ALSO

* [cloudquery](/docs/cli/commands/cloudquery)	 - CloudQuery CLI

