package filterx

import (
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/pngx"
)

const (
	SourcePath01 = "test/source/01.png"
	SourcePath02 = "test/source/02.png"
	SourcePath03 = "test/source/03.png"
	SourcePath04 = "test/source/04.png"
	SourcePath05 = "test/source/05.png"
	SourcePath06 = "test/source/06.png"
	SourcePath07 = "test/source/07.png"
	SourcePath08 = "test/source/08.jpg"
	SourcePath09 = "test/source/09.png"
)

const (
	BlurPath01 = "test/blur/01.png"
	BlurPath02 = "test/blur/02.png"
	BlurPath03 = "test/blur/03.png"
	BlurPath04 = "test/blur/04.png"
	BlurPath05 = "test/blur/05.png"
	BlurPath06 = "test/blur/06.png"
	BlurPath07 = "test/blur/07.png"
	BlurPath08 = "test/blur/08.jpg"
	BlurPath09 = "test/blur/09.jpg"
)

const (
	CVTPath01 = "test/cvt/01.png"
	CVTPath02 = "test/cvt/02.png"
	CVTPath03 = "test/cvt/03.png"
	CVTPath04 = "test/cvt/04.png"
	CVTPath05 = "test/cvt/05.png"
	CVTPath06 = "test/cvt/06.png"
	CVTPath07 = "test/cvt/07.png"
	CVTPath08 = "test/cvt/08.jpg"
	CVTPath09 = "test/cvt/09.png"
)

const (
	DilatePath01 = "test/dilate/01.png"
	DilatePath02 = "test/dilate/02.png"
	DilatePath03 = "test/dilate/03.png"
	DilatePath04 = "test/dilate/04.png"
	DilatePath05 = "test/dilate/05.png"
	DilatePath06 = "test/dilate/06.png"
	DilatePath07 = "test/dilate/07.png"
	DilatePath08 = "test/dilate/08.jpg"
	DilatePath09 = "test/dilate/09.png"
)

const (
	ErodePath01 = "test/erode/01.png"
	ErodePath02 = "test/erode/02.png"
	ErodePath03 = "test/erode/03.png"
	ErodePath04 = "test/erode/04.png"
	ErodePath05 = "test/erode/05.png"
	ErodePath06 = "test/erode/06.png"
	ErodePath07 = "test/erode/07.png"
	ErodePath08 = "test/erode/08.jpg"
	ErodePath09 = "test/erode/09.png"
)

const (
	GrayPath01 = "test/gray/01.png"
	GrayPath02 = "test/gray/02.png"
	GrayPath03 = "test/gray/03.png"
	GrayPath04 = "test/gray/04.png"
	GrayPath05 = "test/gray/05.png"
	GrayPath06 = "test/gray/06.png"
	GrayPath07 = "test/gray/07.png"
	GrayPath08 = "test/gray/08.jpg"
	GrayPath09 = "test/gray/09.png"
)

const (
	RGBPath01 = "test/rgb/01.png"
	RGBPath02 = "test/rgb/02.png"
	RGBPath03 = "test/rgb/03.png"
	RGBPath04 = "test/rgb/04.png"
	RGBPath05 = "test/rgb/05.png"
	RGBPath06 = "test/rgb/06.png"
	RGBPath07 = "test/rgb/07.png"
	RGBPath08 = "test/rgb/08.jpg"
	RGBPath09 = "test/rgb/09.png"
)

const (
	SharpenPath01 = "test/sharpen/01.png"
	SharpenPath02 = "test/sharpen/02.png"
	SharpenPath03 = "test/sharpen/03.png"
	SharpenPath04 = "test/sharpen/04.png"
	SharpenPath05 = "test/sharpen/05.png"
	SharpenPath06 = "test/sharpen/06.png"
	SharpenPath07 = "test/sharpen/07.png"
	SharpenPath08 = "test/sharpen/08.jpg"
	SharpenPath09 = "test/sharpen/09.png"
)

const (
	EdgePath01 = "test/edge/01.png"
	EdgePath02 = "test/edge/02.png"
	EdgePath03 = "test/edge/03.png"
	EdgePath04 = "test/edge/04.png"
	EdgePath05 = "test/edge/05.png"
	EdgePath06 = "test/edge/06.png"
	EdgePath07 = "test/edge/07.png"
	EdgePath08 = "test/edge/08.jpg"
	EdgePath09 = "test/edge/09.png"
)

const (
	EmbossPath01 = "test/emboss/01.png"
	EmbossPath02 = "test/emboss/02.png"
	EmbossPath03 = "test/emboss/03.png"
	EmbossPath04 = "test/emboss/04.png"
	EmbossPath05 = "test/emboss/05.png"
	EmbossPath06 = "test/emboss/06.png"
	EmbossPath07 = "test/emboss/07.png"
	EmbossPath08 = "test/emboss/08.jpg"
	EmbossPath09 = "test/emboss/09.png"
)

var (
	SourcePaths  = []string{SourcePath01, SourcePath02, SourcePath03, SourcePath04, SourcePath05, SourcePath06, SourcePath07, SourcePath08, SourcePath09}
	BlurPaths    = []string{BlurPath01, BlurPath02, BlurPath03, BlurPath04, BlurPath05, BlurPath06, BlurPath07, BlurPath08, BlurPath09}
	CVTPaths     = []string{CVTPath01, CVTPath02, CVTPath03, CVTPath04, CVTPath05, CVTPath06, CVTPath07, CVTPath08, CVTPath09}
	DilatePaths  = []string{DilatePath01, DilatePath02, DilatePath03, DilatePath04, DilatePath05, DilatePath06, DilatePath07, DilatePath08, DilatePath09}
	ErodePaths   = []string{ErodePath01, ErodePath02, ErodePath03, ErodePath04, ErodePath05, ErodePath06, ErodePath07, ErodePath08, ErodePath09}
	GrayPaths    = []string{GrayPath01, GrayPath02, GrayPath03, GrayPath04, GrayPath05, GrayPath06, GrayPath07, GrayPath08, GrayPath09}
	RGBPaths     = []string{RGBPath01, RGBPath02, RGBPath03, RGBPath04, RGBPath05, RGBPath06, RGBPath07, RGBPath08, RGBPath09}
	SharpenPaths = []string{SharpenPath01, SharpenPath02, SharpenPath03, SharpenPath04, SharpenPath05, SharpenPath06, SharpenPath07, SharpenPath08, SharpenPath09}
	EdgePaths    = []string{EdgePath01, EdgePath02, EdgePath03, EdgePath04, EdgePath05, EdgePath06, EdgePath07, EdgePath08, EdgePath09}
	EmbossPaths  = []string{EmbossPath01, EmbossPath02, EmbossPath03, EmbossPath04, EmbossPath05, EmbossPath06, EmbossPath07, EmbossPath08, EmbossPath09}
)
