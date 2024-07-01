using GraphQL.Models;
using GraphQL.Repositories;

namespace GraphQL
{
    public class Query
    {
        public string Hello() => "Hello, world!";

        private readonly IRepository _repository;

        public Query(IRepository repository)
        {
            _repository = repository;
        }
    }
}
