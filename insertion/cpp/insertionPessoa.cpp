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
};

int main(int argc, char *argv[2])
{
    ifstream input("pessoa 1.txt");
    ofstream output("./insCppPessoa.txt", ios_base::app);
    
    vector<Pessoa> vetor;
    unsigned long long comp = 0, swap = 0;
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

    for (int i = 0; i < size(vetor) - 1; i++)
    {
        Pessoa aux;
        for (int j = i + 1; j > 0; j--)
        {
            comp += 3;
            if (vetor[j].codigo < vetor[j - 1].codigo || (vetor[j].codigo == vetor[j - 1].codigo && vetor[j].name < vetor[j - 1].name))
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
    output << argv[1] << "," << float(((fim - inicio) + 0.0) / CLOCKS_PER_SEC) << "," << comp << "," << swap << "\n";

    return 0;
}