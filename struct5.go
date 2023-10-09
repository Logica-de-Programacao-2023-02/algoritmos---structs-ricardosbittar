package main

type Playlist struct {
	nome    string
	musicas []Musica
}

type Musica struct {
	titulo  string
	artista string
	duracao float64
}

func encontrarMusica(playlists []Playlist, tituloMusica string) []Playlist {
	var playlistsEncontradas []Playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == tituloMusica {
				playlistsEncontradas = append(playlistsEncontradas, playlist)
				break
			}
		}
	}

	return playlistsEncontradas
}
