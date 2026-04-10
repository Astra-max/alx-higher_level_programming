using System;
namespace IfStatements.Switches
{
    public class SwitchCases
    {
        public void SwitchCont()
        {
            string? input = UserInput("What is your age?: ");
            int age = Convert.ToInt32(input);

            switch (age)
            {
                case >= 40:
                    Console.WriteLine("You are an adult.");
                    break;
                case <= 18:
                    Console.WriteLine("You are a minor.");
                    break;
                default:
                    Console.WriteLine("You are a young adult.");
                    break;
            }
        }

        static string? UserInput(string message)
        {
            Console.Write(message);
            return Console.ReadLine();
        }
    }
}