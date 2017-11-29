#include <stdio.h>
#include "../example-go-lib/example.h"

// force gcc to link in go runtime (may be a better solution than this)
void dummy(){
    PrintVersion();
}

int main(){
    
}

