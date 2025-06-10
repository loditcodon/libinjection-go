package libinjection

// NewDebugState xuất hiện public để bạn dùng
func NewDebugState(input string) *State {
	// flags = 0 (không bật biến thể DB đặc thù)
	return newState(input, len(input), 0)
}
