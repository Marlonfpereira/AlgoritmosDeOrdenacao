#include <iostream>
#include <fstream>
#include <vector>
#include <time.h>

int main(int argc, char *argv[2])
{
    std::ifstream input("numeros 1.txt");
    std::ofstream output("./insCppNums.txt", std::ios_base::app);
    std::vector<int> vetor;
    int aux;
    unsigned long comp = 0, swap = 0;
    clock_t inicio, fim;

    for (int i = 0; i < std::atoi(argv[1]); i++)
    {
        input >> aux;
        vetor.push_back(aux);
    }

    inicio = clock();

    for (int i = 0; i < std::size(vetor)-1; i++)
    {
        for (int j = i + 1; j > 0; j--)
        {   
            comp++;
            if (vetor[j] < vetor[j - 1])
            {
                aux = vetor[j];
                vetor[j] = vetor[j - 1];
                vetor[j - 1] = aux;
                swap++;
            }
            else
                break;
        }
    }

    fim = clock();
    output << argv[1] << "," << (float)(((fim - inicio) + 0.0) / CLOCKS_PER_SEC) << "," << comp << "," << swap << "\n";

    
    return 0;
}