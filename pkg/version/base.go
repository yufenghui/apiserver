package version

var (
	gitTag string = ""
	// sha1 from git, output of $(git rev-parse HEAD)
	gitCommit string = "$Format:%H$"
	// state of git tree, either "clean" or "dirty"
	gitTreeState string = "not a git tree"
	// build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	buildDate string = "1970-01-01T00:00:00Z"
)
