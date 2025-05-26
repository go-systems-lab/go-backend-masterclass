# Contributing to Go Backend Masterclass

## Commit Guidelines

This project follows [Conventional Commits](https://www.conventionalcommits.org/) specification for consistent commit messages and automated changelog generation.

### Commit Frequency

- **Make small, frequent commits** rather than large, infrequent ones
- Each commit should represent a single logical change
- Commit early and often to maintain a clear development history
- Break down large features into smaller, manageable commits

### Commit Message Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Quick Reference

| Type | Description | When to Use |
|------|-------------|-------------|
| `feat` | New feature | Adding new functionality |
| `fix` | Bug fix | Fixing a bug or issue |
| `docs` | Documentation | README, comments, documentation updates |
| `refactor` | Code refactoring | Improving code without changing functionality |
| `test` | Tests | Adding or updating tests |
| `chore` | Maintenance | Dependencies, build scripts, etc. |

### Examples of Good Commit Practices

#### ✅ Good - Small, focused commits:
```bash
feat: add user model struct
feat: implement user repository interface
feat: add user service layer
test: add unit tests for user service
docs: update API documentation for user endpoints
```

#### ❌ Bad - Large, unfocused commit:
```bash
feat: implement complete user management system with tests and docs
```

### Setting Up Commit Template

Configure git to use the provided commit message template:

```bash
git config commit.template .gitmessage
```

### Pre-commit Checklist

Before committing, ensure:
- [ ] Code compiles without errors
- [ ] Tests pass (if applicable)
- [ ] Commit message follows conventional format
- [ ] Changes are focused and atomic
- [ ] No sensitive information is included

### Workflow Tips

1. **Stage changes incrementally**: Use `git add -p` to stage specific changes
2. **Review before committing**: Use `git diff --staged` to review staged changes
3. **Use meaningful scopes**: Add scope when working on specific modules (e.g., `feat(auth):`, `fix(db):`)
4. **Write descriptive messages**: Explain what and why, not just what changed

### Breaking Changes

For breaking changes, add `!` after the type or include `BREAKING CHANGE:` in the footer:

```bash
feat!: remove deprecated user endpoint
# or
feat: update user authentication

BREAKING CHANGE: The old /auth endpoint has been removed. Use /v2/auth instead.
``` 