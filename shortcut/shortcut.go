package shortcut

// Map used in cli
var Map = map[string][]string{
	"a": {"add"},
	"b": {"branch"},
	"m": {"merge"},

	"om":  {"origin", "master"},
	"cm":  {"commit", "-m"},
	"rao": {"remote", "add", "origin"},
	"po":  {"push", "origin"},
	"pom": {"push", "origin", "master"},
	"cb":  {"checkout", "-b"},
}
