using System;
namespace IfStatements
{
    class Program
    {
        static void Main(string[] args)
        {
            string? input = UserInput("What is your age?: ");
            int age = Convert.ToInt32(input);

            if (age >= 40)
            {
                Console.WriteLine("You are an adult.");
            }
            else if (age <= 18)
            {
                Console.WriteLine("You are a minor.");
            }
            else
            {
                Console.WriteLine("You are a young adult.");
            }
        }

        static string? UserInput(string message)
        {
            Console.Write(message);
            return Console.ReadLine();
        }
    }
}