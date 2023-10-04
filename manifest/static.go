package manifest

var staticManifests = []*Manifest{
	&Base64Decode,
	&Base64Encode,
	&FormatJSON,
	&JwtDecode,
	&LoadFromURL,
	&ShuffleLines,
	&SortLines,
	&UUIDGenerate,
}

func StaticManifests() []*Manifest {
	return staticManifests
}
