package types

const (
	// ModuleName defines the module name
	ModuleName = "agent"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_agent"
)

var (
	ParamsKey = []byte("p_agent")
)

const (
	QuestionKey = "Question/value/"
)

const (
	QuestionCountKey = "Question/count/"
)

const (
	PromptKey = "Prompt/value/"
)

const (
	PromptCountKey = "Prompt/count/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
