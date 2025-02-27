#include <stdlib.h>
#include <stdio.h>

// gcc chal.c -o chal -fno-stack-protector -no-pie

char buf_bss[0x20];

int main() {
    char buf[0x30];
    setvbuf(stdin, NULL, _IONBF, 0);
    setvbuf(stdout, NULL, _IONBF, 0);
    puts("Give Me some string >");
    read(0,buf_bss,0x20);
    puts("Give Me ROP >");
    read(0, buf, 0x100);
}


