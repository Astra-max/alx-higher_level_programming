using System;
using Microsoft.VisualBasic;
namespace Variables
{
    class Variables
    {
        static void Main(string[] args)
        {
            string name = "John";
            int age = 30;
            double height = 1.75;
            bool isStudent = true;

            Console.WriteLine("Name: " + name);
            Console.WriteLine("Age: " + age);
            Console.WriteLine("Height: " + height);
            Console.WriteLine("Is Student: " + isStudent);
            Constants();
        }

        static void Constants()
        {
            const double PI = 3.142;
            const int YEAR = 2024;
            Console.WriteLine("PI: " + PI);
            Console.WriteLine("Year: " + YEAR);
        }
    }
}
