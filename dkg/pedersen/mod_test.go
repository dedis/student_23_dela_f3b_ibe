package pedersen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.dedis.ch/fabric/mino"
	"go.dedis.ch/fabric/mino/minoch"
	"go.dedis.ch/fabric/mino/minogrpc"
	"go.dedis.ch/fabric/mino/minogrpc/routing"
	"go.dedis.ch/kyber/v3"
)

func TestStart(t *testing.T) {
	n := 10

	addrFactory := minoch.AddressFactory{}

	rootAddr := addrFactory.FromText([]byte("127.0.0.1:2000"))

	factory := routing.NewTreeRoutingFactory(4, rootAddr, addrFactory)

	addrs := make([]mino.Address, n)
	pubKeys := make([]kyber.Point, n)
	privKeys := make([]kyber.Scalar, n)

	// manager := minoch.NewManager()
	minos := make([]*minogrpc.Minogrpc, n)
	pedersens := make([]*Pedersen, n)

	for i := 0; i < n; i++ {
		addrs[i] = addrFactory.FromText([]byte(fmt.Sprintf("127.0.0.1:2%03d", i)))
		privKeys[i] = suite.Scalar().Pick(suite.RandomStream())
		pubKeys[i] = suite.Point().Mul(privKeys[i], nil)
		minogrpc, err := minogrpc.NewMinogrpc(addrs[i].String(), factory)
		require.NoError(t, err)
		minos[i] = &minogrpc
	}

	for _, minogrpc := range minos {
		minogrpc.AddNeighbours(minos...)
	}

	for i := 0; i < n; i++ {
		pedersen, err := NewPedersen(pubKeys, privKeys[i], minos[i], addrs, suite)
		require.NoError(t, err)
		pedersens[i] = pedersen
	}

	players := &fakePlayers{
		players: addrs,
	}

	err := pedersens[0].Start(players, uint32(n))
	require.NoError(t, err)

	message := []byte("Hello world")
	K, C, remainder, err := pedersens[0].Encrypt(message)
	require.NoError(t, err)
	require.Empty(t, remainder)

	decryptedMessage, err := pedersens[0].Decrypt(K, C)
	require.NoError(t, err)
	fmt.Println("Here is the decrypted message:", string(decryptedMessage))
}

// ----------------------------------------------------------------------------
// Utility functions

// fakePlayers implements mino.Players{}
type fakePlayers struct {
	players  []mino.Address
	iterator *fakeAddressIterator
}

// AddressIterator implements mino.Players.AddressIterator()
func (p *fakePlayers) AddressIterator() mino.AddressIterator {
	if p.iterator == nil {
		p.iterator = &fakeAddressIterator{players: p.players}
	}
	return p.iterator
}

// Len() implements mino.Players.Len()
func (p *fakePlayers) Len() int {
	return len(p.players)
}

// Take ...
func (p *fakePlayers) Take(filters ...mino.FilterUpdater) mino.Players {
	f := mino.ApplyFilters(filters)
	players := make([]mino.Address, len(p.players))
	for i, k := range f.Indices {
		players[i] = p.players[k]
	}
	return &fakePlayers{
		players: players,
	}
}

// fakeAddressIterator implements mino.addressIterator{}
type fakeAddressIterator struct {
	players []mino.Address
	cursor  int
}

// HasNext implements mino.AddressIterator.HasNext()
func (it *fakeAddressIterator) HasNext() bool {
	return it.cursor < len(it.players)
}

// GetNext implements mino.AddressIterator.GetNext(). It is the responsibility
// of the caller to check there is still elements to get. Otherwise it may
// crash.
func (it *fakeAddressIterator) GetNext() mino.Address {
	p := it.players[it.cursor]
	it.cursor++
	return p
}
