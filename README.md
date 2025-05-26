# Go Backend Masterclass

A comprehensive Go backend development course project with automated commit workflows and conventional commit standards.

## 🚀 Quick Start

### Git Aliases Setup

This project includes pre-configured git aliases for streamlined conventional commits. Run the setup script:

```bash
./scripts/git-aliases.sh
```

### Commit Template Configuration

Configure git to use our commit message template:

```bash
git config commit.template .gitmessage
```

### Required Tools Installation

Install SQLC for SQL code generation:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Install golang-migrate for database migrations:

```bash
brew install golang-migrate
```

> **Note:** 
> - **SQLC** generates type-safe Go code from SQL queries
> - **golang-migrate** handles database schema migrations and versioning
> - Both tools are essential for this project's database layer implementation

## 📝 Available Git Aliases

### Conventional Commit Shortcuts

| Alias | Command | Example Usage |
|-------|---------|---------------|
| `git feat` | `git commit -m "feat: <message>"` | `git feat "add user authentication"` |
| `git fix` | `git commit -m "fix: <message>"` | `git fix "resolve database timeout"` |
| `git docs` | `git commit -m "docs: <message>"` | `git docs "update API documentation"` |
| `git refactor` | `git commit -m "refactor: <message>"` | `git refactor "extract validation logic"` |
| `git test` | `git commit -m "test: <message>"` | `git test "add unit tests for auth"` |
| `git chore` | `git commit -m "chore: <message>"` | `git chore "update dependencies"` |
| `git style` | `git commit -m "style: <message>"` | `git style "fix linting issues"` |
| `git perf` | `git commit -m "perf: <message>"` | `git perf "optimize database queries"` |
| `git build` | `git commit -m "build: <message>"` | `git build "update Dockerfile"` |
| `git ci` | `git commit -m "ci: <message>"` | `git ci "add automated testing"` |

### Workflow Aliases

| Alias | Command | Description |
|-------|---------|-------------|
| `git ac` | `git add -A && git commit` | Add all changes and commit (opens editor) |
| `git acp` | `git add -A && git commit && git push` | Add, commit, and push |
| `git quick` | `git add -A && git commit -m "chore: <message>" && git push` | Quick commit with chore prefix |

### Review Aliases

| Alias | Command | Description |
|-------|---------|-------------|
| `git staged` | `git diff --staged` | Show staged changes |
| `git unstage` | `git reset HEAD --` | Unstage files |
| `git last` | `git log -1 HEAD` | Show last commit |

## 🛠️ Development Workflow

### 1. Making Changes

```bash
# Make your code changes
# Stage specific changes (recommended)
git add -p

# Or stage all changes
git add .

# Review what's staged
git staged
```

### 2. Committing Changes

```bash
# Using aliases (recommended)
git feat "add user login endpoint"
git fix "resolve null pointer in middleware"
git docs "update setup instructions"

# Or traditional way
git commit -m "feat: add user login endpoint"
```

### 3. Quick Workflow

```bash
# For small changes/fixes
git quick "update configuration"

# For regular workflow
git ac  # Opens editor for commit message
```

## 📋 Commit Message Validation

Validate your commit messages before committing:

```bash
./scripts/validate-commit.sh "feat: add user authentication"
```

## 📚 Documentation

- **[CONTRIBUTING.md](CONTRIBUTING.md)** - Detailed contribution guidelines
- **[COMMIT_REFERENCE.md](COMMIT_REFERENCE.md)** - Quick commit reference
- **[.gitmessage](.gitmessage)** - Commit message template

## 🎯 Commit Best Practices

### ✅ Do
- Make small, frequent commits
- Use descriptive commit messages
- Follow conventional commit format
- Commit working code
- Break features into logical commits

### ❌ Don't
- Make huge commits with multiple changes
- Commit broken code
- Use vague messages like "fix stuff"
- Wait too long between commits

### Example: Good Commit Sequence
```bash
git feat "add user model struct"
git feat "implement user repository interface"  
git feat "add user service layer"
git test "add unit tests for user service"
git docs "update API documentation for users"
```

## 🔧 Manual Alias Configuration

If you prefer to set up aliases manually:

```bash
# Conventional commit aliases
git config alias.feat '!f() { git commit -m "feat: $1"; }; f'
git config alias.fix '!f() { git commit -m "fix: $1"; }; f'
git config alias.docs '!f() { git commit -m "docs: $1"; }; f'

# Workflow aliases  
git config alias.ac '!git add -A && git commit'
git config alias.staged 'diff --staged'
git config alias.last 'log -1 HEAD'
```

## 📖 Conventional Commit Types

| Type | Description | Version Impact |
|------|-------------|----------------|
| `feat` | New feature | Minor version bump |
| `fix` | Bug fix | Patch version bump |
| `docs` | Documentation changes | Patch version bump |
| `refactor` | Code refactoring | Patch version bump |
| `test` | Adding/updating tests | Patch version bump |
| `chore` | Maintenance tasks | Patch version bump |
| `perf` | Performance improvements | Patch version bump |
| `style` | Code style changes | Patch version bump |
| `build` | Build system changes | Patch version bump |
| `ci` | CI/CD changes | Patch version bump |

### Breaking Changes
Add `!` after type for breaking changes:
```bash
git feat! "remove deprecated API endpoint"
```

---

**Happy coding! 🚀** Remember to commit early and often with meaningful messages.