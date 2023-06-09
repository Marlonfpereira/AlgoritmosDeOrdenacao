#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <time.h>
using namespace std;

template <typename T>
void quicksort(vector<T> &values, int began, int end, unsigned long &comp, unsigned long &swap)
{
    int i, j;
    T pivo, aux;
    i = began;
    j = end - 1;
    pivo = values[(began + end) / 2];
    while (i <= j)
    {
        while (values[i] < pivo && i < end)
        {
            i++;
            comp+=2;
        }
        while (values[j] > pivo && j > began)
        {
            j--;
        comp+=2;
        }
        if (i <= j)
        {
            aux = values[i];
            values[i] = values[j];
            values[j] = aux;
            swap++;
            i++;
            j--;
        }
    }
    if (j > began)
        quicksort(values, began, j + 1, comp, swap);
    if (i < end)
        quicksort(values, i, end, comp, swap);
}

int main(int argc, char *argv[2])
{
    ifstream input("numeros 1.txt");
    ofstream output("./quickCppNums.txt", ios_base::app);
    vector<int> vetor;
    int aux;
    unsigned long comp = 0, swap = 0;
    clock_t inicio, fim;

    for (int i = 0; i < atoi(argv[1]); i++)
    {
        input >> aux;
        vetor.push_back(aux);
    }

    inicio = clock();

    quicksort(vetor, 0, size(vetor), comp, swap);

    fim = clock();

    output << argv[1] << "," << (float)(((fim - inicio) + 0.0) / CLOCKS_PER_SEC) << "," << comp << "," << swap << "\n";

    return 0;
}