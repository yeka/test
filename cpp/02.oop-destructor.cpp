#include <iostream>

using namespace std;

class MyClass {
public:
    MyClass() {
        cout << "contructing\n";
    }

    ~MyClass() {
        cout << "destructing\n";
    }
};

void InitClass() {
    MyClass a;
    cout << "Return\n";
}

void InitClass2() {
    MyClass* a = new MyClass();
    cout << "Return\n";

    // new keyword assign to heap, and must be removed manually.
    // Uncomment next line and re-run the program
    delete a;
}

int main() {
    InitClass();
    cout << "\n";

    InitClass2();
    cout << "\n";

    cout << "Done\n";
    return 0;
}
