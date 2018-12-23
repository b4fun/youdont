workflow "New workflow" {
  on = "push"
  resolves = ["test"]
}

action "test" {
  uses = "actions/npm@e7aaefe"
}
