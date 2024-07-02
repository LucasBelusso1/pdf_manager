command_exists() {
  command -v "$1" >/dev/null 2>&1
}

if ! command_exists air; then
  echo "Installing air..."
  go install github.com/cosmtrek/air@latest
fi

echo "Starting development server with live reload..."
air