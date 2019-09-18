//
//Created by xuzhuoxi
//on 2019-05-26.
//@author xuzhuoxi
//
package blendx

import (
	"image/color"
	"math"
)

func init() {
	RegisterBlendFunc(SoftLight, BlendSoftLightColor, BlendSoftLightRGBA)
}

// 柔光模式
// 根据绘图色的明暗程度来决定最终色是变亮还是变暗，当绘图色比50%的灰要亮时，则底色图像变亮。当绘图色比50%的灰要暗时，则底色图像就变暗。如果绘图色有纯黑色或纯白色，最终色不是黑色或白色，而是稍微变暗或变亮。
// 如果底色是纯白色或纯黑色，不产生任何效果。此效果与发散的聚光灯照在图像上相似。
// (F<=0.5): R = 2*B*D + S*S*(1 - 2*D)
// (F>0.5): R = 2*B*(1 - D) + (2*D - 1)*√S
//
// (D<=128): R = B*F/128 + (255-2*F)*B*B/65025
// (D>128): R = B*(255-F)/128 + (2*F-255)*√(B/255)
//
// (F<=32768): R = B*F/32768 + (65535-2*F)*B*B/65025
// (F>32768): R = B*(65535-F)/32768 + (2*F-65535)*√(B/65535)
func BlendSoftLightColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendSoftLightRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 柔光模式
// 根据绘图色的明暗程度来决定最终色是变亮还是变暗，当绘图色比50%的灰要亮时，则底色图像变亮。当绘图色比50%的灰要暗时，则底色图像就变暗。如果绘图色有纯黑色或纯白色，最终色不是黑色或白色，而是稍微变暗或变亮。
// 如果底色是纯白色或纯黑色，不产生任何效果。此效果与发散的聚光灯照在图像上相似。
// (F<=0.5): R = 2*B*D + S*S*(1 - 2*D)
// (F>0.5): R = 2*B*(1 - D) + (2*D - 1)*√S
//
// (D<=128): R = B*F/128 + (255-2*F)*B*B/65025
// (D>128): R = B*(255-F)/128 + (2*F-255)*√(B/255)
//
// (F<=32768): R = B*F/32768 + (65535-2*F)*B*B/65025
// (F>32768): R = B*(65535-F)/32768 + (2*F-65535)*√(B/65535)
func BlendSoftLightRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = softLight(foreR, backR)
	G = softLight(foreG, backG)
	B = softLight(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = softLight(foreA, backA)
	}
	return
}

// (F<=0.5): R = 2*B*D + S*S*(1 - 2*D)
// (F>0.5): R = 2*B*(1 - D) + (2*D - 1)*√S
//
// (D<=128): R = B*F/128 + (255-2*F)*B*B/65025
// (D>128): R = B*(255-F)/128 + (2*F-255)*√(B/255)
//
// (F<=32768): R = B*F/32768 + (65535-2*F)*B*B/65025
// (F>32768): R = B*(65535-F)/32768 + (2*F-65535)*√(B/65535)
func softLight(F, B uint32) uint32 {
	if F <= 32768 {
		return B*F/32768 + (65535-2*F)*B*B/65025
	} else {
		return B*(65535-F)/32768 + (2*F-65535)*uint32(math.Sqrt(float64(B)/255))
	}
}
