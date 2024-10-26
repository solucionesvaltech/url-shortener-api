package domain

type URL struct {
	Short    string `json:"short"`
	Original string `json:"original"`
	Enabled  bool   `json:"enabled"`
}
