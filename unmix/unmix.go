package unmix

import (
	"sync"

	"github.com/gellel/amiibo/amiibo"
	"github.com/gellel/amiibo/game"
	"github.com/gellel/amiibo/mix"
)

// Amiibo unmixes the mix.Amiibo map into a sequence of
// normalized amiibo.Amiibo. Performs the normalization in
// O(N) time.
func Amiibo(m map[string]*mix.Amiibo) []*amiibo.Amiibo {
	var (
		s  = []*amiibo.Amiibo{}
		wg sync.WaitGroup
	)
	for _, m := range m {
		wg.Add(1)
		go func(m *mix.Amiibo) {
			defer wg.Done()
			var (
				a, err = amiibo.NewAmiibo(m.Compatability, m.Item, m.Lineup)
			)
			if err != nil {
				return
			}
			s = append(s, a)
		}(m)
	}
	wg.Wait()
	return s
}

// Game unmixes the mix.Game map into a sequence of
// normalized game.Game. Performs the normalization in
// O(N) time.
func Game(m map[string]*mix.Game) []*game.Game {
	var (
		s  = []*game.Game{}
		wg sync.WaitGroup
	)
	for _, m := range m {
		wg.Add(1)
		go func(m *mix.Game) {
			defer wg.Done()
			var (
				g, err = game.NewGame(m.Game, m.Item)
			)
			if err != nil {
				return
			}
			s = append(s, g)
		}(m)
	}
	wg.Wait()
	return s
}

func Unmix(m *mix.Mix) ([]*amiibo.Amiibo, []*game.Game) {
	var (
		amiibo []*amiibo.Amiibo
		game   []*game.Game
		wg     sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		amiibo = Amiibo(m.Amiibo)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		game = Game(m.Games)
	}()
	wg.Wait()
	return amiibo, game
}
