module secretsample

go 1.12

require (
	github.com/eliben/gocdkx v0.40.0
	github.com/eliben/gocdkx/contrib/secrets/vault v0.40.0
)

replace github.com/eliben/gocdkx/contrib/secrets/vault => ../../gocdkx/contrib/secrets/vault
