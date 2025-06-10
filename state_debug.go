package libinjection

type DebugToken struct {
	Type  byte   // ký tự fingerprint, vd 'U', '1', '&', 's'...
	Value string // đoạn text gốc
	From  int    // offset trong chuỗi
	To    int
}

func (s *State) DebugTokens() []DebugToken {
	var out []DebugToken
	for i := 0; i < int(s.toklen); i++ {
		tv := s.tokenvec[i]
		out = append(out, DebugToken{
			Type:  tv.Type,
			Value: string(s.s[tv.Pos : tv.Pos+tv.Len]),
			From:  tv.Pos,
			To:    tv.Pos + tv.Len,
		})
	}
	return out
}
