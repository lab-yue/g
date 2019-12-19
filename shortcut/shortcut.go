package shortcut

// Map used in cli
var Map = map[string][]string{
	"a": {"add"},
	"b": {"branch"},
	"c": {"checkout"},
	"m": {"merge"},
	"s": {"status"},

	"om":  {"origin", "master"},
	"cm":  {"commit", "-m"},
	"rao": {"remote", "add", "origin"},
	"rgo": {"remote", "get-url", "origin"},
	"rso": {"remote", "set-url", "origin"},
	"po":  {"push", "origin"},
	"pom": {"push", "origin", "master"},
	"pox": {"push", "origin", "$(git branch --show-current)"},
	"cb":  {"checkout", "-b"},
	"db":  {"branch", "-d"},
}
