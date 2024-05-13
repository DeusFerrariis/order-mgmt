module github.com/DeusFerrariis/order-mgmt

go 1.22.2

require (
	github.com/DeusFerrariis/order-mgmt/customer v0.0.0-00010101000000-000000000000
	github.com/DeusFerrariis/order-mgmt/order v0.0.0-00010101000000-000000000000
	github.com/charmbracelet/log v0.4.0
	github.com/labstack/echo/v4 v4.12.0
)

replace github.com/DeusFerrariis/order-mgmt/customer => ./customer

replace github.com/DeusFerrariis/order-mgmt/handle => ./handle

replace github.com/DeusFerrariis/order-mgmt/order => ./order

require (
	github.com/DeusFerrariis/order-mgmt/handle v0.0.0-00010101000000-000000000000 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/charmbracelet/lipgloss v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
