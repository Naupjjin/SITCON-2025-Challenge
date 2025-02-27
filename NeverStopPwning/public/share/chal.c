#include <stdio.h>

// gcc  chal.c -o chal -fno-stack-protector -no-pie --static
int main(){
	char buffer[8];
	printf("Go!!!! ");
	gets(buffer);
	return 0;
}
