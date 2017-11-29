#include <tchar.h>
#include <windows.h>
#include <stdio.h>

#define DLL_NAME "example.dll"

typedef int (*PrintVersionFunc)();

#define ErrorExit(funcName) ErrorExitFL((funcName), __FILE__, __LINE__)

void ErrorExitFL(LPSTR lpszFunction, LPSTR file, int line) 
{ 
    // Retrieve the system error message for the last-error code
    LPVOID lpMsgBuf;
    DWORD dw = GetLastError(); 

    FormatMessageA(
        FORMAT_MESSAGE_ALLOCATE_BUFFER | 
        FORMAT_MESSAGE_FROM_SYSTEM |
        FORMAT_MESSAGE_IGNORE_INSERTS,
        NULL,
        dw,
        MAKELANGID(LANG_NEUTRAL, SUBLANG_DEFAULT),
        (LPSTR) &lpMsgBuf,
        0, NULL );

    // Display the error message and exit the process
    fprintf(stderr, "%s:%d:ERROR: %s failed with error %d: %s", file, line, lpszFunction, dw, lpMsgBuf);

    LocalFree(lpMsgBuf);
    ExitProcess(dw); 
}

int main(){
    HINSTANCE hDll = NULL;
    PrintVersionFunc printVersion = NULL;

    hDll = LoadLibrary(DLL_NAME);
    if (!hDll){
        ErrorExit(("LoadLibrary"));
        return 1;
    }

    printVersion = (PrintVersionFunc)GetProcAddress(hDll, ("PrintVersion"));
    if (!printVersion){
        ErrorExit("GetProcAddress");
        return 2;
    }
    
    return printVersion();
}
