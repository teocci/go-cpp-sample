liblibrary.so: routines.cpp
	clang++ -o libroutines.so routines.cpp \
	-std=c++17 -O3 -Wall -Wextra -fPIC -shared

routines: routines.go liblibrary.so
	go build routines.go
