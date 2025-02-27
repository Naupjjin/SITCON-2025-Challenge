#include <stdio.h>

// gcc  chal.c -o chal -fno-stack-protector -no-pie --static
int main(){
	setvbuf(stdin, NULL, _IONBF, 0);
	setvbuf(stdout, NULL, _IONBF, 0);
	char buffer[8];
	printf("Go!!!! ");
	gets(buffer);
	return 0;
}
