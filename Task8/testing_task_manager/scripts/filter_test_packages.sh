#!/bin/bash

# List of packages to exclude (adjust as needed)
exclude_packages=(
  "github.com/Go-learning-path/Task8/testing_task_manager/Repositories/task_repository_test"
  "github.com/Go-learning-path/Task8/testing_task_manager/Repositories/user_repository_test"
)

# Find all packages with .go files
all_packages=$(go list ./... | grep -v "/vendor/")

# Filter out the excluded packages
for exclude in "${exclude_packages[@]}"; do
  all_packages=$(echo "$all_packages" | grep -v "$exclude")
done

# Convert to a space-separated list
test_packages_list=$(echo "$all_packages" | tr '\n' ' ')

# Print the list of packages
echo "$test_packages_list"