package image

import (
    "math"
    "gopkg.in/gographics/imagick.v3/imagick"
    "vidalia/config"
)

func (image *Image) generateRasterVersions(wand *imagick.MagickWand) (err error) {
    ratio := float64(image.Width) / float64(image.Height)

    for version, width := range config.ImageVersions {
        if width < image.Width {
            err = image.createRasterVersion(wand, version, width, ratio)
        } else {
            err = image.linkVersionToImage(version)
        }

        if err != nil { break }
    }

    return err
}

func (image *Image) createRasterVersion(wand *imagick.MagickWand,
        version string, width uint, ratio float64) (err error) {
    path := image.versionStoragePath(version)
    height := uint(math.Floor(float64(width) / ratio))

    versionWand := wand.Clone()
    defer versionWand.Destroy()

    err = versionWand.ResizeImage(width, height, imagick.FILTER_UNDEFINED)
    if err != nil { return err }

    err = versionWand.SetImageCompressionQuality(95)
    if err != nil { return err }

    err = versionWand.WriteImage(path)
    return err
}