#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <assert.h>
#include <link.h>
#include <unistd.h>
#include <string.h>
// Testing in libc-2.35
// gcc slime_machine_re_re.c -o chal -fno-stack-protector -fPIE -fPIC -Wl,-z,relro -Wl,-z,now -s

char *slime_ptrs[0x100];
int slime_size[0x100];
int slime_status[0x100];
int idx = 0;

void init()
{
    setvbuf(stdin, 0, _IONBF, 0);
    setvbuf(stdout, 0, _IONBF, 0);
}

void menu()
{
    puts("=== NAUP'S SLIME MACHINE ===");
    puts("1) Create SLIME");
    puts("2) Get SLIME information");
    puts("3) Set SLIME information");
    puts("4) KILL SLIME");
    puts("5) GIFT");
    puts("6) End");
    printf(">>> ");
}

void create()
{
    int size;

    if (idx >= 0x100) {
        return;
    }

    printf("size: ");
    scanf("%d", &size);
    
    if(size > 0x90){
        puts("Too large");
    }
    else{
        slime_size[idx] = size;
        slime_ptrs[idx] = malloc(size);
        slime_status[idx] = 1;
        
        printf("Create!!! SLIME[%d]\n", idx);

        idx++;
    }
}


void get()
{
    int idx;

    printf("index: ");
    scanf("%d", &idx);
    if(idx < 0 || idx > 0x100){
        printf("error index\n");
        return;
    }
    printf("SLIME[%d]: %s\n", idx, slime_ptrs[idx]);

}

void set()
{
    int idx;

    printf("index: ");
    scanf("%d", &idx);
    if(idx < 0 || idx > 0x100){
        printf("error index\n");
        return;
    }
    printf("data: ");
    read(0, slime_ptrs[idx], slime_size[idx]);
    
}

void delete()
{
    int idx;

    printf("index: ");
    scanf("%d", &idx);
    
    if (slime_ptrs[idx] && slime_status[idx]) {
        free(slime_ptrs[idx]);
        slime_status[idx] = 0;
    }
}

int LD_GIFT(struct dl_phdr_info *info, size_t size, void *data) {

    if (strstr(info->dlpi_name, "/lib64/ld-linux-x86-64.so.2") != NULL) {
        printf("gift: %p\n", (void *)info->dlpi_addr);
    }
    return 0;
}


void win() {
    system("/bin/sh");
}

int a=1;

int main(void)
{
    init();
    while(a) {
        menu();
        int num;
        scanf("%d", &num);
        switch(num) {
        case 1:
            create();
            break;
        case 2:
            get();
            break;
        case 3:
            set();
            break;
        case 4:
            delete();
            break;
        case 5:
            if (dl_iterate_phdr(LD_GIFT, NULL) == -1) {
                perror("dl_iterate_phdr failed");
                return EXIT_FAILURE;
            }
            break;
        case 6:
            a=0;
            break;
        default:
            printf("Err!!!");
            break;
        }
    }
    return 0;
}
