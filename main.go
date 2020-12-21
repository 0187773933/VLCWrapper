package main

import (
	"time"
	vlc "github.com/0187773933/VLCWrapper/wrapper"
)

func main() {
	path := "/media/morphs/14TB/MEDIA_MANAGER/TVShows/GilmoreGirls/006 - Gilmore Girls - S01E06 - Rory's Birthday Parties.mp4"
	p := vlc.NewPlayer(nil)
	p.Start()
	p.Play(path)
	time.Sleep( 3 * time.Second )
	p.Fullscreen()
}