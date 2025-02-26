#include <stdlib.h>
#include <stdio.h>

// gcc chal.c -o chal -fno-stack-protector

void win() {
    puts("Forever ...");
    FILE *fp = fopen("/home/chal/flag.txt", "r");   
    if (fp == NULL) {
        perror("Error opening file");
        return 1;
    }

    char buffer[1024];  
    while (fgets(buffer, sizeof(buffer), fp) != NULL) {
        printf("%s", buffer);  
    }

    fclose(fp);
}

int main() {
    char buf[0x30];
    setvbuf(stdin, NULL, _IONBF, 0);
    setvbuf(stdout, NULL, _IONBF, 0);

    puts("CRYCHIC Never Died !!!!! Go to watch this video https://youtu.be/NJ1tne9u8YM?si=oH2lj2lmY5QXO5HD !");
    read(0, buf, 0x100);
}


