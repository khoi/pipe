package manifest

import (
	"io"
	"io/fs"
)

func ListManifests(assets fs.FS) []*Manifest {
	var manifests []*Manifest

	// Get all the files in the director
	manifestsDir, err := fs.ReadDir(assets, "frontend/src/fixtures")
	if err != nil {
		panic(err)
	}

	// Loop over the files
	for _, manifestDir := range manifestsDir {
		manifestDirPath := "frontend/src/fixtures/" + manifestDir.Name()

		// Open the file
		manifestFile, err := assets.Open(
			manifestDirPath + "/manifest.json",
		)
		if err != nil {
			panic(err)
		}
		defer manifestFile.Close()

		manifestBytes, err := io.ReadAll(manifestFile)
		if err != nil {
			panic(err)
		}

		// Get manifest dir as sub filesystem
		manifestFS, err := fs.Sub(assets, manifestDirPath)
		if err != nil {
			panic(err)
		}

		m, err := NewFromBytes(manifestFS, manifestBytes)
		if err != nil {
			panic(err)
		}

		// Add the manifest to the list
		manifests = append(manifests, &m)
	}

	return manifests
}
