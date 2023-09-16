package manifest

var staticManifests = []*Manifest{
	&FormatJSON,
}

func StaticManifests() []*Manifest {
	return staticManifests
}
