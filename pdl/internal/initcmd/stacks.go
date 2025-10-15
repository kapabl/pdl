package initcmd

import "strings"

var backendChoices = []string{
	"go",
	"java",
	"kotlin",
	"rust",
	"csharp",
	"graphql",
	"proto",
	"ts",
}

var frontendChoices = []string{
	"react",
	"vue",
	"ts",
	"js",
}

var frontendFrameworkSet = map[string]struct{}{
	"react": {},
	"vue":   {},
}

var frontendLanguageSet = map[string]struct{}{
	"ts": {},
	"js": {},
}

func BackendChoices() []string {
	return copyList(backendChoices)
}

func FrontendChoices() []string {
	return copyList(frontendChoices)
}

func FilterSelection(values []string, allowed []string) []string {
	result := make([]string, 0, len(values))
	allowedSet := make(map[string]string, len(allowed))
	for _, entry := range allowed {
		allowedSet[entry] = entry
	}
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		key := strings.ToLower(strings.TrimSpace(value))
		canonical, ok := allowedSet[key]
		if !ok {
			continue
		}
		if _, duplicate := seen[canonical]; duplicate {
			continue
		}
		seen[canonical] = struct{}{}
		result = append(result, canonical)
	}
	return result
}

func ContainsSelection(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func SplitFrontendSelections(values []string) ([]string, []string) {
	frameworks := make([]string, 0)
	languages := make([]string, 0)
	for _, value := range values {
		if _, ok := frontendFrameworkSet[value]; ok {
			frameworks = append(frameworks, value)
			continue
		}
		if _, ok := frontendLanguageSet[value]; ok {
			languages = append(languages, value)
		}
	}
	return frameworks, languages
}

func copyList(values []string) []string {
	result := make([]string, len(values))
	copy(result, values)
	return result
}
