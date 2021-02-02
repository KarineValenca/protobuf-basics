protoc -I golang/ --go_out=golang/ golang/simple/simple.proto
protoc -I golang/ --go_out=golang/ golang/enum_example/enum_example.proto
protoc -I golang/ --go_out=golang/ golang/complex/complex_example.proto
protoc -I golang-exercise/ --go_out=golang-exercise/ golang-exercise/proto/addressbook.proto
