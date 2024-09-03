# codeowners

[![CI](https://github.com/winebarrel/codeowners/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/codeowners/actions/workflows/ci.yml)

List up CODEOWNERS of all repositories.

## Usage

```
Usage: codeowners --user=STRING --org=STRING --token=STRING [flags]

Flags:
  -h, --help            Show help.
  -u, --user=STRING     Organization name.
  -o, --org=STRING      Organization name.
      --token=STRING    Authentication token for github.com API requests ($GITHUB_TOKEN).
      --version
```

```sh
$ codeowners -o my-org
[
  {
    "repo": "my-repo-1",
    "exists": true,
    "content": "* @bob\n"
  },
  {
    "repo": "my-repo-2",
    "exists": true,
    "content": "* @alice\n"
  },
  {
    "repo": "my-repo-3",
    "exists": false
  }
]
```
