module example.com/tester

go 1.23.0

replace example.com => ../practice

replace example.com/practice => ../practice

require example.com/practice v0.0.0-00010101000000-000000000000
