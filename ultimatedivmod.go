package piscine

func UltimateDivMod(a *int, b *int) {
	ab := *a

	*a = ab / *b

	*b = ab % *b
}
