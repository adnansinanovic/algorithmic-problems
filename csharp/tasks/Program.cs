using System;
using System.Collections.Generic;
using tasks.Task.Cr001;
using tasks.Task.Gfg001;
using tasks.Task.Task001;
using tasks.Task.Task002;
using tasks.Task.Task003;
using tasks.Task.Task004;
using tasks.Task.Task005;

namespace tasks
{
    public interface ICodeTask
    {
        string ShortDescription { get; }
        string Title { get; }
        string Description { get; }
        string URL { get; }
        void Show();
    }

    class Program
    {
        enum TaskType
        {
            Task001, Task002, Task003, Task004, Task005, 
            Cr001,
            Gfg001,
        }


        static void Main(string[] args)
        {
            var tasks = Initialize();
            var selection = TaskType.Gfg001;

            tasks[selection].Show();

            Console.WriteLine("............ Press any key to continue ................");
            Console.ReadKey();
        }

        static Dictionary<TaskType, ICodeTask> Initialize()
        {
            return new Dictionary<TaskType, ICodeTask>
            {
                {TaskType.Task001, new Task001()},
                {TaskType.Task002, new Task002()},
                {TaskType.Task003, new Task003()},
                {TaskType.Task004, new Task004()},
                {TaskType.Task005, new Task005()},
                {TaskType.Cr001, new Cr001()},
                {TaskType.Gfg001, new Gfg001()},
            };
        }
    }
}
