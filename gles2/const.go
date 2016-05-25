// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 26 May 2016 00:39:07 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package gles2

/*
#cgo LDFLAGS: -lGLESv2
#include <GLES2/gl2.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// EsVersion20 as defined in GLES2/gl2.h:54
	EsVersion20 = 1
	// DepthBufferBit as defined in GLES2/gl2.h:76
	DepthBufferBit = 0x00000100
	// StencilBufferBit as defined in GLES2/gl2.h:77
	StencilBufferBit = 0x00000400
	// ColorBufferBit as defined in GLES2/gl2.h:78
	ColorBufferBit = 0x00004000
	// False as defined in GLES2/gl2.h:79
	False = 0
	// True as defined in GLES2/gl2.h:80
	True = 1
	// Points as defined in GLES2/gl2.h:81
	Points = 0x0000
	// Lines as defined in GLES2/gl2.h:82
	Lines = 0x0001
	// LineLoop as defined in GLES2/gl2.h:83
	LineLoop = 0x0002
	// LineStrip as defined in GLES2/gl2.h:84
	LineStrip = 0x0003
	// Triangles as defined in GLES2/gl2.h:85
	Triangles = 0x0004
	// TriangleStrip as defined in GLES2/gl2.h:86
	TriangleStrip = 0x0005
	// TriangleFan as defined in GLES2/gl2.h:87
	TriangleFan = 0x0006
	// Zero as defined in GLES2/gl2.h:88
	Zero = 0
	// One as defined in GLES2/gl2.h:89
	One = 1
	// SrcColor as defined in GLES2/gl2.h:90
	SrcColor = 0x0300
	// OneMinusSrcColor as defined in GLES2/gl2.h:91
	OneMinusSrcColor = 0x0301
	// SrcAlpha as defined in GLES2/gl2.h:92
	SrcAlpha = 0x0302
	// OneMinusSrcAlpha as defined in GLES2/gl2.h:93
	OneMinusSrcAlpha = 0x0303
	// DstAlpha as defined in GLES2/gl2.h:94
	DstAlpha = 0x0304
	// OneMinusDstAlpha as defined in GLES2/gl2.h:95
	OneMinusDstAlpha = 0x0305
	// DstColor as defined in GLES2/gl2.h:96
	DstColor = 0x0306
	// OneMinusDstColor as defined in GLES2/gl2.h:97
	OneMinusDstColor = 0x0307
	// SrcAlphaSaturate as defined in GLES2/gl2.h:98
	SrcAlphaSaturate = 0x0308
	// FuncAdd as defined in GLES2/gl2.h:99
	FuncAdd = 0x8006
	// BlendEquation as defined in GLES2/gl2.h:100
	BlendEquation = 0x8009
	// BlendEquationRgb as defined in GLES2/gl2.h:101
	BlendEquationRgb = 0x8009
	// BlendEquationAlpha as defined in GLES2/gl2.h:102
	BlendEquationAlpha = 0x883D
	// FuncSubtract as defined in GLES2/gl2.h:103
	FuncSubtract = 0x800A
	// FuncReverseSubtract as defined in GLES2/gl2.h:104
	FuncReverseSubtract = 0x800B
	// BlendDstRgb as defined in GLES2/gl2.h:105
	BlendDstRgb = 0x80C8
	// BlendSrcRgb as defined in GLES2/gl2.h:106
	BlendSrcRgb = 0x80C9
	// BlendDstAlpha as defined in GLES2/gl2.h:107
	BlendDstAlpha = 0x80CA
	// BlendSrcAlpha as defined in GLES2/gl2.h:108
	BlendSrcAlpha = 0x80CB
	// ConstantColor as defined in GLES2/gl2.h:109
	ConstantColor = 0x8001
	// OneMinusConstantColor as defined in GLES2/gl2.h:110
	OneMinusConstantColor = 0x8002
	// ConstantAlpha as defined in GLES2/gl2.h:111
	ConstantAlpha = 0x8003
	// OneMinusConstantAlpha as defined in GLES2/gl2.h:112
	OneMinusConstantAlpha = 0x8004
	// BlendColor as defined in GLES2/gl2.h:113
	BlendColor = 0x8005
	// ArrayBuffer as defined in GLES2/gl2.h:114
	ArrayBuffer = 0x8892
	// ElementArrayBuffer as defined in GLES2/gl2.h:115
	ElementArrayBuffer = 0x8893
	// ArrayBufferBinding as defined in GLES2/gl2.h:116
	ArrayBufferBinding = 0x8894
	// ElementArrayBufferBinding as defined in GLES2/gl2.h:117
	ElementArrayBufferBinding = 0x8895
	// StreamDraw as defined in GLES2/gl2.h:118
	StreamDraw = 0x88E0
	// StaticDraw as defined in GLES2/gl2.h:119
	StaticDraw = 0x88E4
	// DynamicDraw as defined in GLES2/gl2.h:120
	DynamicDraw = 0x88E8
	// BufferSize as defined in GLES2/gl2.h:121
	BufferSize = 0x8764
	// BufferUsage as defined in GLES2/gl2.h:122
	BufferUsage = 0x8765
	// CurrentVertexAttrib as defined in GLES2/gl2.h:123
	CurrentVertexAttrib = 0x8626
	// Front as defined in GLES2/gl2.h:124
	Front = 0x0404
	// Back as defined in GLES2/gl2.h:125
	Back = 0x0405
	// FrontAndBack as defined in GLES2/gl2.h:126
	FrontAndBack = 0x0408
	// Texture2d as defined in GLES2/gl2.h:127
	Texture2d = 0x0DE1
	// CullFace as defined in GLES2/gl2.h:128
	CullFace = 0x0B44
	// Blend as defined in GLES2/gl2.h:129
	Blend = 0x0BE2
	// Dither as defined in GLES2/gl2.h:130
	Dither = 0x0BD0
	// StencilTest as defined in GLES2/gl2.h:131
	StencilTest = 0x0B90
	// DepthTest as defined in GLES2/gl2.h:132
	DepthTest = 0x0B71
	// ScissorTest as defined in GLES2/gl2.h:133
	ScissorTest = 0x0C11
	// PolygonOffsetFill as defined in GLES2/gl2.h:134
	PolygonOffsetFill = 0x8037
	// SampleAlphaToCoverage as defined in GLES2/gl2.h:135
	SampleAlphaToCoverage = 0x809E
	// SampleCoverage as defined in GLES2/gl2.h:136
	SampleCoverage = 0x80A0
	// NoError as defined in GLES2/gl2.h:137
	NoError = 0
	// InvalidEnum as defined in GLES2/gl2.h:138
	InvalidEnum = 0x0500
	// InvalidValue as defined in GLES2/gl2.h:139
	InvalidValue = 0x0501
	// InvalidOperation as defined in GLES2/gl2.h:140
	InvalidOperation = 0x0502
	// OutOfMemory as defined in GLES2/gl2.h:141
	OutOfMemory = 0x0505
	// Cw as defined in GLES2/gl2.h:142
	Cw = 0x0900
	// Ccw as defined in GLES2/gl2.h:143
	Ccw = 0x0901
	// LineWidth as defined in GLES2/gl2.h:144
	LineWidth = 0x0B21
	// AliasedPointSizeRange as defined in GLES2/gl2.h:145
	AliasedPointSizeRange = 0x846D
	// AliasedLineWidthRange as defined in GLES2/gl2.h:146
	AliasedLineWidthRange = 0x846E
	// CullFaceMode as defined in GLES2/gl2.h:147
	CullFaceMode = 0x0B45
	// FrontFace as defined in GLES2/gl2.h:148
	FrontFace = 0x0B46
	// DepthRange as defined in GLES2/gl2.h:149
	DepthRange = 0x0B70
	// DepthWritemask as defined in GLES2/gl2.h:150
	DepthWritemask = 0x0B72
	// DepthClearValue as defined in GLES2/gl2.h:151
	DepthClearValue = 0x0B73
	// DepthFunc as defined in GLES2/gl2.h:152
	DepthFunc = 0x0B74
	// StencilClearValue as defined in GLES2/gl2.h:153
	StencilClearValue = 0x0B91
	// StencilFunc as defined in GLES2/gl2.h:154
	StencilFunc = 0x0B92
	// StencilFail as defined in GLES2/gl2.h:155
	StencilFail = 0x0B94
	// StencilPassDepthFail as defined in GLES2/gl2.h:156
	StencilPassDepthFail = 0x0B95
	// StencilPassDepthPass as defined in GLES2/gl2.h:157
	StencilPassDepthPass = 0x0B96
	// StencilRef as defined in GLES2/gl2.h:158
	StencilRef = 0x0B97
	// StencilValueMask as defined in GLES2/gl2.h:159
	StencilValueMask = 0x0B93
	// StencilWritemask as defined in GLES2/gl2.h:160
	StencilWritemask = 0x0B98
	// StencilBackFunc as defined in GLES2/gl2.h:161
	StencilBackFunc = 0x8800
	// StencilBackFail as defined in GLES2/gl2.h:162
	StencilBackFail = 0x8801
	// StencilBackPassDepthFail as defined in GLES2/gl2.h:163
	StencilBackPassDepthFail = 0x8802
	// StencilBackPassDepthPass as defined in GLES2/gl2.h:164
	StencilBackPassDepthPass = 0x8803
	// StencilBackRef as defined in GLES2/gl2.h:165
	StencilBackRef = 0x8CA3
	// StencilBackValueMask as defined in GLES2/gl2.h:166
	StencilBackValueMask = 0x8CA4
	// StencilBackWritemask as defined in GLES2/gl2.h:167
	StencilBackWritemask = 0x8CA5
	// Viewport as defined in GLES2/gl2.h:168
	Viewport = 0x0BA2
	// ScissorBox as defined in GLES2/gl2.h:169
	ScissorBox = 0x0C10
	// ColorClearValue as defined in GLES2/gl2.h:170
	ColorClearValue = 0x0C22
	// ColorWritemask as defined in GLES2/gl2.h:171
	ColorWritemask = 0x0C23
	// UnpackAlignment as defined in GLES2/gl2.h:172
	UnpackAlignment = 0x0CF5
	// PackAlignment as defined in GLES2/gl2.h:173
	PackAlignment = 0x0D05
	// MaxTextureSize as defined in GLES2/gl2.h:174
	MaxTextureSize = 0x0D33
	// MaxViewportDims as defined in GLES2/gl2.h:175
	MaxViewportDims = 0x0D3A
	// SubpixelBits as defined in GLES2/gl2.h:176
	SubpixelBits = 0x0D50
	// RedBits as defined in GLES2/gl2.h:177
	RedBits = 0x0D52
	// GreenBits as defined in GLES2/gl2.h:178
	GreenBits = 0x0D53
	// BlueBits as defined in GLES2/gl2.h:179
	BlueBits = 0x0D54
	// AlphaBits as defined in GLES2/gl2.h:180
	AlphaBits = 0x0D55
	// DepthBits as defined in GLES2/gl2.h:181
	DepthBits = 0x0D56
	// StencilBits as defined in GLES2/gl2.h:182
	StencilBits = 0x0D57
	// PolygonOffsetUnits as defined in GLES2/gl2.h:183
	PolygonOffsetUnits = 0x2A00
	// PolygonOffsetFactor as defined in GLES2/gl2.h:184
	PolygonOffsetFactor = 0x8038
	// TextureBinding2d as defined in GLES2/gl2.h:185
	TextureBinding2d = 0x8069
	// SampleBuffers as defined in GLES2/gl2.h:186
	SampleBuffers = 0x80A8
	// Samples as defined in GLES2/gl2.h:187
	Samples = 0x80A9
	// SampleCoverageValue as defined in GLES2/gl2.h:188
	SampleCoverageValue = 0x80AA
	// SampleCoverageInvert as defined in GLES2/gl2.h:189
	SampleCoverageInvert = 0x80AB
	// NumCompressedTextureFormats as defined in GLES2/gl2.h:190
	NumCompressedTextureFormats = 0x86A2
	// CompressedTextureFormats as defined in GLES2/gl2.h:191
	CompressedTextureFormats = 0x86A3
	// DontCare as defined in GLES2/gl2.h:192
	DontCare = 0x1100
	// Fastest as defined in GLES2/gl2.h:193
	Fastest = 0x1101
	// Nicest as defined in GLES2/gl2.h:194
	Nicest = 0x1102
	// GenerateMipmapHint as defined in GLES2/gl2.h:195
	GenerateMipmapHint = 0x8192
	// Byte as defined in GLES2/gl2.h:196
	Byte = 0x1400
	// UnsignedByte as defined in GLES2/gl2.h:197
	UnsignedByte = 0x1401
	// Short as defined in GLES2/gl2.h:198
	Short = 0x1402
	// UnsignedShort as defined in GLES2/gl2.h:199
	UnsignedShort = 0x1403
	// Int as defined in GLES2/gl2.h:200
	Int = 0x1404
	// UnsignedInt as defined in GLES2/gl2.h:201
	UnsignedInt = 0x1405
	// Float as defined in GLES2/gl2.h:202
	Float = 0x1406
	// Fixed as defined in GLES2/gl2.h:203
	Fixed = 0x140C
	// DepthComponent as defined in GLES2/gl2.h:204
	DepthComponent = 0x1902
	// Alpha as defined in GLES2/gl2.h:205
	Alpha = 0x1906
	// Rgb as defined in GLES2/gl2.h:206
	Rgb = 0x1907
	// Rgba as defined in GLES2/gl2.h:207
	Rgba = 0x1908
	// Luminance as defined in GLES2/gl2.h:208
	Luminance = 0x1909
	// LuminanceAlpha as defined in GLES2/gl2.h:209
	LuminanceAlpha = 0x190A
	// UnsignedShort4444 as defined in GLES2/gl2.h:210
	UnsignedShort4444 = 0x8033
	// UnsignedShort5551 as defined in GLES2/gl2.h:211
	UnsignedShort5551 = 0x8034
	// UnsignedShort565 as defined in GLES2/gl2.h:212
	UnsignedShort565 = 0x8363
	// FragmentShader as defined in GLES2/gl2.h:213
	FragmentShader = 0x8B30
	// VertexShader as defined in GLES2/gl2.h:214
	VertexShader = 0x8B31
	// MaxVertexAttribs as defined in GLES2/gl2.h:215
	MaxVertexAttribs = 0x8869
	// MaxVertexUniformVectors as defined in GLES2/gl2.h:216
	MaxVertexUniformVectors = 0x8DFB
	// MaxVaryingVectors as defined in GLES2/gl2.h:217
	MaxVaryingVectors = 0x8DFC
	// MaxCombinedTextureImageUnits as defined in GLES2/gl2.h:218
	MaxCombinedTextureImageUnits = 0x8B4D
	// MaxVertexTextureImageUnits as defined in GLES2/gl2.h:219
	MaxVertexTextureImageUnits = 0x8B4C
	// MaxTextureImageUnits as defined in GLES2/gl2.h:220
	MaxTextureImageUnits = 0x8872
	// MaxFragmentUniformVectors as defined in GLES2/gl2.h:221
	MaxFragmentUniformVectors = 0x8DFD
	// ShaderType as defined in GLES2/gl2.h:222
	ShaderType = 0x8B4F
	// DeleteStatus as defined in GLES2/gl2.h:223
	DeleteStatus = 0x8B80
	// LinkStatus as defined in GLES2/gl2.h:224
	LinkStatus = 0x8B82
	// ValidateStatus as defined in GLES2/gl2.h:225
	ValidateStatus = 0x8B83
	// AttachedShaders as defined in GLES2/gl2.h:226
	AttachedShaders = 0x8B85
	// ActiveUniforms as defined in GLES2/gl2.h:227
	ActiveUniforms = 0x8B86
	// ActiveUniformMaxLength as defined in GLES2/gl2.h:228
	ActiveUniformMaxLength = 0x8B87
	// ActiveAttributes as defined in GLES2/gl2.h:229
	ActiveAttributes = 0x8B89
	// ActiveAttributeMaxLength as defined in GLES2/gl2.h:230
	ActiveAttributeMaxLength = 0x8B8A
	// ShadingLanguageVersion as defined in GLES2/gl2.h:231
	ShadingLanguageVersion = 0x8B8C
	// CurrentProgram as defined in GLES2/gl2.h:232
	CurrentProgram = 0x8B8D
	// Never as defined in GLES2/gl2.h:233
	Never = 0x0200
	// Less as defined in GLES2/gl2.h:234
	Less = 0x0201
	// Equal as defined in GLES2/gl2.h:235
	Equal = 0x0202
	// Lequal as defined in GLES2/gl2.h:236
	Lequal = 0x0203
	// Greater as defined in GLES2/gl2.h:237
	Greater = 0x0204
	// Notequal as defined in GLES2/gl2.h:238
	Notequal = 0x0205
	// Gequal as defined in GLES2/gl2.h:239
	Gequal = 0x0206
	// Always as defined in GLES2/gl2.h:240
	Always = 0x0207
	// Keep as defined in GLES2/gl2.h:241
	Keep = 0x1E00
	// Replace as defined in GLES2/gl2.h:242
	Replace = 0x1E01
	// Incr as defined in GLES2/gl2.h:243
	Incr = 0x1E02
	// Decr as defined in GLES2/gl2.h:244
	Decr = 0x1E03
	// Invert as defined in GLES2/gl2.h:245
	Invert = 0x150A
	// IncrWrap as defined in GLES2/gl2.h:246
	IncrWrap = 0x8507
	// DecrWrap as defined in GLES2/gl2.h:247
	DecrWrap = 0x8508
	// Vendor as defined in GLES2/gl2.h:248
	Vendor = 0x1F00
	// Renderer as defined in GLES2/gl2.h:249
	Renderer = 0x1F01
	// Version as defined in GLES2/gl2.h:250
	Version = 0x1F02
	// Extensions as defined in GLES2/gl2.h:251
	Extensions = 0x1F03
	// Nearest as defined in GLES2/gl2.h:252
	Nearest = 0x2600
	// Linear as defined in GLES2/gl2.h:253
	Linear = 0x2601
	// NearestMipmapNearest as defined in GLES2/gl2.h:254
	NearestMipmapNearest = 0x2700
	// LinearMipmapNearest as defined in GLES2/gl2.h:255
	LinearMipmapNearest = 0x2701
	// NearestMipmapLinear as defined in GLES2/gl2.h:256
	NearestMipmapLinear = 0x2702
	// LinearMipmapLinear as defined in GLES2/gl2.h:257
	LinearMipmapLinear = 0x2703
	// TextureMagFilter as defined in GLES2/gl2.h:258
	TextureMagFilter = 0x2800
	// TextureMinFilter as defined in GLES2/gl2.h:259
	TextureMinFilter = 0x2801
	// TextureWrapS as defined in GLES2/gl2.h:260
	TextureWrapS = 0x2802
	// TextureWrapT as defined in GLES2/gl2.h:261
	TextureWrapT = 0x2803
	// Texture as defined in GLES2/gl2.h:262
	Texture = 0x1702
	// TextureCubeMap as defined in GLES2/gl2.h:263
	TextureCubeMap = 0x8513
	// TextureBindingCubeMap as defined in GLES2/gl2.h:264
	TextureBindingCubeMap = 0x8514
	// TextureCubeMapPositiveX as defined in GLES2/gl2.h:265
	TextureCubeMapPositiveX = 0x8515
	// TextureCubeMapNegativeX as defined in GLES2/gl2.h:266
	TextureCubeMapNegativeX = 0x8516
	// TextureCubeMapPositiveY as defined in GLES2/gl2.h:267
	TextureCubeMapPositiveY = 0x8517
	// TextureCubeMapNegativeY as defined in GLES2/gl2.h:268
	TextureCubeMapNegativeY = 0x8518
	// TextureCubeMapPositiveZ as defined in GLES2/gl2.h:269
	TextureCubeMapPositiveZ = 0x8519
	// TextureCubeMapNegativeZ as defined in GLES2/gl2.h:270
	TextureCubeMapNegativeZ = 0x851A
	// MaxCubeMapTextureSize as defined in GLES2/gl2.h:271
	MaxCubeMapTextureSize = 0x851C
	// Texture0 as defined in GLES2/gl2.h:272
	Texture0 = 0x84C0
	// Texture1 as defined in GLES2/gl2.h:273
	Texture1 = 0x84C1
	// Texture2 as defined in GLES2/gl2.h:274
	Texture2 = 0x84C2
	// Texture3 as defined in GLES2/gl2.h:275
	Texture3 = 0x84C3
	// Texture4 as defined in GLES2/gl2.h:276
	Texture4 = 0x84C4
	// Texture5 as defined in GLES2/gl2.h:277
	Texture5 = 0x84C5
	// Texture6 as defined in GLES2/gl2.h:278
	Texture6 = 0x84C6
	// Texture7 as defined in GLES2/gl2.h:279
	Texture7 = 0x84C7
	// Texture8 as defined in GLES2/gl2.h:280
	Texture8 = 0x84C8
	// Texture9 as defined in GLES2/gl2.h:281
	Texture9 = 0x84C9
	// Texture10 as defined in GLES2/gl2.h:282
	Texture10 = 0x84CA
	// Texture11 as defined in GLES2/gl2.h:283
	Texture11 = 0x84CB
	// Texture12 as defined in GLES2/gl2.h:284
	Texture12 = 0x84CC
	// Texture13 as defined in GLES2/gl2.h:285
	Texture13 = 0x84CD
	// Texture14 as defined in GLES2/gl2.h:286
	Texture14 = 0x84CE
	// Texture15 as defined in GLES2/gl2.h:287
	Texture15 = 0x84CF
	// Texture16 as defined in GLES2/gl2.h:288
	Texture16 = 0x84D0
	// Texture17 as defined in GLES2/gl2.h:289
	Texture17 = 0x84D1
	// Texture18 as defined in GLES2/gl2.h:290
	Texture18 = 0x84D2
	// Texture19 as defined in GLES2/gl2.h:291
	Texture19 = 0x84D3
	// Texture20 as defined in GLES2/gl2.h:292
	Texture20 = 0x84D4
	// Texture21 as defined in GLES2/gl2.h:293
	Texture21 = 0x84D5
	// Texture22 as defined in GLES2/gl2.h:294
	Texture22 = 0x84D6
	// Texture23 as defined in GLES2/gl2.h:295
	Texture23 = 0x84D7
	// Texture24 as defined in GLES2/gl2.h:296
	Texture24 = 0x84D8
	// Texture25 as defined in GLES2/gl2.h:297
	Texture25 = 0x84D9
	// Texture26 as defined in GLES2/gl2.h:298
	Texture26 = 0x84DA
	// Texture27 as defined in GLES2/gl2.h:299
	Texture27 = 0x84DB
	// Texture28 as defined in GLES2/gl2.h:300
	Texture28 = 0x84DC
	// Texture29 as defined in GLES2/gl2.h:301
	Texture29 = 0x84DD
	// Texture30 as defined in GLES2/gl2.h:302
	Texture30 = 0x84DE
	// Texture31 as defined in GLES2/gl2.h:303
	Texture31 = 0x84DF
	// ActiveTexture as defined in GLES2/gl2.h:304
	ActiveTexture = 0x84E0
	// Repeat as defined in GLES2/gl2.h:305
	Repeat = 0x2901
	// ClampToEdge as defined in GLES2/gl2.h:306
	ClampToEdge = 0x812F
	// MirroredRepeat as defined in GLES2/gl2.h:307
	MirroredRepeat = 0x8370
	// FloatVec2 as defined in GLES2/gl2.h:308
	FloatVec2 = 0x8B50
	// FloatVec3 as defined in GLES2/gl2.h:309
	FloatVec3 = 0x8B51
	// FloatVec4 as defined in GLES2/gl2.h:310
	FloatVec4 = 0x8B52
	// IntVec2 as defined in GLES2/gl2.h:311
	IntVec2 = 0x8B53
	// IntVec3 as defined in GLES2/gl2.h:312
	IntVec3 = 0x8B54
	// IntVec4 as defined in GLES2/gl2.h:313
	IntVec4 = 0x8B55
	// Bool as defined in GLES2/gl2.h:314
	Bool = 0x8B56
	// BoolVec2 as defined in GLES2/gl2.h:315
	BoolVec2 = 0x8B57
	// BoolVec3 as defined in GLES2/gl2.h:316
	BoolVec3 = 0x8B58
	// BoolVec4 as defined in GLES2/gl2.h:317
	BoolVec4 = 0x8B59
	// FloatMat2 as defined in GLES2/gl2.h:318
	FloatMat2 = 0x8B5A
	// FloatMat3 as defined in GLES2/gl2.h:319
	FloatMat3 = 0x8B5B
	// FloatMat4 as defined in GLES2/gl2.h:320
	FloatMat4 = 0x8B5C
	// Sampler2d as defined in GLES2/gl2.h:321
	Sampler2d = 0x8B5E
	// SamplerCube as defined in GLES2/gl2.h:322
	SamplerCube = 0x8B60
	// VertexAttribArrayEnabled as defined in GLES2/gl2.h:323
	VertexAttribArrayEnabled = 0x8622
	// VertexAttribArraySize as defined in GLES2/gl2.h:324
	VertexAttribArraySize = 0x8623
	// VertexAttribArrayStride as defined in GLES2/gl2.h:325
	VertexAttribArrayStride = 0x8624
	// VertexAttribArrayType as defined in GLES2/gl2.h:326
	VertexAttribArrayType = 0x8625
	// VertexAttribArrayNormalized as defined in GLES2/gl2.h:327
	VertexAttribArrayNormalized = 0x886A
	// VertexAttribArrayPointer as defined in GLES2/gl2.h:328
	VertexAttribArrayPointer = 0x8645
	// VertexAttribArrayBufferBinding as defined in GLES2/gl2.h:329
	VertexAttribArrayBufferBinding = 0x889F
	// ImplementationColorReadType as defined in GLES2/gl2.h:330
	ImplementationColorReadType = 0x8B9A
	// ImplementationColorReadFormat as defined in GLES2/gl2.h:331
	ImplementationColorReadFormat = 0x8B9B
	// CompileStatus as defined in GLES2/gl2.h:332
	CompileStatus = 0x8B81
	// InfoLogLength as defined in GLES2/gl2.h:333
	InfoLogLength = 0x8B84
	// ShaderSourceLength as defined in GLES2/gl2.h:334
	ShaderSourceLength = 0x8B88
	// ShaderCompiler as defined in GLES2/gl2.h:335
	ShaderCompiler = 0x8DFA
	// ShaderBinaryFormats as defined in GLES2/gl2.h:336
	ShaderBinaryFormats = 0x8DF8
	// NumShaderBinaryFormats as defined in GLES2/gl2.h:337
	NumShaderBinaryFormats = 0x8DF9
	// LowFloat as defined in GLES2/gl2.h:338
	LowFloat = 0x8DF0
	// MediumFloat as defined in GLES2/gl2.h:339
	MediumFloat = 0x8DF1
	// HighFloat as defined in GLES2/gl2.h:340
	HighFloat = 0x8DF2
	// LowInt as defined in GLES2/gl2.h:341
	LowInt = 0x8DF3
	// MediumInt as defined in GLES2/gl2.h:342
	MediumInt = 0x8DF4
	// HighInt as defined in GLES2/gl2.h:343
	HighInt = 0x8DF5
	// Framebuffer as defined in GLES2/gl2.h:344
	Framebuffer = 0x8D40
	// Renderbuffer as defined in GLES2/gl2.h:345
	Renderbuffer = 0x8D41
	// Rgba4 as defined in GLES2/gl2.h:346
	Rgba4 = 0x8056
	// Rgb5A1 as defined in GLES2/gl2.h:347
	Rgb5A1 = 0x8057
	// Rgb565 as defined in GLES2/gl2.h:348
	Rgb565 = 0x8D62
	// DepthComponent16 as defined in GLES2/gl2.h:349
	DepthComponent16 = 0x81A5
	// StencilIndex8 as defined in GLES2/gl2.h:350
	StencilIndex8 = 0x8D48
	// RenderbufferWidth as defined in GLES2/gl2.h:351
	RenderbufferWidth = 0x8D42
	// RenderbufferHeight as defined in GLES2/gl2.h:352
	RenderbufferHeight = 0x8D43
	// RenderbufferInternalFormat as defined in GLES2/gl2.h:353
	RenderbufferInternalFormat = 0x8D44
	// RenderbufferRedSize as defined in GLES2/gl2.h:354
	RenderbufferRedSize = 0x8D50
	// RenderbufferGreenSize as defined in GLES2/gl2.h:355
	RenderbufferGreenSize = 0x8D51
	// RenderbufferBlueSize as defined in GLES2/gl2.h:356
	RenderbufferBlueSize = 0x8D52
	// RenderbufferAlphaSize as defined in GLES2/gl2.h:357
	RenderbufferAlphaSize = 0x8D53
	// RenderbufferDepthSize as defined in GLES2/gl2.h:358
	RenderbufferDepthSize = 0x8D54
	// RenderbufferStencilSize as defined in GLES2/gl2.h:359
	RenderbufferStencilSize = 0x8D55
	// FramebufferAttachmentObjectType as defined in GLES2/gl2.h:360
	FramebufferAttachmentObjectType = 0x8CD0
	// FramebufferAttachmentObjectName as defined in GLES2/gl2.h:361
	FramebufferAttachmentObjectName = 0x8CD1
	// FramebufferAttachmentTextureLevel as defined in GLES2/gl2.h:362
	FramebufferAttachmentTextureLevel = 0x8CD2
	// FramebufferAttachmentTextureCubeMapFace as defined in GLES2/gl2.h:363
	FramebufferAttachmentTextureCubeMapFace = 0x8CD3
	// ColorAttachment0 as defined in GLES2/gl2.h:364
	ColorAttachment0 = 0x8CE0
	// DepthAttachment as defined in GLES2/gl2.h:365
	DepthAttachment = 0x8D00
	// StencilAttachment as defined in GLES2/gl2.h:366
	StencilAttachment = 0x8D20
	// None as defined in GLES2/gl2.h:367
	None = 0
	// FramebufferComplete as defined in GLES2/gl2.h:368
	FramebufferComplete = 0x8CD5
	// FramebufferIncompleteAttachment as defined in GLES2/gl2.h:369
	FramebufferIncompleteAttachment = 0x8CD6
	// FramebufferIncompleteMissingAttachment as defined in GLES2/gl2.h:370
	FramebufferIncompleteMissingAttachment = 0x8CD7
	// FramebufferIncompleteDimensions as defined in GLES2/gl2.h:371
	FramebufferIncompleteDimensions = 0x8CD9
	// FramebufferUnsupported as defined in GLES2/gl2.h:372
	FramebufferUnsupported = 0x8CDD
	// FramebufferBinding as defined in GLES2/gl2.h:373
	FramebufferBinding = 0x8CA6
	// RenderbufferBinding as defined in GLES2/gl2.h:374
	RenderbufferBinding = 0x8CA7
	// MaxRenderbufferSize as defined in GLES2/gl2.h:375
	MaxRenderbufferSize = 0x84E8
	// InvalidFramebufferOperation as defined in GLES2/gl2.h:376
	InvalidFramebufferOperation = 0x0506
)
