package shortcut

// Map used in cli
var Map = map[string][]string{
	"a": {"add"},
	"b": {"branch"},
	"c": {"checkout"},
	"m": {"merge"},

	"om":  {"origin", "master"},
	"cm":  {"commit", "-m"},
	"rao": {"remote", "add", "origin"},
	"po":  {"push", "origin"},
	"pom": {"push", "origin", "master"},
	"cb":  {"checkout", "-b"},
	"db":  {"branch", "-d"},
}
