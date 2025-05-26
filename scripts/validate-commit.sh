#!/bin/bash

# Validate commit message format
# Usage: ./scripts/validate-commit.sh "commit message"

commit_msg="$1"

# Conventional commit regex pattern
pattern="^(feat|fix|docs|style|refactor|perf|test|build|ci|chore)(\(.+\))?(!)?: .{1,50}"

if [[ $commit_msg =~ $pattern ]]; then
    echo "✅ Valid conventional commit message"
    exit 0
else
    echo "❌ Invalid commit message format"
    echo ""
    echo "Expected format: <type>[optional scope]: <description>"
    echo ""
    echo "Valid types: feat, fix, docs, style, refactor, perf, test, build, ci, chore"
    echo ""
    echo "Examples:"
    echo "  feat: add user authentication"
    echo "  fix(api): resolve timeout issue"
    echo "  docs: update README with setup instructions"
    echo ""
    exit 1
fi 