package main

type Playlist struct {
	nome    string
	musicas []musicas
}
type musicas struct {
	titulo  string
	artista string
	duracao int
}

func nomePlaylist(Playlist Playlist) (string, float64, []musicas) {
	nomePlaylist := Playlist.nome
	var duracaoTotal float64
	var musicasInfo []musicas

	for _, musica := range Playlist.musicas {
		duracaoTotal += musicas.duracao
		musicasInfo = append(musicasInfo, musica)
	}

	return nomePlaylist, duracaoTotal, musicasInfo
}
