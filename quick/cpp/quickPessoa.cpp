#include <iostream>
#include <fstream>
#include <vector>
#include <time.h>
#include <string>
#include <sstream>
using namespace std;

class Pessoa
{
public:
    int codigo;
    string name;
    Pessoa(){};
    Pessoa(int cod, string nome)
    {
        codigo = cod;
        name = nome;
    }
    void PrintPessoa()
    {
        cout << codigo << ' ' << name << endl;
    }
    bool operator<(Pessoa b)
    {
        return codigo < b.codigo;
    }
};

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
        while (values[i].name < pivo.name && i < end)
        {
            i++;
        }
        while (values[j].name > pivo.name && j > began)
        {
            j--;
        }
        comp++;
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
    ifstream input("pessoa 1.txt");
    ofstream output("./quickCppPessoa.txt", ios_base::app);
    vector<Pessoa> vetor;
    unsigned long comp = 0, swap = 0;
    clock_t inicio, fim;

    for (int i = 0; i < atoi(argv[1]); i++)
    {
        string linha;

        getline(input, linha);
        stringstream ss(linha);
        string id_str, nome;

        getline(ss, id_str, '\t');
        getline(ss, nome, '\t');

        int id = stoi(id_str);
        vetor.push_back(Pessoa(id, nome));
    }

    inicio = clock();

    quicksort(vetor, 0, size(vetor), comp, swap);

    fim = clock();
    output << argv[1] << "," << (float)(((fim - inicio) + 0.0) / CLOCKS_PER_SEC) << "," << comp << "," << swap << "\n";

    return 0;
}