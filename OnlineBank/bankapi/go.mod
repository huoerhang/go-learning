module bankapi

go 1.16

require (
	"bankcore" v0.0.0
)
replace (
	"bankcore" => "../bankcore"
)