package comb

import "fmt"

// Token accepts the shortest given token. At least one token must
// be provided. If more than one token is given, then a trie is used
// to check for membership.
func Token(tokens ...[]rune) Parser {
	if len(tokens) == 0 {
		panic("at least one token must be specified")
	}

	for _, tok := range tokens {
		if len(tok) == 0 {
			panic("token cannot be empty")
		}
	}

	if len(tokens) == 1 {
		return singleToken(tokens[0])
	}

	return manyTokens(tokens)
}

// StringToken is like Token, but takes multiple strings.
func StringToken(tokens ...string) Parser {
	rTokens := make([][]rune, len(tokens))
	for i, s := range tokens {
		rTokens[i] = []rune(s)
	}

	return Token(rTokens...)
}

func singleToken(runes []rune) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		var r rune
		next := s
		var err error

		for _, c := range runes {
			r, next, err = next.Next()
			if err != nil {
				return Failed(err), next
			}

			if r != c {
				return Failed(tokenError(s.Between(next))), next
			}
		}

		return Result{
			Runes: s.Between(next),
		}, next
	})
}

func manyTokens(tokens [][]rune) Parser {
	t := buildTrie(tokens)

	return ParserFunc(func(s Scanner) (Result, Scanner) {
		t := t

		var r rune
		next := s
		var err error

		for !t.accept {
			r, next, err = next.Next()
			if err != nil {
				return Failed(err), next
			}

			t = t.find(r)
			if t == nil {
				return Failed(tokenError(s.Between(next))), next
			}
		}

		return Result{
			Runes: s.Between(next),
		}, next
	})
}

func tokenError(runes []rune) error {
	return errorFunc(func() string {
		prefix := string(runes)
		return fmt.Sprintf("'%s' is not a prefix of any token", prefix)
	})
}

type tokenTrie struct {
	children map[rune]*tokenTrie
	accept   bool
}

func buildTrie(tokens [][]rune) *tokenTrie {
	t := &tokenTrie{}

	for _, s := range tokens {
		t.add(s)
	}

	return t
}

func (t *tokenTrie) add(runes []rune) {
	if len(runes) == 0 {
		t.accept = true
		return
	}

	r, runes := runes[0], runes[1:]

	if t.children == nil {
		t.children = make(map[rune]*tokenTrie)
	}

	child := t.children[r]
	if child == nil {
		child = &tokenTrie{}
		t.children[r] = child
	}

	child.add(runes)
}

func (t *tokenTrie) find(r rune) *tokenTrie {
	return t.children[r]
}
