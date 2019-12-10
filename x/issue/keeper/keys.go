package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"strings"
)

// Key for getting a the next available proposalID from the store
var (
	KeyDelimiter       = ":"
	KeyNextIssueID     = []byte("newIssueID")
	KeyFirstIssueDenom = []byte("firstIssueDenom")
	KeyLastIssueDenom  = []byte("lastIssueDenom")
)

//func BytesString(b []byte) string {
//	return *(*string)(unsafe.Pointer(&b))
//}
// Key for getting a specific issuer from the store
func KeyIssuer(denom string) []byte {
	return []byte(fmt.Sprintf("issues:%s", denom))
}

// Key for getting a specific address from the store
func KeyAddressIssues(addr string) []byte {
	return []byte(fmt.Sprintf("address:%s", addr))
}

// Key for getting a specific allowed from the store
func KeyAllowance(issueID string, owner sdk.AccAddress, spender sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("allowed:%s:%s:%s", issueID, owner.String(), spender.String()))
}

func KeyFreeze(issueID string, accAddress sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("freeze:%s:%s", issueID, accAddress.String()))
}

func PrefixFreeze(issueID string) []byte {
	return []byte(fmt.Sprintf("freeze:%s", issueID))
}

func KeySymbolIssues(symbol string) []byte {
	return []byte(fmt.Sprintf("symbol:%s", strings.ToUpper(symbol)))
}

func KeyIssueIdStr(seq uint64) string {
	return fmt.Sprintf("%s%x", types.IDPreStr, seq)
}
