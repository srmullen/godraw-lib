package size

type Dimensions struct {
	X Size
	Y Size
}

func (d Dimensions) Width() Size {
	return d.X
}

func (d Dimensions) Height() Size {
	return d.Y
}

type PaperSize struct {
	Name       string
	Dimensions Dimensions
	Unit       *Unit
}

func NewPaperSize(name string, width, height float64, unit *Unit) PaperSize {
	return PaperSize{
		Name: name,
		Dimensions: Dimensions{
			X: Size{
				Value: width,
				Unit:  unit,
			},
			Y: Size{
				Value: height,
				Unit:  unit,
			},
		},
		Unit: unit,
	}
}

func (ps PaperSize) Landscape() Dimensions {
	if ps.Dimensions.X.Value >= ps.Dimensions.Y.Value {
		return ps.Dimensions
	} else {
		return Dimensions{
			X: ps.Dimensions.Y,
			Y: ps.Dimensions.X,
		}
	}
}

func (ps PaperSize) Portrait() Dimensions {
	if ps.Dimensions.X.Value <= ps.Dimensions.Y.Value {
		return ps.Dimensions
	} else {
		return Dimensions{
			X: ps.Dimensions.Y,
			Y: ps.Dimensions.X,
		}
	}
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
	Postcard:     NewPaperSize(Postcard, 101.6, 152.4, MM),
	PosterSmall:  NewPaperSize(PosterSmall, 280, 430, MM),
	Poster:       NewPaperSize(Poster, 460, 610, MM),
	PosterLarge:  NewPaperSize(PosterLarge, 610, 910, MM),
	BusinessCard: NewPaperSize(BusinessCard, 50.8, 88.9, MM),

	// Photographic Print Paper Sizes
	R2:  NewPaperSize(R2, 64, 89, MM),
	R3:  NewPaperSize(R3, 89, 127, MM),
	R4:  NewPaperSize(R4, 102, 152, MM),
	R5:  NewPaperSize(R5, 127, 178, MM),  // 5″x7″
	R6:  NewPaperSize(R6, 152, 203, MM),  // 6″x8″
	R8:  NewPaperSize(R8, 203, 254, MM),  // 8″x10″
	R10: NewPaperSize(R10, 254, 305, MM), // 10″x12″
	R11: NewPaperSize(R11, 279, 356, MM), // 11″x14″
	R12: NewPaperSize(R12, 305, 381, MM),

	// Standard Paper Sizes
	A0:     NewPaperSize(A0, 841, 1189, MM),
	A1:     NewPaperSize(A1, 594, 841, MM),
	A2:     NewPaperSize(A2, 420, 594, MM),
	A3:     NewPaperSize(A3, 297, 420, MM),
	A4:     NewPaperSize(A4, 210, 297, MM),
	A5:     NewPaperSize(A5, 148, 210, MM),
	A6:     NewPaperSize(A6, 105, 148, MM),
	A7:     NewPaperSize(A7, 74, 105, MM),
	A8:     NewPaperSize(A8, 52, 74, MM),
	A9:     NewPaperSize(A9, 37, 52, MM),
	A10:    NewPaperSize(A10, 26, 37, MM),
	TwoA0:  NewPaperSize(TwoA0, 1189, 1682, MM),
	FourA0: NewPaperSize(FourA0, 1682, 2378, MM),
	B0:     NewPaperSize(B0, 1000, 1414, MM),
	B1:     NewPaperSize(B1, 707, 1000, MM),
	B1Plus: NewPaperSize(B1Plus, 720, 1020, MM),
	B2:     NewPaperSize(B2, 500, 707, MM),
	B2Plus: NewPaperSize(B2Plus, 520, 720, MM),
	B3:     NewPaperSize(B3, 353, 500, MM),
	B4:     NewPaperSize(B4, 250, 353, MM),
	B5:     NewPaperSize(B5, 176, 250, MM),
	B6:     NewPaperSize(B6, 125, 176, MM),
	B7:     NewPaperSize(B7, 88, 125, MM),
	B8:     NewPaperSize(B8, 62, 88, MM),
	B9:     NewPaperSize(B9, 44, 62, MM),
	B10:    NewPaperSize(B10, 31, 44, MM),
	B11:    NewPaperSize(B11, 22, 32, MM),
	B12:    NewPaperSize(B12, 16, 22, MM),
	C0:     NewPaperSize(C0, 917, 1297, MM),
	C1:     NewPaperSize(C1, 648, 917, MM),
	C2:     NewPaperSize(C2, 458, 648, MM),
	C3:     NewPaperSize(C3, 324, 458, MM),
	C4:     NewPaperSize(C4, 229, 324, MM),
	C5:     NewPaperSize(C5, 162, 229, MM),
	C6:     NewPaperSize(C6, 114, 162, MM),
	C7:     NewPaperSize(C7, 81, 114, MM),
	C8:     NewPaperSize(C8, 57, 81, MM),
	C9:     NewPaperSize(C9, 40, 57, MM),
	C10:    NewPaperSize(C10, 28, 40, MM),
	C11:    NewPaperSize(C11, 22, 32, MM),
	C12:    NewPaperSize(C12, 16, 22, MM),

	// Use inches for North American sizes,
	// as it produces less float precision errors
	HalfLetter:  NewPaperSize(HalfLetter, 5.5, 8.5, IN),
	Letter:      NewPaperSize(Letter, 8.5, 11, IN),
	Legal:       NewPaperSize(Legal, 8.5, 14, IN),
	JuniorLegal: NewPaperSize(JuniorLegal, 5, 8, IN),
	Ledger:      NewPaperSize(Ledger, 11, 17, IN),
	Tabloid:     NewPaperSize(Tabloid, 11, 17, IN),
	AnsiA:       NewPaperSize(AnsiA, 8.5, 11.0, IN),
	AnsiB:       NewPaperSize(AnsiB, 11.0, 17.0, IN),
	AnsiC:       NewPaperSize(AnsiC, 17.0, 22.0, IN),
	AnsiD:       NewPaperSize(AnsiD, 22.0, 34.0, IN),
	AnsiE:       NewPaperSize(AnsiE, 34.0, 44.0, IN),
	ArchA:       NewPaperSize(ArchA, 9, 12, IN),
	ArchB:       NewPaperSize(ArchB, 12, 18, IN),
	ArchC:       NewPaperSize(ArchC, 18, 24, IN),
	ArchD:       NewPaperSize(ArchD, 24, 36, IN),
	ArchE:       NewPaperSize(ArchE, 36, 48, IN),
	ArchE1:      NewPaperSize(ArchE1, 30, 42, IN),
	ArchE2:      NewPaperSize(ArchE2, 26, 38, IN),
	ArchE3:      NewPaperSize(ArchE3, 27, 39, IN),
}
