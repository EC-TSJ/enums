package enums_test

import (
	"enums"
	"fmt"
	"testing"
)

func TestFunctions(t *testing.T) {

	var z2 = enums.PoblacionFromProvincia(28055, 28)
	var z22 = enums.PoblacionFromProvincia("Estremera", 28)

	var z5 = enums.ProvinciasFromCCAA(8)
	var z51 = enums.ProvinciasFromCCAA("Castilla y León")
	var z52 = enums.ProvinciasFromCCAA("Comunidad Autónoma de Andalucía")

	var z6 = enums.ProvinciasFromDelegacion(8)
	var z61 = enums.ProvinciasFromDelegacion("Castilla y León")
	var z62 = enums.ProvinciasFromDelegacion("Comunidad Autónoma de Andalucía")

	var z7 = enums.ProvinciaFromPoblacion(28005)
	var z10 = enums.ProvinciaFromAdministracion(28005)
	var z8 = enums.CCAAFromDelegacion(28606)
	var z4 = enums.CCAAFromProvincia(51)
	var z41 = enums.CCAAFromProvincia(4)
	var z42 = enums.CCAAFromProvincia("Madrid")
	var z1 = enums.AdministracionFromProvincia(28609, 28)
	var z11 = enums.AdministracionFromProvincia("Madrid", 28)
	var z12 = enums.AdministracionesFromDelegacion(28)
	var z121 = enums.AdministracionesFromDelegacion("Madrid")
	var z122 = enums.AdministracionesFromDelegacion("Cataluña")
	var z123 = enums.AdministracionesFromDelegacion("Barcelona")
	var z9 = enums.AdministracionesFromCCAA(12)
	var z91 = enums.AdministracionesFromCCAA(1)
	var z92 = enums.AdministracionesFromCCAA("Castilla y León")
	var z222 = enums.DelegacionFromAdministracion(8095)
	var z223 = enums.DelegacionFromAdministracion(28616)
	var z13, z43 = enums.Banco(2046)
	var z14, z44 = enums.Banco("Caja Rural Montroy")
	var z3, zX = enums.Sigla("CL")
	z3, zX = enums.Sigla("Calle")

	zLL, cLL := enums.Poblaciones(28)
	zL1, cL1 := enums.Siglas()
	zA, cA := enums.Provincia("Madrid")
	zA, cA = enums.Provincia(28)
	zB, cB := enums.Banco("Banco Guipuzcoano")
	zB, cB = enums.Banco(42)
	zC, cC := enums.Pais("AUSTRALIA")
	zC, cC = enums.Pais(104)
	zD, cD := enums.Sigla("Carretera")
	zD, cD = enums.Sigla("CR")
	zE, cE := enums.Poblacion("Baños De Rio Tobia")
	zE, cE = enums.Poblacion(28026)
	zF, cF := enums.CCAA("Extremadura")
	zF, cF = enums.CCAA(1)
	zG, cG := enums.Administracion("Fuerteventura")
	zG, cG = enums.Administracion(35018)
	zH, cH := enums.Delegacion("Principado de Asturias")
	zH, cH = enums.Delegacion(33600)
	zH, cH = enums.Delegacion(8600)
	zJ, cJ := enums.Administraciones()

	fmt.Println(zA, cA, zB, cB, zC, cC, zD, cD, zE, cE, zF, cF, zG, cG, zH, cH, zJ, zLL, zL1, cLL, cL1, cJ)
	fmt.Println(z1, z11, z43, z44, zX, z2, z22, z8, z9, z91, z92, z10, z12, z121, z122, z123, z222, z223, z7, z3, z13, z14, z4, z41, z42, z5, z51, z52, z6, z61, z62)
}
