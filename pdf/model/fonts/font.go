/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package fonts

import (
	"fmt"

	"github.com/unidoc/unidoc/pdf/core"
	"github.com/unidoc/unidoc/pdf/internal/textencoding"
)

// Font represents a font which is a series of glyphs. Character codes from PDF strings can be
// mapped to and from glyphs.  Each glyph has metrics.
type Font interface {
	Encoder() textencoding.TextEncoder
	SetEncoder(encoder textencoding.TextEncoder)
	GetGlyphCharMetrics(glyph textencoding.GlyphName) (CharMetrics, bool)
	ToPdfObject() core.PdfObject
}

// CharMetrics represents width and height metrics of a glyph.
type CharMetrics struct {
	GlyphName textencoding.GlyphName
	Wx        float64
	Wy        float64
}

func (m CharMetrics) String() string {
	return fmt.Sprintf("<%q,%.1f,%.1f>", m.GlyphName, m.Wx, m.Wy)
}
