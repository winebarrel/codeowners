# codeowners

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
    "has_codeowners": true,
    "content": "* @bob\n"
  },
  {
    "repo": "my-repo-2",
    "has_codeowners": true,
    "content": "* @alice\n"
  },
  {
    "repo": "my-repo-3",
    "has_codeowners": false
  }
]
```
