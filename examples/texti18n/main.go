// Copyright 2023 The Ebitengine Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This example is a demonstration to render languages that cannot be rendered with the `text` package.
// We plan to provide a useful API to render them more easily (#2454). Stay tuned!

package main

import (
	"bytes"
	_ "embed"
	"log"

	"golang.org/x/text/language"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed NotoSansArabic-Regular.ttf
var arabicTTF []byte

var arabicFaceSource *text.GoTextFaceSource

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(arabicTTF))
	if err != nil {
		log.Fatal(err)
	}
	arabicFaceSource = s
}

//go:embed NotoSansDevanagari-Regular.ttf
var devanagariTTF []byte

var devanagariFaceSource *text.GoTextFaceSource

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(devanagariTTF))
	if err != nil {
		log.Fatal(err)
	}
	devanagariFaceSource = s
}

//go:embed NotoSansThai-Regular.ttf
var thaiTTF []byte

var thaiFaceSource *text.GoTextFaceSource

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(thaiTTF))
	if err != nil {
		log.Fatal(err)
	}
	thaiFaceSource = s
}

var japaneseFaceSource *text.GoTextFaceSource

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	japaneseFaceSource = s
}

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const arabicText = "لمّا كان الاعتراف بالكرامة المتأصلة في جميع"

	op := &text.DrawOptions{}
	op.GeoM.Translate(screenWidth-20, 40)
	text.Draw(screen, arabicText, &text.GoTextFace{
		Source:    arabicFaceSource,
		Direction: text.DirectionRightToLeft,
		Size:      24,
		Language:  language.Arabic,
	}, op)

	const hindiText = "चूंकि मानव परिवार के सभी सदस्यों के जन्मजात गौरव और समान"

	op.GeoM.Reset()
	op.GeoM.Translate(20, 110)
	text.Draw(screen, hindiText, &text.GoTextFace{
		Source:   devanagariFaceSource,
		Size:     24,
		Language: language.Hindi,
	}, op)

	const thaiText = "โดยที่การไม่นำพาและการหมิ่นในคุณค่าของสิทธิมนุษยชน"

	op.GeoM.Reset()
	op.GeoM.Translate(20, 160)
	text.Draw(screen, thaiText, &text.GoTextFace{
		Source:   thaiFaceSource,
		Size:     24,
		Language: language.Thai,
	}, op)

	const japaneseText = "あのイーハトーヴォの\nすきとおった風、\n夏でも底に冷たさを\nもつ青いそら…"

	op.GeoM.Reset()
	op.GeoM.Translate(screenWidth-20, 210)
	op.LineSpacingInPixels = 48
	text.Draw(screen, japaneseText, &text.GoTextFace{
		Source:    japaneseFaceSource,
		Direction: text.DirectionTopToBottomAndRightToLeft,
		Size:      24,
		Language:  language.Japanese,
	}, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Text I18N (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
