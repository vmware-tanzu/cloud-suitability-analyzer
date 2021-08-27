module github.com/vmware-tanzu/cloud-suitability-analyzer/go

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20201120081800-1786d5ef83d4 // indirect
	github.com/antchfx/xmlquery v1.3.3
	github.com/blevesearch/bleve v1.0.12
	github.com/cznic/b v0.0.0-20181122101859-a26611c4d92d // indirect
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/cznic/strutil v0.0.0-20181122101858-275e90344537 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/elazarl/go-bindata-assetfs v1.0.1
	github.com/emirpasic/gods v1.12.0
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/fatih/camelcase v1.0.0
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.7.0
	github.com/jinzhu/gorm v1.9.16
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/joho/sqltocsv v0.0.0-20190824231449-5650f27fd5b6
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/lib/pq v1.8.0
	github.com/mattn/go-sqlite3 v1.14.5
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/browser v0.0.0-20201112035734-206646e67786
	github.com/pkg/profile v1.5.0
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/src-d/go-oniguruma v1.1.0 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/tecbot/gorocksdb v0.0.0-20191217155057-f0fad39f321c // indirect
	github.com/toqueteos/trie v1.0.0 // indirect
	github.com/vmware-labs/yaml-jsonpath v0.0.0-20200625154356-ea62dcd51756
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/src-d/enry.v1 v1.7.3
	gopkg.in/toqueteos/substring.v1 v1.0.2 // indirect
	gopkg.in/xmlpath.v2 v2.0.0-20150820204837-860cbeca3ebc
	gopkg.in/yaml.v2 v2.3.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)

replace github.com/vmware-samples/cloud-suitability-analyzer => github.com/vmware-tanzu/cloud-suitability-analyzer/go v3.2.3-rc1.0.20201120140803-ee37f798900b+incompatible

go 1.15
