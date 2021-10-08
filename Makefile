build:
	go build -o bin/TaskSyncAlerts entrypoint/task-sync-alerts/main.go
	go build -o bin/TaskSyncTeams entrypoint/task-sync-teams/main.go
	go build -o bin/TaskSyncIntegrations entrypoint/task-sync-integrations/main.go
test:
	gotestsum --format=pkgname --junitfile tests.xml -- -race -covermode=atomic -coverprofile cover.out ./...

watch:
	gotestsum --format=pkgname --watch -- -race ./...
