module example.com/localmodexample

go 1.13

require (
	example.org/auth v0.0.0
	example.org/deploy v0.0.0
	github.com/chzyer/flagly v0.0.0-20200319234010-7251fe846e8e // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gosuri/uilive v0.0.4 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/olekukonko/ts v0.0.0-20171002115256-78ecb04241c0 // indirect
	github.com/ukautz/reflekt v0.0.0-20180611090553-6ce38d64d188 // indirect
	gopkg.in/ukautz/clif.v1 v1.0.0-20190218144324-df36acc24204
)

replace (
	example.org/auth => ./auth
	example.org/deploy => ./deploy
)
