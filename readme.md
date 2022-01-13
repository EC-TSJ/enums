# enums

[![Home](https://godoc.org/github.com/gookit/event?status.svg)](file:///D:/EC-TSJ/Documents/CODE/SOURCE/Go/pkg/lib/cli)
[![Build Status](https://travis-ci.org/gookit/event.svg?branch=master)](https://travis-ci.org/)
[![Coverage Status](https://coveralls.io/repos/github/gookit/event/badge.svg?branch=master)](https://coveralls.io/github/)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/event)](https://goreportcard.com/report/github.com/)

> **[EN README](README.md)**

Enums es una librería para gestionar un sistema de enumsm basadi en JSON.

## GoDoc

- [godoc for github](https://godoc.org/github.com/)

## Funciones Principales
--- 

Tiene las siguientes funciones:

---

- `Poblaciones(hp ...int) ([]string, []PoblacionesT) `
- `Bancos() ([]string, []BancoT)`
- `Provincias() ([]string, []ProvinciaT) `
- `ComunidadesAutonomas() ([]string, []CCAA_T)` 
- `Delegaciones() ([]string, []ProvinciaT)`
- `Administraciones() ([]string, []AeatT) `
- `Siglas() ([]string, []SiglaT) `
- `Paises() ([]string, []PaisesT)` 
- `AdministracionFromProvincia(name interface{}, code int) (interface{}, AeatT) `
- `PoblacionFromProvincia(name interface{}, code int) (interface{}, PoblacionesT) `
- `CCAAFromDelegacion(delegacion int) (string, CCAA_T) `
- `AdministracionesFromCCAA(ccaa interface{}) ([]string, []AeatT) `
- `ProvinciaFromAdministracion(administracion int) (string, ProvinciaT) `
- `AdministracionesFromDelegacion(delegacion interface{}) ([]string, []AeatT) `
- `DelegacionFromAdministracion(administracion int) (string, ProvinciaT)`
- `ProvinciaFromPoblacion(poblacion int) (string, PoblacionT, error) `
- `ProvinciasFromDelegacion(ccaa interface{}) ([]string, []ProvinciaT)` 
- `ProvinciasFromCCAA(ccaa interface{}) ([]string, []ProvinciaT) `
- `CCAAFromProvincia(province interface{}) (interface{}, CCAA_T, error) `
- `Sigla(name string) (string, SiglaT, error) `
- `Pais(name interface{}) (string, PaisesT, error)`
- `Banco(code interface{}) (string, BancoT, error)  `
- `Provincia(pr interface{}) (string, ProvinciaT) `
- `Poblacion(hp interface{}) (string, PoblacionesT, error) `
- `CCAA(ca interface{}) (string, CCAA_T, error) `
- `Administracion(dg interface{}) (string, AeatT, error)`
- `Delegacion(dg interface{}) (string, ProvinciaT, error)` 

Y tipos:

--- 
- BancoT  
- PoblacionT  
- PoblacionesT 
- ProvinciaT 
- CodeT 
- AeatT  
- PaisesT 
- CCAA_T  
- SiglaT  
- EnumError 

## Ejemplos
```go

```
## Notas




<!-- - [gookit/ini](https://github.com/gookit/ini) INI配置读取管理，支持多文件加载，数据覆盖合并, 解析ENV变量, 解析变量引用
-->
## LICENSE

**[MIT](LICENSE)**
