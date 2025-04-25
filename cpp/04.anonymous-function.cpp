#include <iostream>

// Example 1 -> anonymous function (lambda) that doesn't return anything
void Do(void(*op)(void)) {
    op();
}

void example1() {
    Do([]() {
        std::cout << "Ok" << std::endl;
    });
}

// Example 2 -> anonymous function (lambda) with parameter and return type

int Operate(int a, int b, int (*op)(int, int)) {
    return op(a, b);
}

int add(int a, int b) {
    return a + b;
}

int mul(int a, int b) {
    return a * b;
}

void example2() {
    std::printf("Add: %d\n", Operate(2, 3, add));
    std::printf("Mul: %d\n", Operate(2, 3, mul));
    std::printf("Sub: %d\n", Operate(2, 3, [](int a, int b) -> int {
        return a - b;
    }));
}

// Example 3 -> anonymous function (lambda) using functional library
#include <functional>

typedef std::function<void(void)> Handler;
typedef std::function<int(int, int)> Operator;

void Handle(Handler h) {
    h();
}

void example3a() {
    Handle([]() { std::cout << "Example 3\n"; });
}

int Operate3(int a, int b, Operator op) {
    return op(a, b);
}

void example3b() {
    Operator res = [](int a, int b)->int{
        return a + b;
    };
    std::cout << Operate3(2, 3, res) << std::endl;
}

int main() {
    example1();
    example2();
    example3a();
    example3b();
}