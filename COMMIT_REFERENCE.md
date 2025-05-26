# Quick Commit Reference

## Conventional Commit Types

| Type | Use Case | Example |
|------|----------|---------|
| `feat` | New feature | `feat: add user login endpoint` |
| `fix` | Bug fix | `fix: resolve null pointer in auth middleware` |
| `docs` | Documentation | `docs: update API documentation` |
| `refactor` | Code improvement | `refactor: extract validation logic to helper` |
| `test` | Tests | `test: add unit tests for user service` |
| `chore` | Maintenance | `chore: update dependencies` |
| `style` | Code style | `style: fix linting issues` |
| `perf` | Performance | `perf: optimize database queries` |
| `build` | Build system | `build: update Dockerfile` |
| `ci` | CI/CD | `ci: add automated testing workflow` |

## Quick Commands

### Using Git Aliases
```bash
git feat "add user authentication"
git fix "resolve database timeout"
git docs "update README with setup instructions"
git test "add integration tests for API"
```

### Manual Commits
```bash
git add .
git commit -m "feat: add user authentication middleware"
```

### Workflow Commands
```bash
git staged          # Review staged changes
git ac              # Add all and commit (opens editor)
git quick "message" # Quick commit with chore: prefix
```

## Breaking Changes

Add `!` after type or use `BREAKING CHANGE:` in footer:
```bash
feat!: remove deprecated API endpoint
# or
feat: update authentication system

BREAKING CHANGE: Old JWT format is no longer supported
```

## Commit Frequency Tips

✅ **Do:**
- Commit after each logical change
- Commit working code frequently
- Use descriptive commit messages
- Break large features into smaller commits

❌ **Don't:**
- Commit broken code
- Make huge commits with multiple changes
- Use vague messages like "fix stuff"
- Wait too long between commits 