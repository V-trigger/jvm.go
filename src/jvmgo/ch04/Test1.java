package com.test.stack;

public class Test1 {

    public static void main(String[] args) {
        Test1 test1 = new Test1();
        int add = test1.add(1, 2);
        System.out.println(add);
    }

    private int num;

    public int add(int a, int b){
        num = a + b;
        return num;
    }

}
