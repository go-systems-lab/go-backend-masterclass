#!/bin/bash

# Git aliases for conventional commits and frequent commit workflow
# Run this script to set up helpful git aliases

echo "Setting up git aliases for conventional commits..."

# Conventional commit aliases
git config alias.feat '!f() { git commit -m "feat: $1"; }; f'
git config alias.fix '!f() { git commit -m "fix: $1"; }; f'
git config alias.docs '!f() { git commit -m "docs: $1"; }; f'
git config alias.refactor '!f() { git commit -m "refactor: $1"; }; f'
git config alias.test '!f() { git commit -m "test: $1"; }; f'
git config alias.chore '!f() { git commit -m "chore: $1"; }; f'
git config alias.style '!f() { git commit -m "style: $1"; }; f'
git config alias.perf '!f() { git commit -m "perf: $1"; }; f'
git config alias.build '!f() { git commit -m "build: $1"; }; f'
git config alias.ci '!f() { git commit -m "ci: $1"; }; f'

# Workflow aliases
git config alias.ac '!git add -A && git commit'
git config alias.acp '!git add -A && git commit && git push'
git config alias.quick '!f() { git add -A && git commit -m "chore: $1" && git push; }; f'

# Review aliases
git config alias.staged 'diff --staged'
git config alias.unstage 'reset HEAD --'
git config alias.last 'log -1 HEAD'

echo "✅ Git aliases configured successfully!"
echo ""
echo "Available aliases:"
echo "  git feat 'description'     - feat: description"
echo "  git fix 'description'      - fix: description"
echo "  git docs 'description'     - docs: description"
echo "  git refactor 'description' - refactor: description"
echo "  git test 'description'     - test: description"
echo "  git chore 'description'    - chore: description"
echo ""
echo "Workflow aliases:"
echo "  git ac                     - add all and commit (interactive)"
echo "  git acp                    - add all, commit, and push"
echo "  git quick 'message'       - add all, commit with chore:, and push"
echo ""
echo "Review aliases:"
echo "  git staged                 - show staged changes"
echo "  git unstage                - unstage files"
echo "  git last                   - show last commit" 