/**
 * @fileoverview
 * @author Torres Sacristán, Jesús <ec-tsj@gmail.org>
 * @version 1.0
 * @copyright XCL System
 * @category System
 * @package System
 * @subpackage System.Base
 * @filesource
 */

package enums

import (
	"ec-tsj/core"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

//--------------------------------------------------------
// init()
//--------------------------------------------------------
var p string

func init() {
	p = os.Getenv("GOPATH")
	p += "/pkg/lib/enums/EProvincias.json"
}

//------------------------------------------------
//helpers
//------------------------------------------------
func _helperProvincias() []ProvinciaT {
	provincia := make([]ProvinciaT, 0)
	raw, _ := ioutil.ReadFile(p + "/EProvincias.json")
	json.Unmarshal(raw, &provincia)

	return provincia
}

func _helperPoblacion() []PoblacionT {
	provincia := make([]PoblacionT, 0)
	raw, _ := ioutil.ReadFile(p + "/EPoblaciones.json")
	json.Unmarshal(raw, &provincia)

	return provincia
}

func _helperBancos() []BancoT {
	provincia := make([]BancoT, 0)
	raw, _ := ioutil.ReadFile(p + "/EBancos.json")
	json.Unmarshal(raw, &provincia)

	return provincia
}

func _helperCCAA() []CCAA_T {
	provincia := make([]CCAA_T, 0)
	raw, _ := ioutil.ReadFile(p + "/EComunidadesAutonomas.json")
	json.Unmarshal(raw, &provincia)

	return provincia
}

func _helperSiglas() []SiglaT {
	provincia := make([]SiglaT, 0)
	raw, _ := ioutil.ReadFile(p + "/ESiglas.json")
	json.Unmarshal(raw, &provincia)

	return provincia
}

func _helperPaises() []PaisesT {
	provincia := make([]PaisesT, 0)
	raw, _ := ioutil.ReadFile(p + "/EPaises.json")
	json.Unmarshal(raw, &provincia)

	return provincia
}

//
//-----------------------------------------------
// tipos
//------------------------------------------------
type BancoT struct {
	Nombre string `json:"nombre"`
	Code   int    `json:"code"`
}

//
type PoblacionT struct {
	Nombre      string
	Code        string
	Poblaciones []PoblacionesT
}

type PoblacionesT struct {
	Nombre string
	Code   string
}

//
type ProvinciaT struct {
	Nombre     string  `json:"nombre"`
	Postnombre string  `json:"postnombre"`
	Code       []CodeT `json:"code"`
	CCAA       []CodeT `json:"ccaa"`
	AEAT       []AeatT `json:"aeat"`
}

//
type CodeT struct {
	Data string `json:"data"`
}

//
type AeatT struct {
	Nombre string `json:"nombre"`
	Code   int    `json:"code"`
}

//
type PaisesT struct {
	Cod2   string `json:"Cod2"`
	Cod3   string `json:"Cod3"`
	Code   string `json:"Code"`
	Nombre string `json:"Nombre"`
}

//
type CCAA_T struct {
	Code   string `json:"code"`
	Nombre string `json:"nombre"`
}

//
type SiglaT struct {
	Nombre string `json:"nombre"`
	Sigla  string `json:"sigla"`
}

//------------------------------------------
// Para sort.Sort(interface(elemento))
//------------------------------------------
type poblacionA []PoblacionesT

func (p poblacionA) Len() int           { return len(p) }
func (p poblacionA) Less(i, j int) bool { return p[i].Nombre < p[j].Nombre }
func (p poblacionA) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type aeatA []AeatT

func (f aeatA) Len() int           { return len(f) }
func (f aeatA) Less(i, j int) bool { return f[i].Nombre < f[j].Nombre }
func (f aeatA) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

//------------------------------------------
// error
//------------------------------------------
type EnumError struct {
	message string
}

func (e *EnumError) Error() string {
	return fmt.Sprintf("Error: %s.", e.message)
}

/**
 * Luego se hará:
 *    _, err := Function()
 *  y
 *		if aa, ok := err.(*EnumError) {
 *		  xxx
 *    }
 */

//--------------------------------------------
//funciones
//--------------------------------------------
/**
 * Obtiene un listado de las poblaciones
 *
 * @param {...int} hp (opcional) Número de provincia
 * @return {[]string, []PoblacionesT} listado de poblaciones
 */
func Poblaciones(hp ...int) ([]string, []PoblacionesT) {
	var sst []string
	var ssf []PoblacionesT

	poblacion := _helperPoblacion()
	if hp != nil {
		hp[0]--
		for _, ten := range poblacion[hp[0]].Poblaciones {
			sst = append(sst, ten.Nombre)
			ssf = append(ssf, ten)
		}
	} else {
		for _, ten := range poblacion {
			for _, val := range ten.Poblaciones {
				sst = append(sst, val.Nombre)
				ssf = append(ssf, val)
			}
		}
		sort.Strings(sst)
		sort.Sort(poblacionA(ssf))
	}

	return sst, ssf
}

/**
 * Obtiene un listado de los Bancos
 *
 * @return {[]string, []BancoT} listado de bancos
 */
func Bancos() ([]string, []BancoT) {
	var sst []string

	banco := _helperBancos()
	for _, ten := range banco {
		sst = append(sst, ten.Nombre)
	}

	return sst, banco
}

/**
 * Obtiene un listado de las provincias
 *
 * @return {[]string, []ProvinciaT} listado de provincias
 */
func Provincias() ([]string, []ProvinciaT) {
	var sst []string

	provincia := _helperProvincias()
	for _, ten := range provincia {
		sst = append(sst, ten.Nombre)
		if ten.Nombre == "Melilla" {
			break
		}
	}

	return sst, provincia
}

/**
 * Obtiene un listado de las comunidades autónomas
 *
 * @return {[]string, []CCAA_T} listado de comunidades
 */
func ComunidadesAutonomas() ([]string, []CCAA_T) {
	var s []string

	ccaa := _helperCCAA()
	for _, ten := range ccaa {
		s = append(s, ten.Nombre)
	}
	return s, ccaa
}

/**
 * Obtiene un listado de las delegaciones de Hacienda
 *
 * @return {[]string, []ProvinciaT} listado de delegaciones
 */
func Delegaciones() ([]string, []ProvinciaT) {
	var sst []string

	provincia := _helperProvincias()
	for _, ten := range provincia {
		if ten.Nombre != ten.Postnombre {
			sst = append(sst, ten.Postnombre)
		}
		if ten.Nombre == "Melilla" {
			break
		}
	}

	return sst, provincia
}

/**
 * Obtiene un listado de las delegaciones de Hacienda
 *
 * @return {[]string, []AeatT} listado de administraciones
 */
func Administraciones() ([]string, []AeatT) {
	var sst []string
	var ssf []AeatT

	provincia := _helperProvincias()
	for _, ten := range provincia {
		for _, val := range ten.AEAT {
			sst = append(sst, val.Nombre)
			ssf = append(ssf, val)
		}
	}
	sort.Strings(sst)
	sort.Sort(aeatA(ssf))

	return sst, ssf
}

/**
 * Obtiene un listado de Siglas
 *
 * @return {[]string, []SiglaT} Siglas
 */
func Siglas() ([]string, []SiglaT) {
	var s []string

	sigla := _helperSiglas()
	for _, ten := range sigla {
		s = append(s, ten.Nombre)
	}

	return s, sigla
}

/**
 * Obtiene un listado de paises
 *
 * @return {[]string, []PaisesT} paises
 */
func Paises() ([]string, []PaisesT) {
	var s []string

	pais := _helperPaises()
	for _, ten := range pais {
		s = append(s, ten.Nombre)
	}

	return s, pais
}

/**
 * Busca una Administración de Hacienda en cada provincia.
 *
 * se hace:
 *    <obj>.searchAEAT(valor[27], "Hortaleza");
 *
 *    // la tabla debe ser en la forma tabla[Numero de provincia]
 * ó
 *    <obj>.searchAEAT(valor[27], 28618);
 *
 * @param {core.T} name Nombre o Número de la Administración
 * @param {int} code Número de provincia.
 * @return {core.T} Nombre o Número de la administración
 */
func AdministracionFromProvincia(name core.T, code int) (core.T, AeatT) {
	var s core.T
	var f AeatT
	code--

	provincia := _helperProvincias()
	for v, _ := range provincia[code].AEAT {
		if provincia[code].AEAT[v].Nombre == name {
			s = provincia[code].AEAT[v].Code
			f = provincia[code].AEAT[v]
			break
		} else if provincia[code].AEAT[v].Code == name {
			s = provincia[code].AEAT[v].Nombre
			f = provincia[code].AEAT[v]
			break
		}
	}

	return s, f
}

/**
 * Busca la Población en cada provincia.
 *
 * se hace:
 *    <obj>.searchPobl(valor[27], " Estremera");
 *
 *    // la tabla debe ser en la forma tabla[Numero de provincia]
 * ó
 *    <obj>.searchPobl(valor[27], 28055);
 *
 * @param {core.T} name Nombre o Número de la Administración
 * @param {int} code Número de provincia (optional).
 * @return {core.T,PoblacionT} Nombre o Número de la administración
 */
func PoblacionFromProvincia(name core.T, code int) (core.T, PoblacionesT) {
	var sst core.T
	var b PoblacionesT
	code--

	poblatio := _helperPoblacion()
	switch name.(type) {
	case int:
		name = strconv.Itoa(name.(int))
	}

	for v, _ := range poblatio[code].Poblaciones {
		if poblatio[code].Poblaciones[v].Nombre == name {
			sst = poblatio[code].Poblaciones[v].Code
			b = poblatio[code].Poblaciones[v]
			break
		} else if poblatio[code].Poblaciones[v].Code == name {
			sst = poblatio[code].Poblaciones[v].Nombre
			b = poblatio[code].Poblaciones[v]
			break
		}
	}

	return sst, b
}

/**
 * Obtiene la comunidad autonoma de una delegacion
 *
 * @param {int} delegacion Número de delegación
 * @return {string, CCAA_T} Nombre de la comunidad autonoma
 */
func CCAAFromDelegacion(delegacion int) (string, CCAA_T) {
	var sst string
	var ssl int
	var delegacion_ string
	var ssf CCAA_T

	ccaa := _helperCCAA()
	provincia := _helperProvincias()

	delegacion_ = strconv.Itoa(delegacion)
	if len(delegacion_) == 4 {
		ssl, _ = strconv.Atoi(delegacion_[0:1])
	} else {
		ssl, _ = strconv.Atoi(delegacion_[0:2])
	}
	ssl--

	for v, _ := range provincia[ssl].AEAT {
		if provincia[ssl].AEAT[v].Code == delegacion {
			num, _ := strconv.Atoi(provincia[ssl].CCAA[0].Data)
			num--
			sst = ccaa[num].Nombre
			break
		}
	}

	for _, g := range ccaa {
		fg, _ := strconv.Atoi(g.Code)
		if fg == ssl {
			ssf = g
		}
	}
	return sst, ssf
}

/**
 * Obtiene uns lista de la Administraciones de Hacienda por CCAA
 *
 * @param {core.T} ccaa Nombre/Número de CCAA
 * @return Array con las administraciones
 */
func AdministracionesFromCCAA(ccaa core.T) ([]string, []AeatT) {
	var sst []string
	var sstT []AeatT
	var ssf string

	ccaaT := _helperCCAA()
	provincia := _helperProvincias()

	switch ccaa.(type) {
	case int:
		ssf = strconv.Itoa(ccaa.(int))
	case string:
		ssf = ccaa.(string)
	}

	if len(ssf) == 1 {
		ssf = "0" + ssf
	}

	for _, z := range ccaaT {
		if z.Code == ssf {
			ssf = z.Nombre
			break
		} else if z.Nombre == ssf {
			ssf = z.Code
			break
		}
	}

	for _, z := range provincia {
		if z.CCAA[1].Data == ssf || z.Postnombre == ssf {
			for _, por := range z.AEAT {
				sst = append(sst, por.Nombre)
				sstT = append(sstT, por)
			}
		}
	}

	return sst, sstT
}

/**
 * Obtiene el nombre de la provincia de la Administracion
 *
 * @param {int} administracion Numero de adminiatracion
 * @return {string, ProvinciaT} Nombre de la provincia
 */
func ProvinciaFromAdministracion(administracion int) (string, ProvinciaT) {
	provincia := _helperProvincias()
	administracion_ := strconv.Itoa(administracion)
	if len(administracion_) == 4 {
		ssl, _ := strconv.Atoi(administracion_[0:1])
		ssl--
		return provincia[ssl].Nombre, provincia[ssl]
	} else {
		ssl, _ := strconv.Atoi(administracion_[0:2])
		ssl--
		return provincia[ssl].Nombre, provincia[ssl]
	}
}

/**
 * Obtiene un listado de administraciones de delegación.
 * La delegación es la provincia
 *
 * @param {core.T} delegacion Número de delegación
 * @return {[]string, []AeatT } Array con administraciones de hacienda
 */
func AdministracionesFromDelegacion(delegacion core.T) ([]string, []AeatT) {
	var sst []string
	var ssl int
	var ssf string
	var ssfal ProvinciaT
	var ssfin []AeatT

	provincia := _helperProvincias()
	switch delegacion.(type) {
	case int:
		ssfal = provincia[delegacion.(int)-1]
	case string:
		ssf = delegacion.(string)
		for _, v := range provincia {
			if v.Nombre == ssf || v.Postnombre == ssf {
				ssl, _ = strconv.Atoi(v.Code[1].Data)
				ssl--
				break
			}
		}
		ssfal = provincia[ssl]
	}

	for f, _ := range ssfal.AEAT {
		sst = append(sst, ssfal.AEAT[f].Nombre)
		ssfin = append(ssfin, ssfal.AEAT[f])
	}

	return sst, ssfin
}

/**
 * Obtiene el nombre de la delegción de hacienda
 *
 * @param {int} administracion Número de administracion
 * @return {string, AEAT} Nombre de la Delegación
 */
func DelegacionFromAdministracion(administracion int) (string, ProvinciaT) {
	var ssl int
	var sst string

	provincia := _helperProvincias()
	administracion_ := strconv.Itoa(administracion)

	if len(administracion_) == 4 {
		sst = administracion_[0:1]
	} else {
		sst = administracion_[0:2]
	}

	if len(sst) == 1 {
		sst = "0" + sst
	}
	ssl, _ = strconv.Atoi(sst)
	ssl--

	return provincia[ssl].Postnombre, provincia[ssl]
}

/**
 * Obtiene la provincia de una población
 *
 * @param {int} poblacion Número de la población
 * @return {string, PoblacionT} Nombre de la provincia
 */
func ProvinciaFromPoblacion(poblacion int) (string, PoblacionT, error) {
	var pob string

	poblacionA := _helperPoblacion()
	poblacion_ := strconv.Itoa(poblacion)

	if len(poblacion_) == 4 {
		pob = poblacion_[0:1]
	} else {
		pob = poblacion_[0:2]
	}
	ssl, _ := strconv.Atoi(pob)
	ssl--

	for _, pobl := range poblacionA[ssl].Poblaciones {
		if pobl.Code == poblacion_ {
			return pobl.Nombre, poblacionA[ssl], nil
		}
	}

	return "", poblacionA[ssl], &EnumError{"No hay respuesta"}
}

/**
 * Obtiene las provincias de una Delegacion de Hacienda.
 * La Delegación es la CCAA.
 *
 * La diferencia con getProvinciasFromCCAA es que esta da una lista de
 * las provincias de la comunidad; y getProvinciasFromDelegacion
 * devuelve una lista de los centros de AEAT: Vigo, Gijón, Jerez de la
 * Frontera y Cartagena son centros de AEAT. Y Ceuta y Melilla pertenecen
 * a la delegación de Andalucía.
 * Aparte de la diferencia de nombres de delegación.
 *
 * @param {core.T} delegacion Nombre/Numero de Delegacion
 * @return []string, []ProvinciaT  Array de provincias
 */
func ProvinciasFromDelegacion(ccaa core.T) ([]string, []ProvinciaT) {
	var sabor string
	var sst []string
	var ssf []ProvinciaT

	provincia := _helperProvincias()
	ccaaD := _helperCCAA()

	switch ccaa.(type) {
	case string:
		for _, v := range ccaaD {
			if v.Nombre == ccaa {
				sabor = v.Code
				break
			}
		}
	case int:
		sabor = strconv.Itoa(ccaa.(int))
	}

	if len(sabor) == 1 {
		sabor = "0" + sabor
	}

	for _, v := range provincia {
		if v.CCAA[0].Data == sabor {
			sst = append(sst, v.Nombre)
			ssf = append(ssf, v)
		}
	}

	return sst, ssf
}

/**
 * Obtiene las provincias de una CCAA
 *
 * La diferencia entre getProvinciasFromCCAA es que esta da una lista de
 * las provincias de la comunidad; y getProvinciasFromDelegacion
 * devuelve una lista de los centros de AEAT: Vigo, Gijón, Jerez de la
 * Frontera y Cartagena son centros de AEAT. Y Ceuta y Melilla pertenecen
 * a la delegación de Andalucía.
 * Aparte de la diferencia de nombres de delegación.
 *
 * @param {core.T} ccaa Nombre/Numero de CCAA
 * @return {[]string, []ProvinciaT} Array de provincias
 */
func ProvinciasFromCCAA(ccaa core.T) ([]string, []ProvinciaT) {
	var sabor string
	var sst []string
	var ssf []ProvinciaT

	provincia := _helperProvincias()
	ccaaD := _helperCCAA()

	switch ccaa.(type) {
	case string:
		for _, v := range ccaaD {
			if v.Nombre == ccaa {
				sabor = v.Code
				break
			}
		}
	case int:
		sabor = strconv.Itoa(ccaa.(int))
	}

	if len(sabor) == 1 {
		sabor = "0" + sabor
	}

	for _, v := range provincia {
		if v.CCAA[0].Data == sabor {
			sst = append(sst, v.Nombre)
			ssf = append(ssf, v)
		}
	}

	return sst, ssf
}

/**
 * Obtiene una CCAA de una Provincia
 *
 * @param {core.T} province Número o nombre de provincia
 * @return {core.T, CCAA_T} El nombre o número de la CCAA
 */
func CCAAFromProvincia(province core.T) (core.T, CCAA_T, error) {
	var provincia_ int

	provincia := _helperProvincias()
	ccaa := _helperCCAA()

	switch province.(type) {
	case string:
		for _, v := range provincia {
			if v.Nombre == province {
				provincia_, _ = strconv.Atoi(v.Code[0].Data)
				break
			}
		}
	case int:
		provincia_ = province.(int)
	}
	provincia_--

	zzCod, _ := strconv.Atoi(provincia[provincia_].CCAA[0].Data)
	str := ccaa[zzCod-1].Nombre
	code := ccaa[zzCod-1].Code

	var sst string
	for _, v := range provincia {
		if v.CCAA[0].Data == code {
			return str, ccaa[zzCod-1], nil
		} else if v.Nombre == str {
			return code, ccaa[zzCod-1], nil
		}
	}

	return sst, ccaa[zzCod-1], &EnumError{"No hay respuesta"}
}

/**
 * Busca una sigla o su nombre
 *
 * @param {string} name Nombre ó Sigla
 * @return {string, SiglaT} nombre o sigla
 */
func Sigla(name string) (string, SiglaT, error) {
	var sst SiglaT

	sigla := _helperSiglas()
	for _, m := range sigla {
		if m.Nombre == name {
			return m.Nombre, m, nil
		} else if m.Sigla == name {
			return m.Nombre, m, nil
		}
	}

	return "", sst, &EnumError{"No hay respuesta"}
}

/**
 * Busca el nombre o el número del Pais
 *
 * @param {string} name Nombre ó Número
 * @return {string, PaisesT} nombre o sigla
 */
func Pais(name core.T) (string, PaisesT, error) {
	var sst string
	var ffg PaisesT
	var into int

	pais := _helperPaises()
	switch name.(type) {
	case int:
		into = name.(int)
		sst = strconv.Itoa(into)
		if len(sst) < 3 {
			sst = strings.Repeat("0", 3-len(sst)) + sst
		}
	case string:
		sst = name.(string)
		sst = strings.ToUpper(sst)
	}

	for _, m := range pais {
		if m.Nombre == sst {
			return m.Nombre, m, nil
		} else if m.Code == sst {
			return m.Nombre, m, nil
		}
	}

	return "", ffg, &EnumError{"No hay respuesta"}
}

/**
 * Obtiene un Banco
 *
 * @param {core.T} code Número o Nombre del banco
 * @return {string, BancoT} Nombre o Número del banco
 */
func Banco(code core.T) (string, BancoT, error) {
	var sst BancoT

	banco := _helperBancos()
	for _, v := range banco {
		if v.Code == code {
			return v.Nombre, v, nil
		}
		if v.Nombre == code {
			return v.Nombre, v, nil
		}
	}

	return "", sst, &EnumError{"No hay respuesta"}
}

/**
 * Obtiene una provincia
 *
 * @return {string, ProvinciaT} provincia
 */
func Provincia(pr core.T) (string, ProvinciaT) {
	var val int

	provincia := _helperProvincias()

	switch pr.(type) {
	case int:
		val = pr.(int)
		val--
	case string:
		for z, ten := range provincia {
			if ten.Nombre == pr {
				val = z
			}
		}
	}

	return provincia[val].Nombre, provincia[val]
}

/**
 * Obtiene una poblacion
 *
 * @param {...int} hp Nombre/Número de provincia
 * @return {string, PoblacionesT} listado de poblaciones
 */
func Poblacion(hp core.T) (string, PoblacionesT, error) {
	var s PoblacionesT
	var sd string

	poblacion := _helperPoblacion()
	switch hp.(type) {
	case int:
		sd = strconv.Itoa(hp.(int))
	}

	for _, ten := range poblacion {
		for _, val := range ten.Poblaciones {
			if val.Nombre == hp || val.Code == sd {
				return val.Nombre, val, nil
			}
		}
	}
	return "", s, &EnumError{"No hay respuesta"}
}

/**
 * Obtiene un listado de las comunidades autónomas
 *
 * @return {string, CCAA_T} listado de comunidades
 */
func CCAA(ca core.T) (string, CCAA_T, error) {
	var s CCAA_T
	var sd string

	ccaa := _helperCCAA()
	switch ca.(type) {
	case int:
		sd = strconv.Itoa(ca.(int))
		sd = strings.Repeat("0", 2-len(sd)) + sd
	}

	for _, ten := range ccaa {
		if ten.Nombre == ca || ten.Code == sd {
			return ten.Nombre, ten, nil
		}
	}

	return "", s, &EnumError{"No hay respuesta"}
}

/**
 * Obtiene una Administracion de hacienda
 *
 * @param {core.T} Nombre/Número
 * @return {string, Aeat} Administracion de Hacienda
 */
func Administracion(dg core.T) (string, AeatT, error) {
	var sst AeatT
	var sd int

	provincia := _helperProvincias()
	switch dg.(type) {
	case int:
		sd = dg.(int)
	}

	for _, ten := range provincia {
		for _, val := range ten.AEAT {
			if val.Nombre == dg || val.Code == sd {
				return val.Nombre, val, nil
			}
		}
	}

	return "", sst, &EnumError{"No hay respuesta"}
}

/**
 * Obtiene una Delegacion de hacienda
 *
 * @param {core.T} Nombre/Número
 * @return {string, ProvinciaT} Delegacion de Hacienda
 */
func Delegacion(dg core.T) (string, ProvinciaT, error) {
	var sst ProvinciaT
	var sds string

	provincia := _helperProvincias()
	switch dg.(type) {
	case int:
		sds = strconv.Itoa(dg.(int))
		sds = strings.Repeat("0", 5-len(sds)) + sds
	case string:
		sds = dg.(string)
	}

	if sds[2:5] == "600" {
		var sf int
		sf, _ = strconv.Atoi(sds[0:2])
		sf--
		return provincia[sf].Postnombre, provincia[sf], nil
	}

	for _, val := range provincia {
		if val.Postnombre == sds {
			return val.Postnombre, val, nil
		}
	}

	sds = ""
	return sds, sst, &EnumError{"No hay respuesta"}
}
