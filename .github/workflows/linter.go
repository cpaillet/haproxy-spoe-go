workflow "Go" {
  on = "pull_request"
  resolves = ["gofmt"]
}

action "gofmt" {
  uses    = "sjkaliski/go-github-actions/fmt@v1.0.0"
  needs   = "previous-action"
  secrets = ["GITHUB_TOKEN"]

  env {
    GO_WORKING_DIR = "./"
  }
}
