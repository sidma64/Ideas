BUILD="cc $(pkg-config --cflags gtk4) -o app main.c $(pkg-config --libs gtk4)"

eval $BUILD
bear -- $BUILD
