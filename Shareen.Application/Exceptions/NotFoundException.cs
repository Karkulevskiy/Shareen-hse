namespace Shareen.Application.Exceptions;

public class NotFoundException(string property, string entity)
    : Exception($"'{entity}' with '{property}' NOT FOUND\n");