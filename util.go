package astify

func buildSet(strs []string) map[string]struct{} {
	m := map[string]struct{}{}
	for _, s := range strs {
		m[s] = struct{}{}
	}
	return m
}
