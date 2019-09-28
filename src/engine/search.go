package engine

type SearchEngine interface {
	Search(txt string, pat string) []int
}
