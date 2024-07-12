package size

type PaperSize struct {
	Name   string
	width  float64
	height float64
	Unit   *Unit
}

func (ps *PaperSize) Width() *Size {
	return &Size{
		ps.Unit,
		ps.width,
	}
}

func (ps *PaperSize) Height() *Size {
	return &Size{
		ps.Unit,
		ps.height,
	}
}

type Dimensions struct {
	X *Size
	Y *Size
}

const (
	Postcard     = "postcard"
	PosterSmall  = "poster-small"
	Poster       = "poster"
	PosterLarge  = "poster-large"
	BusinessCard = "business-card"
	R2           = "2r"
	R3           = "3r"
	R4           = "4r"
	R5           = "5r"
	R6           = "6r"
	R8           = "8r"
	R10          = "10r"
	R11          = "11r"
	R12          = "12r"
	A0           = "a0"
	A1           = "a1"
	A2           = "a2"
	A3           = "a3"
	A4           = "a4"
	A5           = "a5"
	A6           = "a6"
	A7           = "a7"
	A8           = "a8"
	A9           = "a9"
	A10          = "a10"
	TwoA0        = "2a0"
	FourA0       = "4a0"
	B0           = "b0"
	B1           = "b1"
	B1Plus       = "b1+"
	B2           = "b2"
	B2Plus       = "b2+"
	B3           = "b3"
	B4           = "b4"
	B5           = "b5"
	B6           = "b6"
	B7           = "b7"
	B8           = "b8"
	B9           = "b9"
	B10          = "b10"
	B11          = "b11"
	B12          = "b12"
	C0           = "c0"
	C1           = "c1"
	C2           = "c2"
	C3           = "c3"
	C4           = "c4"
	C5           = "c5"
	C6           = "c6"
	C7           = "c7"
	C8           = "c8"
	C9           = "c9"
	C10          = "c10"
	C11          = "c11"
	C12          = "c12"
	HalfLetter   = "half-letter"
	Letter       = "letter"
	Legal        = "legal"
	JuniorLegal  = "junior-legal"
	Ledger       = "ledger"
	Tabloid      = "tabloid"
	AnsiA        = "ansi-a"
	AnsiB        = "ansi-b"
	AnsiC        = "ansi-c"
	AnsiD        = "ansi-d"
	AnsiE        = "ansi-e"
	ArchA        = "arch-a"
	ArchB        = "arch-b"
	ArchC        = "arch-c"
	ArchD        = "arch-d"
	ArchE        = "arch-e"
	ArchE1       = "arch-e1"
	ArchE2       = "arch-e2"
	ArchE3       = "arch-e3"
)

var PaperSizes = map[string]PaperSize{
	// Common Paper Sizes
	// (Mostly North-American based)
	Postcard:     {Postcard, 101.6, 152.4, MM},
	PosterSmall:  {PosterSmall, 280, 430, MM},
	Poster:       {Poster, 460, 610, MM},
	PosterLarge:  {PosterLarge, 610, 910, MM},
	BusinessCard: {BusinessCard, 50.8, 88.9, MM},

	// Photographic Print Paper Sizes
	R2:  {R2, 64, 89, MM},
	R3:  {R3, 89, 127, MM},
	R4:  {R4, 102, 152, MM},
	R5:  {R5, 127, 178, MM},  // 5″x7″
	R6:  {R6, 152, 203, MM},  // 6″x8″
	R8:  {R8, 203, 254, MM},  // 8″x10″
	R10: {R10, 254, 305, MM}, // 10″x12″
	R11: {R11, 279, 356, MM}, // 11″x14″
	R12: {R12, 305, 381, MM},

	// Standard Paper Sizes
	A0:     {A0, 841, 1189, MM},
	A1:     {A1, 594, 841, MM},
	A2:     {A2, 420, 594, MM},
	A3:     {A3, 297, 420, MM},
	A4:     {A4, 210, 297, MM},
	A5:     {A5, 148, 210, MM},
	A6:     {A6, 105, 148, MM},
	A7:     {A7, 74, 105, MM},
	A8:     {A8, 52, 74, MM},
	A9:     {A9, 37, 52, MM},
	A10:    {A10, 26, 37, MM},
	TwoA0:  {TwoA0, 1189, 1682, MM},
	FourA0: {FourA0, 1682, 2378, MM},
	B0:     {B0, 1000, 1414, MM},
	B1:     {B1, 707, 1000, MM},
	B1Plus: {B1Plus, 720, 1020, MM},
	B2:     {B2, 500, 707, MM},
	B2Plus: {B2Plus, 520, 720, MM},
	B3:     {B3, 353, 500, MM},
	B4:     {B4, 250, 353, MM},
	B5:     {B5, 176, 250, MM},
	B6:     {B6, 125, 176, MM},
	B7:     {B7, 88, 125, MM},
	B8:     {B8, 62, 88, MM},
	B9:     {B9, 44, 62, MM},
	B10:    {B10, 31, 44, MM},
	B11:    {B11, 22, 32, MM},
	B12:    {B12, 16, 22, MM},
	C0:     {C0, 917, 1297, MM},
	C1:     {C1, 648, 917, MM},
	C2:     {C2, 458, 648, MM},
	C3:     {C3, 324, 458, MM},
	C4:     {C4, 229, 324, MM},
	C5:     {C5, 162, 229, MM},
	C6:     {C6, 114, 162, MM},
	C7:     {C7, 81, 114, MM},
	C8:     {C8, 57, 81, MM},
	C9:     {C9, 40, 57, MM},
	C10:    {C10, 28, 40, MM},
	C11:    {C11, 22, 32, MM},
	C12:    {C12, 16, 22, MM},

	// Use inches for North American sizes,
	// as it produces less float precision errors
	HalfLetter:  {HalfLetter, 5.5, 8.5, IN},
	Letter:      {Letter, 8.5, 11, IN},
	Legal:       {Legal, 8.5, 14, IN},
	JuniorLegal: {JuniorLegal, 5, 8, IN},
	Ledger:      {Ledger, 11, 17, IN},
	Tabloid:     {Tabloid, 11, 17, IN},
	AnsiA:       {AnsiA, 8.5, 11.0, IN},
	AnsiB:       {AnsiB, 11.0, 17.0, IN},
	AnsiC:       {AnsiC, 17.0, 22.0, IN},
	AnsiD:       {AnsiD, 22.0, 34.0, IN},
	AnsiE:       {AnsiE, 34.0, 44.0, IN},
	ArchA:       {ArchA, 9, 12, IN},
	ArchB:       {ArchB, 12, 18, IN},
	ArchC:       {ArchC, 18, 24, IN},
	ArchD:       {ArchD, 24, 36, IN},
	ArchE:       {ArchE, 36, 48, IN},
	ArchE1:      {ArchE1, 30, 42, IN},
	ArchE2:      {ArchE2, 26, 38, IN},
	ArchE3:      {ArchE3, 27, 39, IN},
}
