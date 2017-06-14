package image

import (
    "os"
    "path"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "vidalia/config"
)

func TestSmallImage(t *testing.T) {
    cached := path.Join(config.CacheDir, "small.jpg")
    targetDir := path.Join(config.StorageDir, "24")
    target := path.Join(targetDir, "source.jpg")

    image, err := NewImage(cached, 24)

    require.Nil(t, err)
    assert.Equal(t, target, image.path,
        "image.path should be updated with the storage location.")

    if _, err := os.Stat(cached); os.IsExist(err) {
        assert.Fail(t, "Cached file should be destroyed (moved)")
    }
    if _, err := os.Stat(target); os.IsNotExist(err) {
        assert.Fail(t, "Target file should exist")
    }

    err = image.Process()
    require.Nil(t, err)

    /* Reset test image location for subsequent tests */
    defer func() {
        os.Rename(target, cached)
        os.RemoveAll(targetDir)
    }()

    /* Analysis test */
    assert.Equal(t, uint(200), image.width)
    assert.Equal(t, uint(198), image.height)
    assert.Equal(t, uint64(13597467515327605323), image.phash)

    /* Version test */
    source, err := os.Stat(
        path.Join(config.StorageDir, "24/source.jpg"))
    require.Nil(t, err)

    preview, err := os.Stat(
        path.Join(config.StorageDir, "24/preview.jpg"))
    require.Nil(t, err)

    thumb, err := os.Stat(
        path.Join(config.StorageDir, "24/thumbnail.jpg"))
    require.Nil(t, err)

    assert.True(t, os.SameFile(preview, source),
        "preview should be hardlinked to the source")

    assert.True(t, os.SameFile(thumb, source),
        "thumbnail should be hardlinked to the source")
}
