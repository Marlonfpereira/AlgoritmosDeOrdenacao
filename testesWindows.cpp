#include "Windows.h"
#include <string>
#include <vector>
using namespace std;

void run(string call) {
    for (int i = 10; i <= 100000; i*=10)
    {
        char aux[7];
        string chamada = call;
        itoa(i, aux, 10);
        chamada.append(aux);
        system(chamada.c_str());
        // printf("%s\n", chamada.c_str());
    }
}

int main() {
    vector<string> lista;
    // lista.push_back("start /wait ./insertion/cpp/insertionNomes.exe ");
    // lista.push_back("start /wait ./insertion/cpp/insertionNums.exe ");
    // lista.push_back("start /wait ./insertion/cpp/insertionPessoa.exe ");
    lista.push_back("start /wait go run ./insertion/golang/insertionNomes.go ");
    // lista.push_back("start /wait go run ./insertion/golang/insertionNums.go ");
    // lista.push_back("start /wait go run ./insertion/golang/insertionPessoa.go ");
    // lista.push_back("start /wait ./quick/cpp/quickNomes.exe ");
    // lista.push_back("start /wait ./quick/cpp/quickNums.exe ");
    // lista.push_back("start /wait ./quick/cpp/quickPessoa.exe ");
    // lista.push_back("start /wait go run ./quick/golang/quickNomes.go ");
    // lista.push_back("start /wait go run ./quick/golang/quickNums.go ");
    // lista.push_back("start /wait go run ./quick/golang/quickPessoa.go ");

    for (const auto& i: lista)
    {
        run(i);
    }
    

    return 0;
}