package main

import (
	"image"
	"image/color"
	"log"

	"github.com/tajtiattila/blur"
)

func process(current, previous image.Image) image.Image {
	img := AbsDiff(blur.Gaussian(previous, 18, blur.ReuseSrc), blur.Gaussian(current, 18, blur.ReuseSrc))
	img = blur.Gaussian(img, 12, blur.ReuseSrc)
	return Threshold(img, 10)
}

func Threshold(source image.Image, distance uint8) image.Image {
	b := source.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, b.Max.X, b.Max.Y))

	var maxDist = uint16(distance) * (256)
	var high = uint32(65535 - maxDist)
	var low = uint32(maxDist)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {

			r, g, b, _ := source.At(x, y).RGBA()
			rn, gn, bn := 65535, 65535, 65535

			if r > high || r < low {
				rn, gn, bn = 0, 0, 0
			}

			if g > high || g < low {
				rn, gn, bn = 0, 0, 0
			}

			if b > high || b < low {
				rn, gn, bn = 0, 0, 0
			}

			c := color.RGBA64{uint16(rn), uint16(gn), uint16(bn), 65535}
			result.Set(x, y, c)
		}
	}

	return result
}

func AbsDiff(source, target image.Image) image.Image {
	b := source.Bounds()
	diff := image.NewRGBA(image.Rect(0, 0, b.Max.X, b.Max.Y))

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r1, g1, b1, _ := source.At(x, y).RGBA()
			r2, g2, b2, _ := target.At(x, y).RGBA()

			rd := abs(r1, r2)
			gd := abs(g1, g2)
			bd := abs(b1, b2)

			c := color.RGBA64{rd, gd, bd, 65535}

			diff.Set(x, y, c)
		}
	}

	return diff
}

func abs(a, b uint32) uint16 {
	if a < b {
		return uint16(b - a)
	}
	return uint16(a - b)
}

func HotPixels(current image.Image) (hot int, pct float64) {
	total := 0
	b := current.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			total++
			r, g, b, _ := current.At(x, y).RGBA()
			if r > 0 || g > 0 || b > 0 {
				hot++
			}
		}
	}

	return hot, (float64(hot) / float64(total) * 100.0)
}

func Extract(source, target image.Image) image.Image {
	b := source.Bounds()
	extracted := image.NewRGBA(image.Rect(0, 0, b.Max.X, b.Max.Y))

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			rs, gs, bs, _ := source.At(x, y).RGBA()
			rt, gt, bt, _ := target.At(x, y).RGBA()

			if rt > 0 || gt > 0 || bt > 0 {
				c := color.RGBA64{uint16(rs), uint16(gs), uint16(bs), 65535}

				extracted.Set(x, y, c)
			}
		}
	}

	return extracted
}

func Contrast(img image.Image, contrast int) image.Image {

	b := img.Bounds()
	new := image.NewRGBA(image.Rect(0, 0, b.Max.X, b.Max.Y))

	factor := float64(259*(contrast+255)) / float64(255*(259-contrast))
	log.Println(factor)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			rn := factor*(float64(r)-32640) + 32640
			gn := factor*(float64(g)-32640) + 32640
			bn := factor*(float64(b)-32640) + 32640

			// log.Println(rn, gn, bn)

			c := color.RGBA64{uint16(rn), uint16(gn), uint16(bn), 65535}
			new.Set(x, y, c)
		}
	}

	return new
}
