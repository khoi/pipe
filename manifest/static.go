package manifest

var staticManifests = []*Manifest{
	&FormatJSON,
	&Base64Encode,
	&Base64Decode,
	&ShuffleLines,
	&UUIDGenerate,
	&JwtDecode,
}

func StaticManifests() []*Manifest {
	return staticManifests
}
