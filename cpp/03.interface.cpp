#include <iostream>

using namespace std;

// A class is made abstract by declaring at least one of its functions as pure virtual function.
// A pure virtual function is specified by placing "= 0"
class Shape {
public:
    virtual string Name() = 0;
    virtual int Area() = 0;
};

class Triangle: public Shape{
private:
    int height, base;
public:
    Triangle(int height, int base): height(height), base(base) {}
    string Name() {
        return "Triangle";
    }
    int Area() {
        return height * base / 2;
    }
};

class Square: public Shape {
private:
    int s;
public:
    Square(int size): s(size) {}
    string Name() {
        return "Square";
    }
    int Area() {
        return s * s;
    }
};

void Print(Shape& s) {
    printf("Area of %s is %d\n", s.Name().c_str(), s.Area());
}

struct Hello {
    Triangle t;
    Square s;
};

int main() {
    Hello b = {
        Triangle(10, 2),
        Square(3),
    };
    // Triangle t(10, 2);
    // Square s(3);
    // Print(t);
    // Print(s);
    Print(b.t);
    Print(b.s);
    return 0;
}